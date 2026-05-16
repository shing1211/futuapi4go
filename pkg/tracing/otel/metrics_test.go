package otel

import (
	"testing"
	"time"
)

func TestNewOTelMeter_ReturnsNonNil(t *testing.T) {
	m, err := NewOTelMeter()
	if err != nil {
		t.Fatalf("NewOTelMeter failed: %v", err)
	}
	if m == nil {
		t.Fatal("expected non-nil meter")
	}
	m.Close()
}

func TestOTelMeter_RecordFunctionsNoPanic(t *testing.T) {
	m, err := NewOTelMeter()
	if err != nil {
		t.Fatalf("NewOTelMeter failed: %v", err)
	}
	defer m.Close()

	m.RecordConnection("tcp")
	m.RecordDisconnection("tcp")
	m.RecordReconnect("connection_lost")
	m.RecordAPICall("3004", "success", 50*time.Millisecond)
	m.RecordPushMessage("quote")
	m.RecordOpenDUp(true)
	m.RecordOpenDUp(false)
	m.RecordAPIError("3004", "ERR_1000")
	m.RecordRateLimited("3004")
	m.RecordRetry("3004", "1")
	m.RecordBreakerState("test-breaker", 0.0)
}

func TestOTelMeter_CloseMultipleTimes(t *testing.T) {
	m, err := NewOTelMeter()
	if err != nil {
		t.Fatalf("NewOTelMeter failed: %v", err)
	}

	if err := m.Close(); err != nil {
		t.Errorf("first Close failed: %v", err)
	}

	if err := m.Close(); err != nil {
		t.Errorf("second Close failed: %v", err)
	}
}
