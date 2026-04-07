package futuapi

import (
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
}

func TestNewWithOptionsDefaults(t *testing.T) {
	client := NewWithOptions("", 0, 0)
	if client.maxRetries != DefaultMaxRetries {
		t.Errorf("expected maxRetries %d, got %d", DefaultMaxRetries, client.maxRetries)
	}
	if client.reconnectInterval != DefaultReconnectInterval {
		t.Errorf("expected reconnectInterval %v, got %v", DefaultReconnectInterval, client.reconnectInterval)
	}
}

func TestNewWithOptionsCustom(t *testing.T) {
	maxRetries := 5
	reconnectInterval := 10 * time.Second
	client := NewWithOptions("", maxRetries, reconnectInterval)
	if client.maxRetries != maxRetries {
		t.Errorf("expected maxRetries %d, got %d", maxRetries, client.maxRetries)
	}
	if client.reconnectInterval != reconnectInterval {
		t.Errorf("expected reconnectInterval %v, got %v", reconnectInterval, client.reconnectInterval)
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
}
