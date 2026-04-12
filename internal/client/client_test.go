package futuapi

import (
	"context"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := New()
	if client == nil {
		t.Fatal("New() returned nil")
	}
	if client.conn == nil {
		t.Error("conn should not be nil")
	}
	if client.handlers == nil {
		t.Error("handlers should not be nil")
	}
	if client.opts == nil {
		t.Error("opts should not be nil")
	}
}

func TestNewWithOptions(t *testing.T) {
	opts := []Option{
		WithMaxRetries(5),
		WithReconnectInterval(10 * time.Second),
		WithDialTimeout(5 * time.Second),
		WithAPITimeout(15 * time.Second),
		WithLogLevel(2),
	}
	client := New(opts...)
	if client == nil {
		t.Fatal("New() returned nil")
	}
	if client.opts.MaxRetries != 5 {
		t.Errorf("expected MaxRetries 5, got %d", client.opts.MaxRetries)
	}
	if client.opts.ReconnectInterval != 10*time.Second {
		t.Errorf("expected ReconnectInterval 10s, got %v", client.opts.ReconnectInterval)
	}
	if client.opts.DialTimeout != 5*time.Second {
		t.Errorf("expected DialTimeout 5s, got %v", client.opts.DialTimeout)
	}
	if client.opts.APITimeout != 15*time.Second {
		t.Errorf("expected APITimeout 15s, got %v", client.opts.APITimeout)
	}
	if client.opts.LogLevel != 2 {
		t.Errorf("expected LogLevel 2, got %d", client.opts.LogLevel)
	}
}

func TestNewOptionsDefaults(t *testing.T) {
	opts := NewOptions()
	if opts.DialTimeout != DefaultDialTimeout {
		t.Errorf("expected DialTimeout %v, got %v", DefaultDialTimeout, opts.DialTimeout)
	}
	if opts.APITimeout != DefaultTimeout {
		t.Errorf("expected APITimeout %v, got %v", DefaultTimeout, opts.APITimeout)
	}
	if opts.KeepAliveInterval != DefaultKeepAliveInterval {
		t.Errorf("expected KeepAliveInterval %v, got %v", DefaultKeepAliveInterval, opts.KeepAliveInterval)
	}
	if opts.MaxRetries != DefaultMaxRetries {
		t.Errorf("expected MaxRetries %d, got %d", DefaultMaxRetries, opts.MaxRetries)
	}
	if opts.ReconnectInterval != DefaultReconnectInterval {
		t.Errorf("expected ReconnectInterval %v, got %v", DefaultReconnectInterval, opts.ReconnectInterval)
	}
	if opts.ReconnectBackoff != 1.5 {
		t.Errorf("expected ReconnectBackoff 1.5, got %f", opts.ReconnectBackoff)
	}
}

func TestDefaultConstants(t *testing.T) {
	if DefaultTimeout != 30*time.Second {
		t.Errorf("expected DefaultTimeout 30s, got %v", DefaultTimeout)
	}
	if DefaultKeepAliveInterval != 30*time.Second {
		t.Errorf("expected DefaultKeepAliveInterval 30s, got %v", DefaultKeepAliveInterval)
	}
	if DefaultMaxRetries != 3 {
		t.Errorf("expected DefaultMaxRetries 3, got %d", DefaultMaxRetries)
	}
	if DefaultReconnectInterval != 3*time.Second {
		t.Errorf("expected DefaultReconnectInterval 3s, got %v", DefaultReconnectInterval)
	}
	if DefaultDialTimeout != 10*time.Second {
		t.Errorf("expected DefaultDialTimeout 10s, got %v", DefaultDialTimeout)
	}
}

func TestEnsureConnectedNotConnected(t *testing.T) {
	client := New()
	defer client.Close()

	err := client.EnsureConnected()
	if err == nil {
		t.Error("EnsureConnected should return error when not connected")
	}
	if err != ErrNotConnected {
		t.Errorf("expected ErrNotConnected, got %v", err)
	}
}

func TestIsConnectedInitialState(t *testing.T) {
	client := New()
	defer client.Close()

	if client.IsConnected() {
		t.Error("client should not be connected initially")
	}
}

func TestContextReturnsContext(t *testing.T) {
	client := New()
	defer client.Close()

	ctx := client.Context()
	if ctx == nil {
		t.Error("Context() returned nil")
	}
}

func TestWithContext(t *testing.T) {
	client := New()
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newClient := client.WithContext(ctx)
	if newClient == nil {
		t.Fatal("WithContext returned nil")
	}
	if newClient.Context() != ctx {
		t.Error("WithContext did not set context correctly")
	}
	// Original client should still be usable
	if client.Context() == ctx {
		t.Error("WithContext should not modify original client's context")
	}
}

func TestSetPushHandler(t *testing.T) {
	client := New()
	defer client.Close()

	client.SetPushHandler(func(pkt *Packet) {
		// Handler registered successfully
	})

	// We can't easily test push handling without a real connection,
	// but we can verify the handler is set without panic
	t.Log("SetPushHandler executed successfully")
}

func TestRegisterHandler(t *testing.T) {
	client := New()
	defer client.Close()

	client.RegisterHandler(9999, func(protoID uint32, body []byte) {
		// Handler registered
	})

	// Verify handler is registered
	client.handlersMu.RLock()
	_, ok := client.handlers[9999]
	client.handlersMu.RUnlock()

	if !ok {
		t.Error("handler not registered")
	}
}

func TestSerialNoIncrement(t *testing.T) {
	client := New()
	defer client.Close()

	first := client.NextSerialNo()
	second := client.NextSerialNo()
	third := client.NextSerialNo()

	if second <= first {
		t.Errorf("serial numbers not incrementing: %d, %d", first, second)
	}
	if third <= second {
		t.Errorf("serial numbers not incrementing: %d, %d", second, third)
	}
}
