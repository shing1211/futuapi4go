package metrics

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestRecordConnection(t *testing.T) {
	RecordConnection("test")
	// No panic is success; prometheus metrics are global
	RecordDisconnect("test")
}

func TestRecordReconnect(t *testing.T) {
	RecordReconnect("test_reason")
}

func TestRecordAPICall(t *testing.T) {
	RecordAPICall("test_proto", "success", 50*time.Millisecond)
	RecordAPICall("test_proto", "error", 100*time.Millisecond)
}

func TestRecordPushMessage(t *testing.T) {
	RecordPushMessage("test_type")
}

func TestRecordOpenDUp(t *testing.T) {
	RecordOpenDUp(true)
	RecordOpenDUp(false)
}

func TestRecordRateLimited(t *testing.T) {
	RecordRateLimited("test_proto")
}

func TestRecordRetry(t *testing.T) {
	RecordRetry("test_proto", "1")
}

func TestRecordBreakerState(t *testing.T) {
	RecordBreakerState("test", 0.5)
}

func TestAPICallTracker(t *testing.T) {
	tracker := StartAPITracking("test_proto")
	if tracker == nil {
		t.Fatal("StartAPITracking returned nil")
	}
	if tracker.finished {
		t.Error("tracker should not be finished on creation")
	}
	// Calling End should succeed
	tracker.End(true)
	if !tracker.finished {
		t.Error("tracker should be finished after End(true)")
	}
	// Calling End again should be a no-op
	tracker.End(false)
}

func TestAPICallTracker_DoubleEnd(t *testing.T) {
	tracker := StartAPITracking("test_proto")
	tracker.End(true)
	// Second call should be ignored (no double-counting)
	tracker.End(false)
}

func TestInit_Idempotent(t *testing.T) {
	Init()
	Init() // should not panic
}

func TestInitWithServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/metrics" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer srv.Close()
	// Just ensure it doesn't panic; the real InitWithServer spawns its own goroutine
	Init()
}

func TestHandler(t *testing.T) {
	h := Handler()
	if h == nil {
		t.Fatal("Handler returned nil")
	}
}

func TestConcurrentRecording(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			RecordConnection("concurrent_test")
			RecordAPICall("proto", "success", time.Duration(idx)*time.Millisecond)
			RecordDisconnect("concurrent_test")
		}(i)
	}
	wg.Wait()
	// Success = no race conditions (detected by -race flag)
}