package retry

import (
	"context"
	"math"
	"math/rand"
	"time"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

type Config struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
	Jitter      bool
	IsRecoverable func(error) bool
}

func DefaultConfig() Config {
	return Config{
		MaxAttempts:    3,
		BaseDelay:     500 * time.Millisecond,
		MaxDelay:      10 * time.Second,
		Jitter:        true,
		IsRecoverable: defaultIsRecoverable,
	}
}

func defaultIsRecoverable(err error) bool {
	cat := constant.CategoryOf(err)
	switch cat {
	case constant.CategoryTimeout, constant.CategoryConnection:
		return true
	default:
		return false
	}
}

func (c *Config) delay(attempt int) time.Duration {
	d := time.Duration(math.Pow(2, float64(attempt))) * c.BaseDelay
	if d > c.MaxDelay {
		d = c.MaxDelay
	}
	if c.Jitter {
		d = time.Duration(float64(d) * (0.8 + 0.4*rand.Float64()))
	}
	return d
}

func Do(ctx context.Context, cfg Config, fn func() error) error {
	var lastErr error
	for attempt := 0; attempt < cfg.MaxAttempts; attempt++ {
		if attempt > 0 {
			delay := cfg.delay(attempt - 1)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(delay):
			}
		}

		err := fn()
		if err == nil {
			return nil
		}
		lastErr = err

		if !cfg.IsRecoverable(err) {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}
	return lastErr
}

type RetryableFunc func() (interface{}, error)

func DoWithResult(ctx context.Context, cfg Config, fn RetryableFunc) (interface{}, error) {
	var lastErr error
	for attempt := 0; attempt < cfg.MaxAttempts; attempt++ {
		if attempt > 0 {
			delay := cfg.delay(attempt - 1)
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(delay):
			}
		}

		result, err := fn()
		if err == nil {
			return result, nil
		}
		lastErr = err

		if !cfg.IsRecoverable(err) {
			return nil, err
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}
	return nil, lastErr
}
