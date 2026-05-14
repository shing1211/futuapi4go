package ratelimit

import (
	"context"
	"testing"
	"time"
)

func TestLimiterAllow_Basic(t *testing.T) {
	l := NewLimiter(10, 10, ModeReject)
	for i := 0; i < 10; i++ {
		if !l.Allow() {
			t.Fatalf("expected Allow() = true for attempt %d", i+1)
		}
	}
	// 11th should be rejected (no tokens left, rate=10/sec but 0 elapsed)
	if l.Allow() {
		t.Error("expected Allow() = false after exhausting tokens")
	}
}

func TestLimiterAllow_RefillOverTime(t *testing.T) {
	l := NewLimiter(100, 100, ModeReject)
	// Drain all tokens
	for i := 0; i < 100; i++ {
		l.Allow()
	}
	if l.Allow() {
		t.Fatal("expected no tokens left")
	}
	// Wait for refill
	time.Sleep(110 * time.Millisecond)
	if !l.Allow() {
		t.Error("expected tokens to be refilled after 110ms at rate 100/sec")
	}
}

func TestLimiterAllow_CapacityCap(t *testing.T) {
	l := NewLimiter(5, 10, ModeReject)
	// Use 3 tokens
	for i := 0; i < 3; i++ {
		l.Allow()
	}
	// Wait long enough to accumulate more than capacity
	time.Sleep(3 * time.Second)
	// Should be capped at capacity=10, not 5*3+7=22
	l.mu.Lock()
	if l.tokens > l.capacity {
		t.Errorf("tokens %f exceeded capacity %f", l.tokens, l.capacity)
	}
	l.mu.Unlock()
}

func TestLimiterWait_ModeReject(t *testing.T) {
	l := NewLimiter(1, 1, ModeReject)
	// Drain the single token
	l.Allow()
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	err := l.Wait(ctx)
	if err != ErrRateLimited {
		t.Errorf("expected ErrRateLimited, got %v", err)
	}
}

func TestLimiterWait_ModeWait(t *testing.T) {
	l := NewLimiter(1, 1, ModeWait)
	// Drain the single token
	l.Allow()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// Should wait for refill (~1 second at rate 1) but context expires first
	err := l.Wait(ctx)
	if err != context.DeadlineExceeded {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestLimiterWait_ContextCancellation(t *testing.T) {
	l := NewLimiter(1, 1, ModeWait)
	l.Allow() // drain token
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()
	err := l.Wait(ctx)
	if err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}

func TestLimiterWait_TokenAvailable(t *testing.T) {
	l := NewLimiter(100, 100, ModeWait)
	// Tokens should be available immediately
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	err := l.Wait(ctx)
	if err != nil {
		t.Errorf("expected no error when tokens available, got %v", err)
	}
}

func TestModeValues(t *testing.T) {
	if ModeReject == ModeWait {
		t.Error("ModeReject should differ from ModeWait")
	}
}

func TestProtoLimiter_AllowNoPerProto(t *testing.T) {
	pl := NewProtoLimiter(100, 100, ModeReject)
	// No per-proto limit set for protoID 1001
	// Should only check global
	if !pl.Allow(1001) {
		t.Error("expected Allow=true when no per-proto limit is set")
	}
}

func TestProtoLimiter_AllowWithPerProto(t *testing.T) {
	pl := NewProtoLimiter(100, 100, ModeReject)
	pl.SetProtoLimit(1001, 1, 1, ModeReject)
	// Drain per-proto token
	pl.Allow(1001)
	// Should be rejected by per-proto limiter
	if pl.Allow(1001) {
		t.Error("expected Allow=false after draining per-proto tokens")
	}
	// Global should still allow
	if !pl.Allow(9999) {
		t.Error("expected Allow=true for unrelated proto with global tokens available")
	}
}

func TestProtoLimiter_SetProtoLimit(t *testing.T) {
	pl := NewProtoLimiter(100, 100, ModeReject)
	pl.SetProtoLimit(2001, 50, 50, ModeWait)
	// Should not panic and allow calls
	for i := 0; i < 50; i++ {
		if !pl.Allow(2001) {
			t.Fatalf("expected Allow=true for attempt %d", i+1)
		}
	}
}

func TestProtoLimiter_Concurrent(t *testing.T) {
	pl := NewProtoLimiter(1000, 1000, ModeReject)
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				pl.Allow(1001)
			}
			done <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	// Should not have panicked or corrupted state
}

func TestErrRateLimited(t *testing.T) {
	if ErrRateLimited == nil {
		t.Fatal("ErrRateLimited should not be nil")
	}
	if ErrRateLimited.Error() != "rate limit exceeded" {
		t.Errorf("unexpected error message: %s", ErrRateLimited.Error())
	}
}