package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

func TestDo_SuccessFirstAttempt(t *testing.T) {
	calls := 0
	err := Do(context.Background(), DefaultConfig(), func() error {
		calls++
		return nil
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if calls != 1 {
		t.Errorf("expected 1 call, got %d", calls)
	}
}

func TestDo_RetryOnRecoverableError(t *testing.T) {
	calls := 0
	err := Do(context.Background(), Config{
		MaxAttempts:    3,
		BaseDelay:     1 * time.Millisecond,
		MaxDelay:      10 * time.Millisecond,
		IsRecoverable: defaultIsRecoverable,
	}, func() error {
		calls++
		if calls < 3 {
			return constant.NewFutuError(constant.ErrCodeNetworkError, "TestDo_RetryOnRecoverableError", "network error")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if calls != 3 {
		t.Errorf("expected 3 calls, got %d", calls)
	}
}

func TestDo_UnrecoverableError(t *testing.T) {
	calls := 0
	err := Do(context.Background(), Config{
		MaxAttempts:    3,
		BaseDelay:     1 * time.Millisecond,
		MaxDelay:      10 * time.Millisecond,
		IsRecoverable: defaultIsRecoverable,
	}, func() error {
		calls++
		return constant.NewFutuError(constant.ErrCodeInvalidParams, "TestDo_UnrecoverableError", "invalid params")
	})
	if err == nil {
		t.Fatal("expected error for unrecoverable failure")
	}
	if calls != 1 {
		t.Errorf("expected 1 call (no retry for unrecoverable), got %d", calls)
	}
}

func TestDo_AllAttemptsFail(t *testing.T) {
	calls := 0
	cfg := Config{
		MaxAttempts:    2,
		BaseDelay:     1 * time.Millisecond,
		MaxDelay:      10 * time.Millisecond,
		IsRecoverable: defaultIsRecoverable,
	}
	err := Do(context.Background(), cfg, func() error {
		calls++
		return constant.NewFutuError(constant.ErrCodeNetworkError, "TestDo_AllAttemptsFail", "always fails")
	})
	if err == nil {
		t.Fatal("expected error after all attempts exhausted")
	}
	if calls != 2 {
		t.Errorf("expected 2 calls, got %d", calls)
	}
}

func TestDo_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	calls := 0
err := Do(ctx, Config{
		MaxAttempts:    100,
		BaseDelay:     500 * time.Millisecond,
		MaxDelay:      500 * time.Millisecond,
		IsRecoverable: defaultIsRecoverable,
	}, func() error {
		calls++
		return constant.NewFutuError(constant.ErrCodeNetworkError, "TestDo_ZeroMaxAttempts", "fail")
	})
	if err != context.DeadlineExceeded {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
	if calls > 2 {
		t.Errorf("expected at most 2 calls before context expired, got %d", calls)
	}
}

func TestDo_ContextCancelBeforeDelay(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately
	calls := 0
	err := Do(ctx, Config{
		MaxAttempts:    3,
		BaseDelay:     1 * time.Second,
		MaxDelay:      1 * time.Second,
		IsRecoverable: defaultIsRecoverable,
	}, func() error {
		calls++
		return constant.NewFutuError(constant.ErrCodeNetworkError, "TestDo_ContextCancelBeforeDelay", "fail")
	})
	if err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", err)
	}
	if calls != 1 {
		t.Errorf("expected 1 call, got %d", calls)
	}
}

func TestDoWithResult_Success(t *testing.T) {
	result, err := DoWithResult(context.Background(), DefaultConfig(), func() (interface{}, error) {
		return "hello", nil
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "hello" {
		t.Errorf("expected 'hello', got %v", result)
	}
}

func TestDoWithResult_RetryWithResult(t *testing.T) {
	calls := 0
	result, err := DoWithResult(context.Background(), Config{
		MaxAttempts:    3,
		BaseDelay:     1 * time.Millisecond,
		MaxDelay:      10 * time.Millisecond,
		IsRecoverable: defaultIsRecoverable,
	}, func() (interface{}, error) {
		calls++
		if calls < 2 {
			return nil, constant.NewFutuError(constant.ErrCodeNetworkError, "TestDoWithResult_RetryWithResult", "fail")
		}
		return 42, nil
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 42 {
		t.Errorf("expected 42, got %v", result)
	}
	if calls != 2 {
		t.Errorf("expected 2 calls, got %d", calls)
	}
}

func TestDoWithResult_Unrecoverable(t *testing.T) {
_, err := DoWithResult(context.Background(), Config{
		MaxAttempts:    3,
		BaseDelay:     1 * time.Millisecond,
		MaxDelay:      10 * time.Millisecond,
		IsRecoverable: defaultIsRecoverable,
	}, func() (interface{}, error) {
		return nil, constant.NewFutuError(constant.ErrCodeInvalidParams, "TestDoWithResult_Unrecoverable", "bad request")
	})
	if err == nil {
		t.Fatal("expected error for unrecoverable failure")
	}
}

func TestDoDelay_ExponentialBackoff(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Jitter = false
	d1 := cfg.delay(0) // 2^0 * 500ms = 500ms
	d2 := cfg.delay(1) // 2^1 * 500ms = 1000ms
	d3 := cfg.delay(2) // 2^2 * 500ms = 2000ms
	if d1 != 500*time.Millisecond {
		t.Errorf("attempt 0 delay: got %v, want 500ms", d1)
	}
	if d2 != 1*time.Second {
		t.Errorf("attempt 1 delay: got %v, want 1s", d2)
	}
	if d3 != 2*time.Second {
		t.Errorf("attempt 2 delay: got %v, want 2s", d3)
	}
}

func TestDoDelay_CappedAtMax(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Jitter = false
	cfg.MaxDelay = 1 * time.Second
	d := cfg.delay(10) // 2^10 * 500ms = 512s, but capped at 1s
	if d != 1*time.Second {
		t.Errorf("expected delay capped at 1s, got %v", d)
	}
}

func TestDoDelay_JitterRange(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Jitter = true
	cfg.BaseDelay = 100 * time.Millisecond
	cfg.MaxDelay = 1 * time.Second
	// Jitter factor: 0.8 + 0.4*rand.Float64() => range [0.8, 1.2)
	// For attempt 0: delay = 100ms * jitter, range [80ms, 120ms)
	seenBelow := false
	seenAbove := false
	for i := 0; i < 100; i++ {
		d := cfg.delay(0)
		if d < 80*time.Millisecond {
			seenBelow = true
		}
		if d >= 120*time.Millisecond {
			seenAbove = true
		}
	}
	if seenBelow {
		t.Error("jitter produced delay below 80% of base")
	}
	if seenAbove {
		t.Error("jitter produced delay at or above 120% of base")
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.MaxAttempts != 3 {
		t.Errorf("MaxAttempts: got %d, want 3", cfg.MaxAttempts)
	}
	if cfg.BaseDelay != 500*time.Millisecond {
		t.Errorf("BaseDelay: got %v, want 500ms", cfg.BaseDelay)
	}
	if cfg.MaxDelay != 10*time.Second {
		t.Errorf("MaxDelay: got %v, want 10s", cfg.MaxDelay)
	}
	if !cfg.Jitter {
		t.Error("Jitter should be true by default")
	}
	if cfg.IsRecoverable == nil {
		t.Error("IsRecoverable should not be nil")
	}
}

func TestDefaultIsRecoverable(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		recoverable bool
	}{
		{"timeout", constant.NewFutuError(constant.ErrCodeTimeout, "TestDefaultIsRecoverable", "timeout"), true},
		{"disconnected", constant.NewFutuError(constant.ErrCodeDisconnected, "TestDefaultIsRecoverable", "disconnected"), true},
		{"network", constant.NewFutuError(constant.ErrCodeNetworkError, "TestDefaultIsRecoverable", "network"), true},
		{"invalid_params", constant.NewFutuError(constant.ErrCodeInvalidParams, "TestDefaultIsRecoverable", "bad"), false},
		{"trading", constant.NewFutuError(constant.ErrCodeOrderRejected, "TestDefaultIsRecoverable", "rejected"), false},
		{"plain error", errors.New("something else"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := defaultIsRecoverable(tt.err)
			if got != tt.recoverable {
				t.Errorf("defaultIsRecoverable(%v) = %v, want %v", tt.err, got, tt.recoverable)
			}
		})
	}
}

func TestDo_ZeroMaxAttempts(t *testing.T) {
	cfg := DefaultConfig()
	cfg.MaxAttempts = 0
	calls := 0
	err := Do(context.Background(), cfg, func() error {
		calls++
		return nil
	})
	if err != nil {
		t.Errorf("expected no error with 0 attempts and nil fn error, got %v", err)
	}
	if calls != 0 {
		t.Errorf("expected 0 calls with MaxAttempts=0, got %d", calls)
	}
}