package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/shing1211/futuapi4go/pkg/tracing"
)

func NewTracer(serviceName string, opts ...Option) tracing.Tracer {
	cfg := &config{
		serviceName: serviceName,
	}
	for _, o := range opts {
		o(cfg)
	}

	tp := cfg.tracerProvider
	if tp == nil {
		tp = otel.GetTracerProvider()
	}

	tracer := tp.Tracer(
		"github.com/shing1211/futuapi4go",
		trace.WithInstrumentationVersion("0.6.2"),
	)

	return &oteltracer{
		tracer: tracer,
		cfg:    cfg,
	}
}

type oteltracer struct {
	tracer trace.Tracer
	cfg    *config
}

func (t *oteltracer) Start(ctx context.Context, name string, attrs ...tracing.Attribute) (context.Context, tracing.Span) {
	otelAttrs := make([]attribute.KeyValue, len(attrs))
	for i, a := range attrs {
		otelAttrs[i] = toKeyValue(a)
	}

	ctx, span := t.tracer.Start(ctx, name,
		trace.WithAttributes(otelAttrs...),
	)

	return ctx, &otelspan{span: span}
}

type otelspan struct {
	span trace.Span
}

func (s *otelspan) SetAttribute(key string, value interface{}) {
	s.span.SetAttributes(attribute.KeyValue{
		Key:   attribute.Key(key),
		Value: toValue(value),
	})
}

func (s *otelspan) End() {
	s.span.End()
}

func toKeyValue(a tracing.Attribute) attribute.KeyValue {
	return attribute.KeyValue{
		Key:   attribute.Key(a.Key),
		Value: toValue(a.Value),
	}
}

func toValue(v interface{}) attribute.Value {
	switch val := v.(type) {
	case string:
		return attribute.StringValue(val)
	case int:
		return attribute.IntValue(val)
	case int64:
		return attribute.Int64Value(val)
	case bool:
		return attribute.BoolValue(val)
	case float64:
		return attribute.Float64Value(val)
	default:
		return attribute.StringValue("")
	}
}

type Option func(*config)

type config struct {
	serviceName    string
	tracerProvider trace.TracerProvider
}

func WithTracerProvider(tp trace.TracerProvider) Option {
	return func(c *config) {
		c.tracerProvider = tp
	}
}
