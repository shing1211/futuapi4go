package history

import (
	"sync"
	"testing"
	"time"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

func TestDownloadProgress_Percent(t *testing.T) {
	tests := []struct {
		name     string
		downloaded int
		total    int
		want     float64
	}{
		{"zero total", 5, 0, 0},
		{"half", 50, 100, 50.0},
		{"full", 100, 100, 100.0},
		{"over", 150, 100, 150.0},
		{"zero both", 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tracker := NewProgressTracker(tt.total, func(p DownloadProgress) {})
			tracker.Add(tt.downloaded)
			if got := tracker.Percent(); got != tt.want {
				t.Errorf("Percent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgressTracker_Add(t *testing.T) {
	var lastProgress DownloadProgress
	var mu sync.Mutex
	tracker := NewProgressTracker(100, func(p DownloadProgress) {
		mu.Lock()
		defer mu.Unlock()
		lastProgress = p
	})
	tracker.Add(10)
	mu.Lock()
	if lastProgress.Downloaded != 10 {
		t.Errorf("expected Downloaded=10, got %d", lastProgress.Downloaded)
	}
	if lastProgress.Total != 100 {
		t.Errorf("expected Total=100, got %d", lastProgress.Total)
	}
	mu.Unlock()
	tracker.Add(20)
	mu.Lock()
	if lastProgress.Downloaded != 30 {
		t.Errorf("expected Downloaded=30, got %d", lastProgress.Downloaded)
	}
	mu.Unlock()
}

func TestProgressTracker_Percent(t *testing.T) {
	tracker := NewProgressTracker(200, func(p DownloadProgress) {})
	tracker.Add(50)
	p := tracker.Percent()
	if p != 25.0 {
		t.Errorf("Percent() = %v, want 25.0", p)
	}
}

func TestProgressTracker_PercentZeroTotal(t *testing.T) {
	tracker := NewProgressTracker(0, func(p DownloadProgress) {})
	tracker.Add(100)
	if p := tracker.Percent(); p != 0 {
		t.Errorf("Percent() with total=0 should be 0, got %v", p)
	}
}

func TestProgressTracker_AddConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	tracker := NewProgressTracker(1000, func(p DownloadProgress) {})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				tracker.Add(1)
			}
		}()
	}
	wg.Wait()
	if p := tracker.Percent(); p != 100.0 {
		t.Errorf("expected 100%%, got %v%%", p)
	}
}

func TestKLineRequest_Defaults(t *testing.T) {
	req := KLineRequest{
		Code:     "00700",
		Market:   constant.Market_HK,
		KLType:   KLType_1Day,
		StartDate: "2020-01-01",
		EndDate:  "2024-01-01",
	}
	if req.MaxPerPage != 0 {
		t.Errorf("default MaxPerPage should be 0, got %d", req.MaxPerPage)
	}
}

func TestNewDownloader_Defaults(t *testing.T) {
	// Nil client would panic on actual use, but we can test option defaults
	_ = NewDownloader(nil)
	// If we got here without panic, the constructor works with nil client
	// (actual API calls would fail, but that's expected)
}

func TestDownloader_WithMaxRetries(t *testing.T) {
	d := NewDownloader(nil, WithMaxRetries(5))
	if d.maxRetries != 5 {
		t.Errorf("expected maxRetries=5, got %d", d.maxRetries)
	}
}

func TestDownloader_WithPageDelay(t *testing.T) {
	delay := 200 * time.Millisecond
	d := NewDownloader(nil, WithPageDelay(delay))
	if d.pageDelay != delay {
		t.Errorf("expected pageDelay=%v, got %v", delay, d.pageDelay)
	}
}

func TestDownloader_WithProgress(t *testing.T) {
	called := false
	var mu sync.Mutex
	cb := func(p DownloadProgress) {
		mu.Lock()
		defer mu.Unlock()
		called = true
	}
	d := NewDownloader(nil, WithProgress(cb))
	d.progress(DownloadProgress{Downloaded: 10, Total: 100})
	mu.Lock()
	defer mu.Unlock()
	if !called {
		t.Error("progress callback was not called")
	}
}

func TestNewProgressTracker(t *testing.T) {
	tracker := NewProgressTracker(500, func(p DownloadProgress) {})
	if tracker.total != 500 {
		t.Errorf("expected total=500, got %d", tracker.total)
	}
	if tracker.downloaded != 0 {
		t.Errorf("expected downloaded=0, got %d", tracker.downloaded)
	}
}

func TestConcurrentDownloader_New(t *testing.T) {
	cd := NewConcurrentDownloader(nil)
	if cd.workers != 4 {
		t.Errorf("expected workers=4, got %d", cd.workers)
	}
	if cd.pageDelay != 50*time.Millisecond {
		t.Errorf("expected pageDelay=50ms, got %v", cd.pageDelay)
	}
}

func TestConcurrentDownloader_WithWorkers(t *testing.T) {
	cd := NewConcurrentDownloader(nil, WithWorkers(8))
	if cd.workers != 8 {
		t.Errorf("expected workers=8, got %d", cd.workers)
	}
}

// Note: DownloadKLine, DownloadWithStats, and DownloadMultiple require
// a real client connection and are tested via integration tests.
// These are unit tests for the non-I/O logic only.