package breaker

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestBreakerInitialState(t *testing.T) {
	cb := New()
	if got := cb.State(); got != StateClosed {
		t.Errorf("initial state = %v, want StateClosed", got)
	}
}

func TestBreakerAllowsInClosedState(t *testing.T) {
	cb := New()
	if !cb.Allow() {
		t.Error("Allow() = false, want true in closed state")
	}
}

func TestBreakerRecordSuccess(t *testing.T) {
	cb := New()
	cb.RecordSuccess()
	cb.RecordSuccess()
	if got := cb.State(); got != StateClosed {
		t.Errorf("state = %v after successes, want StateClosed", got)
	}
}

func TestBreakerRecordFailureBelowThreshold(t *testing.T) {
	cb := New(WithThreshold(3))
	for i := 0; i < 2; i++ {
		cb.RecordFailure()
	}
	if got := cb.State(); got != StateClosed {
		t.Errorf("state = %v after 2 failures (threshold=3), want StateClosed", got)
	}
}

func TestBreakerOpensAfterThreshold(t *testing.T) {
	cb := New(WithThreshold(3))
	for i := 0; i < 3; i++ {
		cb.RecordFailure()
	}
	if got := cb.State(); got != StateOpen {
		t.Errorf("state = %v after 3 failures (threshold=3), want StateOpen", got)
	}
}

func TestBreakerBlocksInOpenState(t *testing.T) {
	cb := New(WithThreshold(2), WithCooldown(10*time.Second))
	for i := 0; i < 2; i++ {
		cb.RecordFailure()
	}
	if cb.Allow() {
		t.Error("Allow() = true in open state, want false")
	}
}

func TestBreakerHalfOpenAfterCooldown(t *testing.T) {
	cb := New(WithThreshold(1), WithCooldown(50*time.Millisecond))
	cb.RecordFailure()
	if got := cb.State(); got != StateOpen {
		t.Errorf("state = %v, want StateOpen", got)
	}

	time.Sleep(60 * time.Millisecond)
	if !cb.Allow() {
		t.Error("Allow() = false after cooldown, want true (half-open)")
	}
	if got := cb.State(); got != StateHalfOpen {
		t.Errorf("state = %v after cooldown, want StateHalfOpen", got)
	}
}

func TestBreakerClosesOnSuccessInHalfOpen(t *testing.T) {
	cb := New(WithThreshold(1), WithCooldown(1*time.Millisecond))
	cb.RecordFailure()
	time.Sleep(2 * time.Millisecond)
	cb.Allow()
	cb.RecordSuccess()
	if got := cb.State(); got != StateClosed {
		t.Errorf("state = %v after success in half-open, want StateClosed", got)
	}
}

func TestBreakerReopensOnFailureInHalfOpen(t *testing.T) {
	cb := New(WithThreshold(1), WithCooldown(1*time.Millisecond))
	cb.RecordFailure()
	time.Sleep(2 * time.Millisecond)
	cb.Allow()
	cb.RecordFailure()
	if got := cb.State(); got != StateOpen {
		t.Errorf("state = %v after failure in half-open, want StateOpen", got)
	}
}

func TestBreakerDoExecutesInClosedState(t *testing.T) {
	cb := New()
	result, err := cb.Do(func() (interface{}, error) {
		return "success", nil
	})
	if err != nil {
		t.Errorf("Do() returned error: %v", err)
	}
	if result != "success" {
		t.Errorf("Do() result = %v, want success", result)
	}
	if got := cb.State(); got != StateClosed {
		t.Errorf("state = %v after successful Do, want StateClosed", got)
	}
}

func TestBreakerDoRecordsFailure(t *testing.T) {
	cb := New(WithThreshold(3))
	wantErr := errors.New("test error")
	for i := 0; i < 2; i++ {
		cb.Do(func() (interface{}, error) { return nil, wantErr })
	}
	cb.Do(func() (interface{}, error) { return nil, wantErr })
	if got := cb.State(); got != StateOpen {
		t.Errorf("state = %v after 3 Do failures, want StateOpen", got)
	}
}

func TestBreakerDoReturnsErrOpen(t *testing.T) {
	cb := New(WithThreshold(1))
	cb.Do(func() (interface{}, error) { return nil, errors.New("fail") })
	_, err := cb.Do(func() (interface{}, error) { return "should not run", nil })
	if err != ErrOpen {
		t.Errorf("Do() after open = %v, want ErrOpen", err)
	}
}

func TestBreakerDoVoidSuccess(t *testing.T) {
	cb := New()
	err := cb.DoVoid(func() error { return nil })
	if err != nil {
		t.Errorf("DoVoid() returned error: %v", err)
	}
}

func TestBreakerDoVoidFailure(t *testing.T) {
	cb := New(WithThreshold(2))
	errFail := errors.New("test")
	cb.DoVoid(func() error { return errFail })
	cb.DoVoid(func() error { return errFail })
	if got := cb.State(); got != StateOpen {
		t.Errorf("state = %v, want StateOpen", got)
	}
}

func TestBreakerReset(t *testing.T) {
	cb := New(WithThreshold(2))
	cb.RecordFailure()
	cb.RecordFailure()
	if got := cb.State(); got != StateOpen {
		t.Errorf("state = %v before reset, want StateOpen", got)
	}
	cb.Reset()
	if got := cb.State(); got != StateClosed {
		t.Errorf("state = %v after reset, want StateClosed", got)
	}
	if got := cb.Failures(); got != 0 {
		t.Errorf("failures = %d after reset, want 0", got)
	}
}

func TestBreakerOnChangeCallback(t *testing.T) {
	var changes []string
	cb := New(WithThreshold(1), WithOnChange(func(from, to State) {
		changes = append(changes, from.String()+"->"+to.String())
	}))
	cb.RecordFailure()
	if len(changes) != 1 || changes[0] != "closed->open" {
		t.Errorf("onChange calls = %v, want [closed->open]", changes)
	}
}

func TestBreakerStats(t *testing.T) {
	cb := New(WithThreshold(5), WithCooldown(30*time.Second))
	stats := cb.Stats()
	if stats.State != StateClosed {
		t.Errorf("stats.State = %v, want StateClosed", stats.State)
	}
	if stats.Threshold != 5 {
		t.Errorf("stats.Threshold = %d, want 5", stats.Threshold)
	}
	if stats.CooldownSecs != 30 {
		t.Errorf("stats.CooldownSecs = %d, want 30", stats.CooldownSecs)
	}
}

func TestBreakerOnOpenCallback(t *testing.T) {
	var called bool
	cb := New(WithThreshold(1), WithOnOpen(func() { called = true }))
	cb.RecordFailure()
	if !called {
		t.Error("onOpen callback not called when circuit opened")
	}
}

func TestBreakerOnCloseCallback(t *testing.T) {
	var called bool
	cb := New(WithThreshold(1), WithCooldown(1*time.Millisecond), WithOnClose(func() { called = true }))
	cb.RecordFailure()
	time.Sleep(2 * time.Millisecond)
	cb.Allow()
	cb.RecordSuccess()
	if !called {
		t.Error("onClose callback not called when circuit closed")
	}
}

func TestBreakerConcurrency(t *testing.T) {
	cb := New(WithThreshold(100))
	done := make(chan bool)
	for i := 0; i < 50; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				cb.RecordFailure()
			}
			done <- true
		}()
	}
	for i := 0; i < 50; i++ {
		<-done
	}
	if cb.State() != StateOpen {
		t.Errorf("state = %v after concurrent failures, want StateOpen", cb.State())
	}
}

func TestBreakerDoubleTransition(t *testing.T) {
	cb := New(WithThreshold(5))
	for i := 0; i < 10; i++ {
		cb.RecordFailure()
	}
	if cb.State() != StateOpen {
		t.Errorf("state = %v, want StateOpen", cb.State())
	}
}

func TestStateString(t *testing.T) {
	tests := []struct {
		s   State
		want string
	}{
		{StateClosed, "closed"},
		{StateOpen, "open"},
		{StateHalfOpen, "half-open"},
		{State(99), "unknown"},
	}
	for _, tt := range tests {
		if got := tt.s.String(); got != tt.want {
			t.Errorf("State(%d).String() = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestDoVoidBlocksWhenOpen(t *testing.T) {
	cb := New(WithThreshold(1))
	cb.DoVoid(func() error { return errors.New("fail") })
	err := cb.DoVoid(func() error { return nil })
	if err != ErrOpen {
		t.Errorf("DoVoid() = %v, want ErrOpen", err)
	}
}

func TestExecuteFailureMessage(t *testing.T) {
	if !strings.Contains(ErrOpen.Error(), "circuit breaker") {
		t.Errorf("ErrOpen.Error() = %q, want circuit breaker message", ErrOpen.Error())
	}
}
