package tracing

import (
	"context"
	"testing"
)

func TestNoopTracer(t *testing.T) {
	tr := NoopTracer{}
	ctx, span := tr.Start(context.Background(), "test")
	if ctx == nil {
		t.Error("expected non-nil context")
	}
	if span == nil {
		t.Error("expected non-nil span")
	}
	span.SetAttribute("key", "value") // should not panic
	span.End()                         // should not panic
}

func TestNoopTracer_StartReturnsNoopSpan(t *testing.T) {
	tr := NoopTracer{}
	_, span := tr.Start(context.Background(), "test")
	if _, ok := span.(noopSpan); !ok {
		t.Error("expected noopSpan from NoopTracer")
	}
}

func TestSetTracer(t *testing.T) {
	custom := &customTracer{}
	SetTracer(custom)
	got := GetTracer()
	if got != custom {
		t.Error("GetTracer should return the custom tracer after SetTracer")
	}
	// Reset to noop
	SetTracer(NoopTracer{})
}

func TestSetTracer_NilNoPanic(t *testing.T) {
	SetTracer(nil) // should not panic
	// Default tracer should remain unchanged
	if GetTracer() == nil {
		t.Error("tracer should not be nil after SetTracer(nil)")
	}
}

func TestStartSpan(t *testing.T) {
	ctx := context.Background()
	ctx2, span := StartSpan(ctx, "test-span")
	if ctx2 == nil {
		t.Error("expected non-nil context")
	}
	if span == nil {
		t.Error("expected non-nil span")
	}
	span.SetAttribute("attr1", "val1")
	span.End()
}

func TestStartSpan_WithAttributes(t *testing.T) {
	ctx := context.Background()
	_, span := StartSpan(ctx, "test",
		StringAttr("string_key", "string_val"),
		IntAttr("int_key", 42),
		Int64Attr("int64_key", 1234567890),
	)
	span.End()
}

func TestSpanFromContext_NoopSpan(t *testing.T) {
	ctx := context.Background()
	span := SpanFromContext(ctx)
	if span == nil {
		t.Error("expected non-nil span from empty context")
	}
	span.SetAttribute("key", "val")
	span.End()
}

func TestSpanFromContext_WithContextKey(t *testing.T) {
	// Verify SpanFromContext retrieves spans stored in context
	ctx := context.WithValue(context.Background(), spanKey, noopSpan{})
	span := SpanFromContext(ctx)
	if span == nil {
		t.Error("expected non-nil span from context with spanKey")
	}
}

func TestAttributeConstructors(t *testing.T) {
	sa := StringAttr("k", "v")
	if sa.Key != "k" || sa.Value != "v" {
		t.Errorf("StringAttr mismatch: %+v", sa)
	}
	ia := IntAttr("k", 42)
	if ia.Key != "k" || ia.Value != 42 {
		t.Errorf("IntAttr mismatch: %+v", ia)
	}
	i64a := Int64Attr("k", 9223372036854775807)
	if i64a.Key != "k" || i64a.Value != int64(9223372036854775807) {
		t.Errorf("Int64Attr mismatch: %+v", i64a)
	}
}

// customTracer implements Tracer for testing SetTracer
type customTracer struct{}

func (c *customTracer) Start(ctx context.Context, name string, attrs ...Attribute) (context.Context, Span) {
	return ctx, noopSpan{}
}