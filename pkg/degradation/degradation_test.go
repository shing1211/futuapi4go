package degradation

import (
	"sync"
	"testing"
)

func TestManager_NewManager(t *testing.T) {
	m := NewManager()
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	// Initial status should be empty
	status := m.AllStatus()
	if len(status) != 0 {
		t.Errorf("expected empty status map, got %d entries", len(status))
	}
}

func TestManager_SetStatus(t *testing.T) {
	m := NewManager()
	m.SetStatus(ComponentPush, LevelDegraded, "push lag")
	if got := m.GetStatus(ComponentPush); got != LevelDegraded {
		t.Errorf("expected LevelDegraded, got %v", got)
	}
	if got := m.GetStatus(ComponentHistory); got != LevelNormal {
		t.Errorf("expected LevelNormal for unset component, got %v", got)
	}
}

func TestManager_SetStatus_FiresWatcher(t *testing.T) {
	m := NewManager()
	var firedEvents []Event
	var mu sync.Mutex
	m.AddWatcher(func(evt Event) {
		mu.Lock()
		defer mu.Unlock()
		firedEvents = append(firedEvents, evt)
	})
	m.SetStatus(ComponentPush, LevelDegraded, "push lag")
	m.SetStatus(ComponentPush, LevelDegraded, "push lag again") // same level, no fire
	m.SetStatus(ComponentPush, LevelNormal, "recovered")

	mu.Lock()
	defer mu.Unlock()
	if len(firedEvents) != 2 {
		t.Errorf("expected 2 watcher events, got %d: %v", len(firedEvents), firedEvents)
		return
	}
	if firedEvents[0].Component != ComponentPush || firedEvents[0].Level != LevelDegraded {
		t.Errorf("first event mismatch: %+v", firedEvents[0])
	}
	if firedEvents[1].Component != ComponentPush || firedEvents[1].Level != LevelNormal {
		t.Errorf("second event mismatch: %+v", firedEvents[1])
	}
}

func TestManager_SetStatus_NoFireOnSameLevel(t *testing.T) {
	m := NewManager()
	callCount := 0
	m.AddWatcher(func(Event) { callCount++ })
	m.SetStatus(ComponentPush, LevelDegraded, "msg1")
	m.SetStatus(ComponentPush, LevelDegraded, "msg2") // same level
	if callCount != 1 {
		t.Errorf("expected 1 watcher call, got %d", callCount)
	}
}

func TestManager_GetStatus(t *testing.T) {
	m := NewManager()
	// Unset component should return zero value (LevelNormal=0)
	if got := m.GetStatus(ComponentHistory); got != LevelNormal {
		t.Errorf("expected LevelNormal for unset component, got %v", got)
	}
	m.SetStatus(ComponentBreaker, LevelFailed, "breaker tripped")
	if got := m.GetStatus(ComponentBreaker); got != LevelFailed {
		t.Errorf("expected LevelFailed, got %v", got)
	}
}

func TestManager_IsDegraded(t *testing.T) {
	m := NewManager()
	if m.IsDegraded() {
		t.Error("fresh manager should not be degraded")
	}
	m.SetStatus(ComponentPush, LevelDegraded, "push lag")
	if !m.IsDegraded() {
		t.Error("manager should be degraded after setting LevelDegraded")
	}
	m.SetStatus(ComponentPush, LevelNormal, "recovered")
	if m.IsDegraded() {
		t.Error("manager should not be degraded after recovery")
	}
}

func TestManager_IsDegraded_LevelFailed(t *testing.T) {
	m := NewManager()
	m.SetStatus(ComponentBreaker, LevelFailed, "critical")
	if !m.IsDegraded() {
		t.Error("LevelFailed should count as degraded")
	}
}

func TestManager_AllStatus(t *testing.T) {
	m := NewManager()
	m.SetStatus(ComponentPush, LevelDegraded, "lag")
	m.SetStatus(ComponentHistory, LevelNormal, "ok")
	status := m.AllStatus()
	if len(status) != 2 {
		t.Errorf("expected 2 entries, got %d", len(status))
	}
	// Verify returned map is a copy (not the internal map)
	status[ComponentPush] = LevelNormal
	if m.GetStatus(ComponentPush) != LevelDegraded {
		t.Error("modifying returned map should not affect internal state")
	}
}

func TestManager_ConcurrentAccess(t *testing.T) {
	m := NewManager()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			comp := ComponentPush
			if idx%2 == 0 {
				comp = ComponentHistory
			}
			m.SetStatus(comp, LevelDegraded, "concurrent test")
			_ = m.GetStatus(comp)
			_ = m.IsDegraded()
			_ = m.AllStatus()
		}(i)
	}
	wg.Wait()
	// Success = no race conditions (detected by -race flag)
}

func TestManager_MultipleWatchers(t *testing.T) {
	m := NewManager()
	var counts [3]int
	var mu sync.Mutex
	for i := 0; i < 3; i++ {
		idx := i
		m.AddWatcher(func(Event) {
			mu.Lock()
			defer mu.Unlock()
			counts[idx]++
		})
	}
	m.SetStatus(ComponentRateLimit, LevelDegraded, "rate limited")
	mu.Lock()
	defer mu.Unlock()
	for i, c := range counts {
		if c != 1 {
			t.Errorf("watcher %d: expected 1 call, got %d", i, c)
		}
	}
}

func TestManager_NilWatcherNoPanic(t *testing.T) {
	m := NewManager()
	// Register nil watcher - should not panic
	m.AddWatcher(nil)
	m.SetStatus(ComponentPush, LevelDegraded, "test")
	// Should not panic when iterating watchers
}