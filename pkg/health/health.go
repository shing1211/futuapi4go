package health

import (
	"net/http"
	"sync"
)

type Status string

const (
	StatusHealthy   Status = "healthy"
	StatusDegraded  Status = "degraded"
	StatusUnhealthy Status = "unhealthy"
)

type CheckFunc func() Status

type Checker struct {
	mu     sync.RWMutex
	checks map[string]CheckFunc
}

func NewChecker() *Checker {
	return &Checker{checks: make(map[string]CheckFunc)}
}

func (h *Checker) Register(name string, fn CheckFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.checks[name] = fn
}

type Result struct {
	Name   string
	Status Status
}

func (h *Checker) Check() []Result {
	h.mu.RLock()
	defer h.mu.RUnlock()
	results := make([]Result, 0, len(h.checks))
	for name, fn := range h.checks {
		results = append(results, Result{Name: name, Status: fn()})
	}
	return results
}

func (h *Checker) IsHealthy() bool {
	for _, r := range h.Check() {
		if r.Status == StatusUnhealthy {
			return false
		}
	}
	return true
}

func (h *Checker) IsReady() bool {
	for _, r := range h.Check() {
		if r.Status != StatusHealthy {
			return false
		}
	}
	return true
}

func (h *Checker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/healthz" {
		if h.IsHealthy() {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("unhealthy"))
		}
		return
	}
	if r.URL.Path == "/readyz" {
		if h.IsReady() {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("not ready"))
		}
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
