package tracing

import (
	"context"
)

type Span interface {
	SetAttribute(key string, value interface{})
	End()
}

type Tracer interface {
	Start(ctx context.Context, name string, attrs ...Attribute) (context.Context, Span)
}

type Attribute struct {
	Key   string
	Value interface{}
}

func StringAttr(k, v string) Attribute  { return Attribute{Key: k, Value: v} }
func IntAttr(k string, v int) Attribute  { return Attribute{Key: k, Value: v} }
func Int64Attr(k string, v int64) Attribute { return Attribute{Key: k, Value: v} }

type noopSpan struct{}

func (noopSpan) SetAttribute(string, interface{}) {}
func (noopSpan) End()                            {}

type NoopTracer struct{}

func (NoopTracer) Start(ctx context.Context, _ string, _ ...Attribute) (context.Context, Span) {
	return ctx, noopSpan{}
}

var defaultTracer Tracer = NoopTracer{}

func SetTracer(t Tracer) {
	if t != nil {
		defaultTracer = t
	}
}

func GetTracer() Tracer {
	return defaultTracer
}

func StartSpan(ctx context.Context, name string, attrs ...Attribute) (context.Context, Span) {
	return defaultTracer.Start(ctx, name, attrs...)
}

func SpanFromContext(ctx context.Context) Span {
	if s, ok := ctx.Value(spanKey).(Span); ok {
		return s
	}
	return noopSpan{}
}

type contextKey struct{}

var spanKey = contextKey{}
