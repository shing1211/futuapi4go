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
	mu          sync.RWMutex
	config      *PoolConfig
	clients     map[PoolType][]*PoolConn
	clientIndex map[*Client]*PoolConn
	ctx         context.Context
	cancel      context.CancelFunc
	wg          sync.WaitGroup
	closed      bool
	cond        *sync.Cond
}

func NewClientPool(config *PoolConfig) *ClientPool {
	ctx, cancel := context.WithCancel(context.Background())
	p := &ClientPool{
		config:      config,
		clients:     make(map[PoolType][]*PoolConn),
		clientIndex: make(map[*Client]*PoolConn),
		ctx:        ctx,
		cancel:     cancel,
	}
	p.cond = sync.NewCond(&p.mu)
	return p
}

func (p *ClientPool) Get(ctx context.Context, poolType PoolType) (*Client, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil, NewError(CodePoolClosed, "pool is closed")
	}

	for {
		select {
		case <-ctx.Done():
			return nil, NewError(CodePoolExhausted, "pool exhausted: context timed out waiting for available connection")
		default:
		}

		conns := p.clients[poolType]

		for _, pc := range conns {
			if !pc.InUse && time.Since(pc.LastUsed) < p.config.MaxIdleTime {
				pc.InUse = true
				pc.LastUsed = time.Now()
				return pc.Client, nil
			}
		}

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
			p.clientIndex[client] = pc
			return client, nil
		}

		done := make(chan struct{})
		go func() {
			p.cond.Wait()
			close(done)
		}()

		p.mu.Unlock()
		select {
		case <-ctx.Done():
			p.mu.Lock()
			p.cond.Broadcast()
			<-done
			return nil, NewError(CodePoolExhausted, "pool exhausted: context timed out waiting for available connection")
		case <-done:
		}
		p.mu.Lock()

		if p.closed {
			return nil, NewError(CodePoolClosed, "pool is closed")
		}
	}
}

func (p *ClientPool) Put(client *Client) {
	p.mu.Lock()
	defer p.mu.Unlock()

	pc, ok := p.clientIndex[client]
	if !ok {
		return
	}
	pc.InUse = false
	pc.LastUsed = time.Now()
	p.cond.Signal()
}

// Remove removes a client from the pool (e.g., if it's broken).
func (p *ClientPool) Remove(client *Client) {
	p.mu.Lock()
	defer p.mu.Unlock()

	pc, ok := p.clientIndex[client]
	if !ok {
		return
	}
	pc.Client.Close()
	poolType := pc.PoolType

	// Remove from slice
	conns := p.clients[poolType]
	for i, c := range conns {
		if c == pc {
			p.clients[poolType] = append(conns[:i], conns[i+1:]...)
			break
		}
	}

	// Remove from index
	delete(p.clientIndex, client)
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

// PoolStats holds connection pool statistics for a single pool type.
type PoolStats struct {
	Total  int // Total connections in the pool
	InUse  int // Connections currently in use
	Idle   int // Connections available for use
}

// Stats returns a snapshot of pool statistics keyed by PoolType.
func (p *ClientPool) Stats() map[PoolType]PoolStats {
	p.mu.RLock()
	defer p.mu.RUnlock()
	result := make(map[PoolType]PoolStats)
	for pt, conns := range p.clients {
		s := PoolStats{Total: len(conns)}
		for _, pc := range conns {
			if pc.InUse {
				s.InUse++
			} else {
				s.Idle++
			}
		}
		result[pt] = s
	}
	return result
}

// GetPoolType returns the PoolType of the given client. Returns (PoolTypeGeneral, false)
// if the client is not in the pool.
func (p *ClientPool) GetPoolType(client *Client) (PoolType, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	pc, ok := p.clientIndex[client]
	if !ok {
		return PoolTypeGeneral, false
	}
	return pc.PoolType, true
}

func (p *ClientPool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil
	}
	p.closed = true
	p.cancel()
	p.cond.Broadcast()

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
