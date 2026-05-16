// Package otel provides an OpenTelemetry-backed Tracer implementation for
// the Futu OpenD SDK tracing framework. Use it to export distributed traces
// to any OTLP-compatible backend (Jaeger, Zipkin, Grafana Tempo, etc.).
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/tracing/otel"
//
//	tracing.SetTracer(otel.NewTracer("my-trading-app"))
//
//	// With custom TracerProvider:
//	tp := otel.GetTracerProvider()
//	tracing.SetTracer(otel.NewTracer("my-app", otel.WithTracerProvider(tp)))
//
// 简体中文:
// otel 包为富途 OpenD SDK 追踪框架提供基于 OpenTelemetry 的 Tracer 实现。
// 可将分布式追踪导出到任何 OTLP 兼容的后端。
//
// 繁體中文:
// otel 包為富途 OpenD SDK 追踪框架提供基於 OpenTelemetry 的 Tracer 實現。
// 可將分佈式追踪導出到任何 OTLP 兼容的後端。
package otel
