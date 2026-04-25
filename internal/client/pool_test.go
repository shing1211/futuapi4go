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
	"sync"
	"testing"
	"time"
)

func TestPoolConfigDefaults(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	if config.MaxSize != 3 {
		t.Errorf("expected MaxSize 3, got %d", config.MaxSize)
	}
	if config.MinIdle != 1 {
		t.Errorf("expected MinIdle 1, got %d", config.MinIdle)
	}
	if config.MaxIdleTime != 5*time.Minute {
		t.Errorf("expected MaxIdleTime 5m, got %v", config.MaxIdleTime)
	}
	if config.HealthCheckInterval != 30*time.Second {
		t.Errorf("expected HealthCheckInterval 30s, got %v", config.HealthCheckInterval)
	}
	if config.Addr != "127.0.0.1:11111" {
		t.Errorf("expected Addr 127.0.0.1:11111, got %s", config.Addr)
	}
}

func TestPoolTypeString(t *testing.T) {
	tests := []struct {
		poolType PoolType
		expected string
	}{
		{PoolTypeMarketData, "MarketData"},
		{PoolTypeTrading, "Trading"},
		{PoolTypeGeneral, "General"},
		{PoolType(999), "Unknown"},
	}

	for _, tt := range tests {
		if got := tt.poolType.String(); got != tt.expected {
			t.Errorf("PoolType(%d).String() = %q, want %q", tt.poolType, got, tt.expected)
		}
	}
}

func TestNewClientPool(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	pool := NewClientPool(config)
	if pool == nil {
		t.Fatal("NewClientPool returned nil")
	}
	if pool.closed {
		t.Error("pool should not be closed initially")
	}
}

func TestPoolGetReturnsErrorWhenClosed(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	pool := NewClientPool(config)
	pool.Close()

	_, err := pool.Get(context.Background(), PoolTypeMarketData)
	if err == nil {
		t.Error("Get should return error when pool is closed")
	}
}

func TestPoolCloseIsIdempotent(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	pool := NewClientPool(config)

	// Close multiple times should not panic
	pool.Close()
	pool.Close()
	pool.Close()
}

func TestPoolSizeInitialState(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	pool := NewClientPool(config)
	defer pool.Close()

	if pool.Size(PoolTypeMarketData) != 0 {
		t.Errorf("expected Size 0, got %d", pool.Size(PoolTypeMarketData))
	}
}

func TestPoolAvailableInitialState(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	pool := NewClientPool(config)
	defer pool.Close()

	if pool.Available(PoolTypeMarketData) != 0 {
		t.Errorf("expected Available 0, got %d", pool.Available(PoolTypeMarketData))
	}
}

func TestPoolConnReuse(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	config.MaxSize = 2
	config.MinIdle = 0
	pool := NewClientPool(config)
	defer pool.Close()

	client1, err := pool.Get(context.Background(), PoolTypeGeneral)
	if err != nil {
		t.Skip("Cannot connect to server:", err)
	}
	if client1 == nil {
		t.Fatal("Get returned nil client")
	}

	pool.Put(client1)

	client2, err := pool.Get(context.Background(), PoolTypeGeneral)
	if err != nil {
		t.Fatalf("Failed to get second client: %v", err)
	}
	pool.Put(client2)

	if pool.Size(PoolTypeGeneral) != 1 {
		t.Errorf("expected pool size 1, got %d", pool.Size(PoolTypeGeneral))
	}
	if pool.Available(PoolTypeGeneral) != 1 {
		t.Errorf("expected 1 available, got %d", pool.Available(PoolTypeGeneral))
	}
}

func TestPoolConnReuseWithDifferentTypes(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	config.MaxSize = 2
	config.MinIdle = 0
	pool := NewClientPool(config)
	defer pool.Close()

	clientMD, err := pool.Get(context.Background(), PoolTypeMarketData)
	if err != nil {
		t.Skip("Cannot connect to server:", err)
	}
	pool.Put(clientMD)

	clientTrd, err := pool.Get(context.Background(), PoolTypeTrading)
	if err != nil {
		t.Skip("Cannot connect to server:", err)
	}
	pool.Put(clientTrd)

	if pool.Size(PoolTypeMarketData) != 1 {
		t.Errorf("expected market data pool size 1, got %d", pool.Size(PoolTypeMarketData))
	}
	if pool.Size(PoolTypeTrading) != 1 {
		t.Errorf("expected trading pool size 1, got %d", pool.Size(PoolTypeTrading))
	}
}

func TestPoolMaxSizeLimit(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	config.MaxSize = 2
	config.MinIdle = 0
	pool := NewClientPool(config)
	defer pool.Close()

	_, err := pool.Get(context.Background(), PoolTypeGeneral)
	if err != nil {
		t.Skip("Cannot connect to server:", err)
	}
	_, err = pool.Get(context.Background(), PoolTypeGeneral)
	if err != nil {
		t.Skip("Cannot connect to server:", err)
	}

	_, err = pool.Get(context.Background(), PoolTypeGeneral)
	if err == nil {
		t.Error("Expected error when pool exhausted, got nil")
	}
}

func TestPoolRemove(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	config.MaxSize = 2
	config.MinIdle = 0
	pool := NewClientPool(config)
	defer pool.Close()

	client, err := pool.Get(context.Background(), PoolTypeGeneral)
	if err != nil {
		t.Skip("Cannot connect to server:", err)
	}

	initialSize := pool.Size(PoolTypeGeneral)
	pool.Remove(client)

	if pool.Size(PoolTypeGeneral) >= initialSize {
		t.Errorf("expected size to decrease after Remove, was %d now %d", initialSize, pool.Size(PoolTypeGeneral))
	}
}

func TestPoolConcurrentAccess(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	config.MaxSize = 5
	config.MinIdle = 0
	pool := NewClientPool(config)
	defer pool.Close()

	const goroutines = 10
	const requestsPerGoroutine = 20

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < requestsPerGoroutine; j++ {
				ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
				defer cancel()

				client, err := pool.Get(ctx, PoolTypeGeneral)
				if err != nil {
					continue
				}

				pool.Put(client)
			}
		}(i)
	}

	wg.Wait()
}

func TestPoolConcurrentGetPutRemove(t *testing.T) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	config.MaxSize = 10
	config.MinIdle = 0
	pool := NewClientPool(config)
	defer pool.Close()

	const goroutines = 8
	const operationsPerGoroutine = 30

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
				defer cancel()

				client, err := pool.Get(ctx, PoolTypeGeneral)
				if err != nil {
					continue
				}

				if goroutineID%3 == 0 {
					pool.Remove(client)
				} else {
					pool.Put(client)
				}
			}
		}(i)
	}

	wg.Wait()
}
