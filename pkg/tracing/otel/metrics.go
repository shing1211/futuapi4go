package otel

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type OTelMeter struct {
	Connections   metric.Int64UpDownCounter
	Reconnects    metric.Int64Counter
	APICalls      metric.Int64Counter
	APILatency    metric.Float64Histogram
	PushMessages  metric.Int64Counter
	APIErrors     metric.Int64Counter
	RateLimited   metric.Int64Counter
	RetryAttempts metric.Int64Counter

	openDUpValue         atomic.Int64
	lastConnectTimeValue atomic.Int64
	breakerStates        sync.Map

	gaugeReg metric.Registration
}

func NewOTelMeter(opts ...Option) (*OTelMeter, error) {
	cfg := &config{}
	for _, o := range opts {
		o(cfg)
	}

	mp := cfg.meterProvider
	if mp == nil {
		mp = otel.GetMeterProvider()
	}

	meter := mp.Meter(
		"github.com/shing1211/futuapi4go",
		metric.WithInstrumentationVersion("0.7.0"),
	)

	m := &OTelMeter{}

	var err error

	m.Connections, err = meter.Int64UpDownCounter(
		"futuapi.connections",
		metric.WithDescription("Active connections by type"),
	)
	if err != nil {
		return nil, err
	}

	m.Reconnects, err = meter.Int64Counter(
		"futuapi.reconnects",
		metric.WithDescription("Reconnection attempts by reason"),
	)
	if err != nil {
		return nil, err
	}

	m.APICalls, err = meter.Int64Counter(
		"futuapi.api.calls",
		metric.WithDescription("API call count by proto and status"),
	)
	if err != nil {
		return nil, err
	}

	m.APILatency, err = meter.Float64Histogram(
		"futuapi.api.latency",
		metric.WithDescription("API call latency in milliseconds"),
		metric.WithUnit("ms"),
	)
	if err != nil {
		return nil, err
	}

	m.PushMessages, err = meter.Int64Counter(
		"futuapi.push.messages",
		metric.WithDescription("Push messages received by type"),
	)
	if err != nil {
		return nil, err
	}

	m.APIErrors, err = meter.Int64Counter(
		"futuapi.api.errors",
		metric.WithDescription("API errors by proto and error code"),
	)
	if err != nil {
		return nil, err
	}

	m.RateLimited, err = meter.Int64Counter(
		"futuapi.rate.limited",
		metric.WithDescription("Rate-limited requests by proto"),
	)
	if err != nil {
		return nil, err
	}

	m.RetryAttempts, err = meter.Int64Counter(
		"futuapi.retry.attempts",
		metric.WithDescription("Retry attempts by proto and attempt number"),
	)
	if err != nil {
		return nil, err
	}

	openDUpGauge, err := meter.Int64ObservableGauge(
		"futuapi.opend_up",
		metric.WithDescription("OpenD connection status (1=up, 0=down)"),
	)
	if err != nil {
		return nil, err
	}

	lastConnectGauge, err := meter.Int64ObservableGauge(
		"futuapi.last.connect.timestamp",
		metric.WithDescription("Timestamp of last successful connection"),
	)
	if err != nil {
		return nil, err
	}

	breakerStateGauge, err := meter.Float64ObservableGauge(
		"futuapi.breaker.state",
		metric.WithDescription("Circuit breaker state (0=closed, 0.5=half-open, 1=open)"),
	)
	if err != nil {
		return nil, err
	}

	m.gaugeReg, err = meter.RegisterCallback(
		func(ctx context.Context, obs metric.Observer) error {
			obs.ObserveInt64(openDUpGauge, m.openDUpValue.Load())
			obs.ObserveInt64(lastConnectGauge, m.lastConnectTimeValue.Load())
			m.breakerStates.Range(func(key, value interface{}) bool {
				obs.ObserveFloat64(breakerStateGauge, value.(float64),
					metric.WithAttributes(
						attribute.String("name", key.(string)),
					),
				)
				return true
			})
			return nil
		},
		openDUpGauge,
		lastConnectGauge,
		breakerStateGauge,
	)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *OTelMeter) Close() error {
	if m.gaugeReg != nil {
		return m.gaugeReg.Unregister()
	}
	return nil
}

func (m *OTelMeter) RecordConnection(connType string) {
	m.Connections.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("type", connType),
	))
	m.lastConnectTimeValue.Store(time.Now().Unix())
}

func (m *OTelMeter) RecordDisconnection(connType string) {
	m.Connections.Add(context.Background(), -1, metric.WithAttributes(
		attribute.String("type", connType),
	))
}

func (m *OTelMeter) RecordReconnect(reason string) {
	m.Reconnects.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("reason", reason),
	))
}

func (m *OTelMeter) RecordAPICall(protoID, status string, latency time.Duration) {
	m.APICalls.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("proto", protoID),
		attribute.String("status", status),
	))
	m.APILatency.Record(context.Background(), float64(latency.Milliseconds()),
		metric.WithAttributes(
			attribute.String("proto", protoID),
		),
	)
}

func (m *OTelMeter) RecordPushMessage(msgType string) {
	m.PushMessages.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("type", msgType),
	))
}

func (m *OTelMeter) RecordOpenDUp(up bool) {
	var val int64 = 0
	if up {
		val = 1
	}
	m.openDUpValue.Store(val)
}

func (m *OTelMeter) RecordAPIError(protoID, errorCode string) {
	m.APIErrors.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("proto", protoID),
		attribute.String("error_code", errorCode),
	))
}

func (m *OTelMeter) RecordRateLimited(protoID string) {
	m.RateLimited.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("proto", protoID),
	))
}

func (m *OTelMeter) RecordRetry(protoID, attempt string) {
	m.RetryAttempts.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("proto", protoID),
		attribute.String("attempt", attempt),
	))
}

func (m *OTelMeter) RecordBreakerState(name string, state float64) {
	m.breakerStates.Store(name, state)
}
