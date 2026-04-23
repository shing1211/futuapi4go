// Package breaker implements the Circuit Breaker pattern for the futuapi4go SDK.
//
// A circuit breaker prevents cascading failures by stopping requests to a
// failing service. When the circuit is "open", calls return immediately with
// an error. After a cooldown period, the circuit moves to "half-open" and
// allows one test call. If it succeeds, the circuit closes; if it fails,
// the circuit opens again.
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/breaker"
//
//	cb := breaker.New(breaker.WithThreshold(5))
//
//	result, err := cb.Do(func() (interface{}, error) {
//	    return qot.GetBasicQot(cli, market, code)
//	})
//
//	if err == breaker.ErrOpen {
//	    log.Println("circuit open, service unavailable")
//	}
//
// The breaker is safe for concurrent use.
//
// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package breaker

import (
	"errors"
	"sync"
	"time"
)

var ErrOpen = errors.New("circuit breaker is open")

type State int

const (
	StateClosed   State = 0
	StateOpen     State = 1
	StateHalfOpen State = 2
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateOpen:
		return "open"
	case StateHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}

type Breaker struct {
	mu       sync.RWMutex
	state    State
	failures int

	threshold    int
	cooldown     time.Duration
	halfOpenMax int

	lastFailure time.Time
	openTime   time.Time

	onOpen   func()
	onClose  func()
	onChange func(State, State)
}

type Config struct {
	Threshold    int
	Cooldown     time.Duration
	HalfOpenMax int
	OnOpen      func()
	OnClose     func()
	OnChange    func(from, to State)
}

type Option func(*Config)

func WithThreshold(n int) Option {
	return func(c *Config) { c.Threshold = n }
}

func WithCooldown(d time.Duration) Option {
	return func(c *Config) { c.Cooldown = d }
}

func WithHalfOpenMax(n int) Option {
	return func(c *Config) { c.HalfOpenMax = n }
}

func WithOnOpen(fn func()) Option {
	return func(c *Config) { c.OnOpen = fn }
}

func WithOnClose(fn func()) Option {
	return func(c *Config) { c.OnClose = fn }
}

func WithOnChange(fn func(from, to State)) Option {
	return func(c *Config) { c.OnChange = fn }
}

func New(opts ...Option) *Breaker {
	cfg := Config{
		Threshold:    5,
		Cooldown:    30 * time.Second,
		HalfOpenMax: 1,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return &Breaker{
		state:        StateClosed,
		threshold:    cfg.Threshold,
		cooldown:     cfg.Cooldown,
		halfOpenMax:  cfg.HalfOpenMax,
		onOpen:       cfg.OnOpen,
		onClose:      cfg.OnClose,
		onChange:     cfg.OnChange,
	}
}

func (b *Breaker) State() State {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.state
}

func (b *Breaker) Failures() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.failures
}

func (b *Breaker) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	switch b.state {
	case StateClosed:
		return true

	case StateOpen:
		if time.Since(b.openTime) >= b.cooldown {
			b.transitionTo(StateHalfOpen)
			return true
		}
		return false

	case StateHalfOpen:
		return true

	default:
		return false
	}
}

func (b *Breaker) RecordSuccess() {
	b.mu.Lock()
	defer b.mu.Unlock()

	switch b.state {
	case StateClosed:
		if b.failures > 0 {
			b.failures--
		}

	case StateHalfOpen:
		b.transitionTo(StateClosed)

	case StateOpen:
	}
}

func (b *Breaker) RecordFailure() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.lastFailure = time.Now()

	switch b.state {
	case StateClosed:
		b.failures++
		if b.failures >= b.threshold {
			b.transitionTo(StateOpen)
		}

	case StateHalfOpen:
		b.transitionTo(StateOpen)

	case StateOpen:
	}
}

func (b *Breaker) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.transitionTo(StateClosed)
	b.failures = 0
}

func (b *Breaker) transitionTo(newState State) {
	if b.state == newState {
		return
	}
	oldState := b.state
	b.state = newState
	if newState == StateOpen {
		b.openTime = time.Now()
	}
	if newState == StateClosed {
		b.failures = 0
		if b.onClose != nil {
			b.onClose()
		}
	}
	if newState == StateOpen && b.onOpen != nil {
		b.onOpen()
	}
	if b.onChange != nil && oldState != newState {
		b.onChange(oldState, newState)
	}
}

func (b *Breaker) Do(fn func() (interface{}, error)) (interface{}, error) {
	if !b.Allow() {
		return nil, ErrOpen
	}

	result, err := fn()
	if err != nil {
		b.RecordFailure()
	} else {
		b.RecordSuccess()
	}
	return result, err
}

func (b *Breaker) DoVoid(fn func() error) error {
	if !b.Allow() {
		return ErrOpen
	}

	err := fn()
	if err != nil {
		b.RecordFailure()
	} else {
		b.RecordSuccess()
	}
	return err
}

func (b *Breaker) Stats() Stats {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return Stats{
		State:         b.state,
		Failures:      b.failures,
		Threshold:     b.threshold,
		CooldownSecs: int(b.cooldown.Seconds()),
		LastFailure:   b.lastFailure,
		OpenSince:     b.openTime,
	}
}

type Stats struct {
	State         State
	Failures      int
	Threshold     int
	CooldownSecs int
	LastFailure   time.Time
	OpenSince     time.Time
}
