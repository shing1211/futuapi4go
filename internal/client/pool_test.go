package futuapi

import (
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

	_, err := pool.Get(PoolTypeMarketData)
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

