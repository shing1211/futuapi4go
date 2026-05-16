package otel

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel/trace"

	"github.com/shing1211/futuapi4go/pkg/tracing"
)

func TestNewTracer(t *testing.T) {
	tr := NewTracer("test-service")
	if tr == nil {
		t.Fatal("expected non-nil tracer")
	}
	ctx, span := tr.Start(context.Background(), "test-span")
	if ctx == nil {
		t.Error("expected non-nil context")
	}
	if span == nil {
		t.Error("expected non-nil span")
	}
	span.SetAttribute("key", "value")
	span.End()
}

func TestNewTracer_WithAttributes(t *testing.T) {
	tr := NewTracer("test-service")
	_, span := tr.Start(context.Background(), "test",
		tracing.StringAttr("str", "hello"),
		tracing.IntAttr("int", 42),
		tracing.Int64Attr("int64", 9223372036854775807),
		tracing.BoolAttr("bool", true),
	)
	span.End()
}

func TestNewTracer_AllValueTypes(t *testing.T) {
	tr := NewTracer("test-service")
	_, span := tr.Start(context.Background(), "test-values",
		tracing.StringAttr("str", "hello"),
		tracing.IntAttr("int", 42),
		tracing.Int64Attr("int64", 9223372036854775807),
		tracing.BoolAttr("bool", true),
	)
	span.SetAttribute("late_str", "world")
	span.SetAttribute("late_int", 100)
	span.SetAttribute("late_float", 3.14)
	span.End()
}

func TestWithTracerProvider(t *testing.T) {
	tp := trace.NewNoopTracerProvider()
	tr := NewTracer("test", WithTracerProvider(tp))
	if tr == nil {
		t.Fatal("expected non-nil tracer with custom provider")
	}
	_, span := tr.Start(context.Background(), "test")
	span.End()
}

func TestNewTracer_NilProviderDefaultsToGlobal(t *testing.T) {
	tr := NewTracer("test")
	_, span := tr.Start(context.Background(), "test")
	span.End()
}

func TestSpan_ImplementsInterface(t *testing.T) {
	tr := NewTracer("test")
	_, span := tr.Start(context.Background(), "test")
	otelspan := span.(*otelspan)
	if otelspan.span == nil {
		t.Error("expected non-nil underlying OTel span")
	}
	span.End()
}

func TestMultipleSpans(t *testing.T) {
	tr := NewTracer("test-service")
	ctx1, span1 := tr.Start(context.Background(), "parent",
		tracing.StringAttr("level", "parent"),
	)
	_, span2 := tr.Start(ctx1, "child",
		tracing.StringAttr("level", "child"),
	)
	span2.End()
	span1.End()
}

func TestToValue_DefaultCase(t *testing.T) {
	v := toValue(struct{}{})
	_ = v
}

func TestNewTracer_WithNoAttrs(t *testing.T) {
	tr := NewTracer("no-attrs")
	_, span := tr.Start(context.Background(), "bare")
	span.End()
}
