package ratelimit

import (
	"context"
	"sync"
	"time"
)

type Limiter struct {
	mu       sync.Mutex
	tokens   float64
	maxRate  float64
	capacity float64
	lastTime time.Time
	mode     Mode
}

type Mode int

const (
	ModeReject Mode = iota
	ModeWait
)

func NewLimiter(rate float64, capacity float64, mode Mode) *Limiter {
	return &Limiter{
		tokens:   capacity,
		maxRate:  rate,
		capacity: capacity,
		lastTime: time.Now(),
		mode:     mode,
	}
}

func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(l.lastTime).Seconds()
	l.tokens += elapsed * l.maxRate
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastTime = now

	if l.tokens >= 1 {
		l.tokens--
		return true
	}
	return false
}

func (l *Limiter) Wait(ctx context.Context) error {
	for {
		if l.Allow() {
			return nil
		}
		if l.mode == ModeReject {
			return ErrRateLimited
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(10 * time.Millisecond):
		}
	}
}

var ErrRateLimited = &RateLimitError{}

type RateLimitError struct{}

func (e *RateLimitError) Error() string { return "rate limit exceeded" }

type ProtoLimiter struct {
	mu      sync.RWMutex
	global  *Limiter
	perProto map[uint32]*Limiter
}

func NewProtoLimiter(globalRate, globalCapacity float64, mode Mode) *ProtoLimiter {
	return &ProtoLimiter{
		global:   NewLimiter(globalRate, globalCapacity, mode),
		perProto: make(map[uint32]*Limiter),
	}
}

func (pl *ProtoLimiter) SetProtoLimit(protoID uint32, rate, capacity float64, mode Mode) {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	pl.perProto[protoID] = NewLimiter(rate, capacity, mode)
}

func (pl *ProtoLimiter) Allow(protoID uint32) bool {
	if !pl.global.Allow() {
		return false
	}
	pl.mu.RLock()
	limiter, ok := pl.perProto[protoID]
	pl.mu.RUnlock()
	if ok {
		return limiter.Allow()
	}
	return true
}

func (pl *ProtoLimiter) Wait(ctx context.Context, protoID uint32) error {
	if err := pl.global.Wait(ctx); err != nil {
		return err
	}
	pl.mu.RLock()
	limiter, ok := pl.perProto[protoID]
	pl.mu.RUnlock()
	if ok {
		return limiter.Wait(ctx)
	}
	return nil
}
