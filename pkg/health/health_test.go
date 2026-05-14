package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewChecker(t *testing.T) {
	c := NewChecker()
	if c == nil {
		t.Fatal("NewChecker returned nil")
	}
	results := c.Check()
	if len(results) != 0 {
		t.Errorf("expected 0 results from fresh checker, got %d", len(results))
	}
}

func TestChecker_Register(t *testing.T) {
	c := NewChecker()
	c.Register("test", func() Status { return StatusHealthy })
	results := c.Check()
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Name != "test" {
		t.Errorf("expected name 'test', got %q", results[0].Name)
	}
	if results[0].Status != StatusHealthy {
		t.Errorf("expected StatusHealthy, got %v", results[0].Status)
	}
}

func TestChecker_RegisterMultiple(t *testing.T) {
	c := NewChecker()
	c.Register("a", func() Status { return StatusHealthy })
	c.Register("b", func() Status { return StatusDegraded })
	c.Register("c", func() Status { return StatusUnhealthy })
	results := c.Check()
	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
}

func TestChecker_IsHealthy(t *testing.T) {
	c := NewChecker()
	c.Register("ok", func() Status { return StatusHealthy })
	if !c.IsHealthy() {
		t.Error("expected IsHealthy=true when all checks healthy")
	}
	c.Register("bad", func() Status { return StatusUnhealthy })
	if c.IsHealthy() {
		t.Error("expected IsHealthy=false when a check is unhealthy")
	}
}

func TestChecker_IsHealthy_Empty(t *testing.T) {
	c := NewChecker()
	// No checks registered — vacuously healthy
	if !c.IsHealthy() {
		t.Error("expected IsHealthy=true with no checks")
	}
}

func TestChecker_IsReady(t *testing.T) {
	c := NewChecker()
	c.Register("ok", func() Status { return StatusHealthy })
	if !c.IsReady() {
		t.Error("expected IsReady=true when all checks healthy")
	}
	c.Register("degraded", func() Status { return StatusDegraded })
	if c.IsReady() {
		t.Error("expected IsReady=false when a check is degraded")
	}
}

func TestChecker_IsReady_Unhealthy(t *testing.T) {
	c := NewChecker()
	c.Register("bad", func() Status { return StatusUnhealthy })
	if c.IsReady() {
		t.Error("expected IsReady=false when a check is unhealthy")
	}
}

func TestChecker_IsReady_Empty(t *testing.T) {
	c := NewChecker()
	if !c.IsReady() {
		t.Error("expected IsReady=true with no checks")
	}
}

func TestChecker_Check_ReturnsCopy(t *testing.T) {
	c := NewChecker()
	c.Register("test", func() Status { return StatusHealthy })
	results := c.Check()
	// Modifying returned slice should not affect internal state
	results[0].Status = StatusUnhealthy
	results2 := c.Check()
	if results2[0].Status != StatusHealthy {
		t.Error("modifying returned result should not affect checker")
	}
}

func TestServeHTTP_Healthz(t *testing.T) {
	c := NewChecker()
	c.Register("ok", func() Status { return StatusHealthy })
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != "ok" {
		t.Errorf("expected body 'ok', got %q", w.Body.String())
	}
}

func TestServeHTTP_Healthz_Unhealthy(t *testing.T) {
	c := NewChecker()
	c.Register("bad", func() Status { return StatusUnhealthy })
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)
	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("expected 503, got %d", w.Code)
	}
	if w.Body.String() != "unhealthy" {
		t.Errorf("expected body 'unhealthy', got %q", w.Body.String())
	}
}

func TestServeHTTP_Readyz(t *testing.T) {
	c := NewChecker()
	c.Register("ok", func() Status { return StatusHealthy })
	req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != "ok" {
		t.Errorf("expected body 'ok', got %q", w.Body.String())
	}
}

func TestServeHTTP_Readyz_NotReady(t *testing.T) {
	c := NewChecker()
	c.Register("degraded", func() Status { return StatusDegraded })
	req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)
	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("expected 503, got %d", w.Code)
	}
	if w.Body.String() != "not ready" {
		t.Errorf("expected body 'not ready', got %q", w.Body.String())
	}
}

func TestServeHTTP_NotFound(t *testing.T) {
	c := NewChecker()
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}

func TestStatusValues(t *testing.T) {
	if StatusHealthy == StatusDegraded {
		t.Error("StatusHealthy should differ from StatusDegraded")
	}
	if StatusDegraded == StatusUnhealthy {
		t.Error("StatusDegraded should differ from StatusUnhealthy")
	}
	if StatusHealthy == StatusUnhealthy {
		t.Error("StatusHealthy should differ from StatusUnhealthy")
	}
}