// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package futuapi

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// PoolType represents the type of connection in the pool.
type PoolType int

const (
	PoolTypeMarketData PoolType = iota // Market data connection
	PoolTypeTrading                    // Trading connection
	PoolTypeGeneral                    // General purpose connection
)

func (t PoolType) String() string {
	switch t {
	case PoolTypeMarketData:
		return "MarketData"
	case PoolTypeTrading:
		return "Trading"
	case PoolTypeGeneral:
		return "General"
	default:
		return "Unknown"
	}
}

// PoolConfig holds configuration for the ClientPool.
type PoolConfig struct {
	MaxSize             int           // Maximum number of connections per pool type
	MinIdle             int           // Minimum idle connections to maintain
	MaxIdleTime         time.Duration // Maximum time a connection can stay idle
	HealthCheckInterval time.Duration // Interval between health checks
	Addr                string        // Futu OpenD address
	Options             []Option      // Client options to apply to all connections
}

// DefaultPoolConfig returns a PoolConfig with sensible defaults.
func DefaultPoolConfig(addr string, opts ...Option) *PoolConfig {
	return &PoolConfig{
		MaxSize:             3,
		MinIdle:             1,
		MaxIdleTime:         5 * time.Minute,
		HealthCheckInterval: 30 * time.Second,
		Addr:                addr,
		Options:             opts,
	}
}

// PoolConn wraps a Client with metadata for pool management.
type PoolConn struct {
	Client    *Client
	PoolType  PoolType
	CreatedAt time.Time
	LastUsed  time.Time
	InUse     bool
}

// ClientPool manages a pool of FutuAPI clients for different purposes.
type ClientPool struct {
	mu      sync.RWMutex
	config  *PoolConfig
	clients map[PoolType][]*PoolConn
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
	closed  bool
}

// NewClientPool creates a new client pool with the given configuration.
func NewClientPool(config *PoolConfig) *ClientPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &ClientPool{
		config:  config,
		clients: make(map[PoolType][]*PoolConn),
		ctx:     ctx,
		cancel:  cancel,
	}
}

// Get retrieves a client of the specified type from the pool.
// If no idle client is available, it waits until a connection becomes available or context times out.
// The context controls the maximum wait time.
func (p *ClientPool) Get(ctx context.Context, poolType PoolType) (*Client, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil, NewError(CodePoolClosed, "pool is closed")
	}

	// Wait loop: retry until we get a connection or context expires
	for {
		// Check for context cancellation
		select {
		case <-ctx.Done():
			return nil, NewError(CodePoolExhausted, "pool exhausted: context timed out waiting for available connection")
		default:
		}

		conns := p.clients[poolType]

		// Try to find an idle connection
		for _, pc := range conns {
			if !pc.InUse && time.Since(pc.LastUsed) < p.config.MaxIdleTime {
				pc.InUse = true
				pc.LastUsed = time.Now()
				return pc.Client, nil
			}
		}

		// Create new connection if pool size permits
		if len(conns) < p.config.MaxSize {
			client, err := p.newClientLocked()
			if err != nil {
				return nil, fmt.Errorf("create new client: %w", err)
			}
			pc := &PoolConn{
				Client:    client,
				PoolType:  poolType,
				CreatedAt: time.Now(),
				LastUsed:  time.Now(),
				InUse:     true,
			}
			p.clients[poolType] = append(conns, pc)
			return client, nil
		}

		// Pool is full, wait for a connection to be released
		// Release lock briefly to allow other goroutines to Put() back
		p.mu.Unlock()
		time.Sleep(50 * time.Millisecond)
		p.mu.Lock()
	}
}

// Put returns a client to the pool after use.
func (p *ClientPool) Put(client *Client) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, conns := range p.clients {
		for _, pc := range conns {
			if pc.Client == client {
				pc.InUse = false
				pc.LastUsed = time.Now()
				return
			}
		}
	}
}

// Remove removes a client from the pool (e.g., if it's broken).
func (p *ClientPool) Remove(client *Client) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for poolType, conns := range p.clients {
		for i, pc := range conns {
			if pc.Client == client {
				pc.Client.Close()
				p.clients[poolType] = append(conns[:i], conns[i+1:]...)
				return
			}
		}
	}
}

// Size returns the number of connections in the pool for a given type.
func (p *ClientPool) Size(poolType PoolType) int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return len(p.clients[poolType])
}

// Available returns the number of available (idle) connections.
func (p *ClientPool) Available(poolType PoolType) int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	count := 0
	for _, pc := range p.clients[poolType] {
		if !pc.InUse {
			count++
		}
	}
	return count
}

// Close closes all connections in the pool and stops the health checker.
func (p *ClientPool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil
	}
	p.closed = true
	p.cancel()

	for _, conns := range p.clients {
		for _, pc := range conns {
			pc.Client.Close()
		}
	}
	p.clients = make(map[PoolType][]*PoolConn)

	p.wg.Wait()
	return nil
}

// StartHealthChecker starts a background goroutine that periodically
// checks connection health and removes stale connections.
func (p *ClientPool) StartHealthChecker() {
	p.wg.Add(1)
	go p.healthCheckLoop()
}

func (p *ClientPool) healthCheckLoop() {
	defer p.wg.Done()

	ticker := time.NewTicker(p.config.HealthCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-ticker.C:
			p.healthCheck()
		}
	}
}

func (p *ClientPool) healthCheck() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for poolType, conns := range p.clients {
		var healthy []*PoolConn
		for _, pc := range conns {
			if pc.InUse {
				healthy = append(healthy, pc)
				continue
			}

			// Check if connection is still alive
			if pc.Client.IsConnected() {
				// Check idle time
				if time.Since(pc.LastUsed) < p.config.MaxIdleTime {
					healthy = append(healthy, pc)
				} else {
					pc.Client.Close()
				}
			} else {
				pc.Client.Close()
			}
		}

		// Ensure minimum idle connections
		for len(healthy) < p.config.MinIdle {
			client, err := p.newClientLocked()
			if err != nil {
				break
			}
			healthy = append(healthy, &PoolConn{
				Client:    client,
				PoolType:  poolType,
				CreatedAt: time.Now(),
				LastUsed:  time.Now(),
				InUse:     false,
			})
		}

		p.clients[poolType] = healthy
	}
}

func (p *ClientPool) newClient() (*Client, error) {
	opts := append([]Option{}, p.config.Options...)
	client := New(opts...)
	if err := client.Connect(p.config.Addr); err != nil {
		return nil, err
	}
	return client, nil
}

func (p *ClientPool) newClientLocked() (*Client, error) {
	return p.newClient()
}
