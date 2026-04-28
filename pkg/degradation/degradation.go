package degradation

import (
	"sync"
)

type Component string

const (
	ComponentPush      Component = "push"
	ComponentHistory   Component = "history"
	ComponentPool      Component = "pool"
	ComponentBreaker   Component = "breaker"
	ComponentRateLimit Component = "ratelimit"
)

type Event struct {
	Component Component
	Level     Level
	Message   string
}

type Level int

const (
	LevelNormal Level = iota
	LevelDegraded
	LevelFailed
)

type Watcher func(evt Event)

type Manager struct {
	mu       sync.RWMutex
	status   map[Component]Level
	watchers []Watcher
}

func NewManager() *Manager {
	return &Manager{
		status: make(map[Component]Level),
	}
}

func (m *Manager) AddWatcher(w Watcher) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.watchers = append(m.watchers, w)
}

func (m *Manager) SetStatus(component Component, level Level, msg string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	prev := m.status[component]
	m.status[component] = level
	if prev != level {
		evt := Event{Component: component, Level: level, Message: msg}
		for _, w := range m.watchers {
			w(evt)
		}
	}
}

func (m *Manager) GetStatus(component Component) Level {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.status[component]
}

func (m *Manager) IsDegraded() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, level := range m.status {
		if level >= LevelDegraded {
			return true
		}
	}
	return false
}

func (m *Manager) AllStatus() map[Component]Level {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make(map[Component]Level, len(m.status))
	for k, v := range m.status {
		result[k] = v
	}
	return result
}
