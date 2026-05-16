// Package tracing provides the core interfaces for distributed tracing
// within the Futu OpenD SDK. It defines Tracer and Span abstractions with
// a no-op default implementation, allowing users to plug in any backend
// (e.g., OpenTelemetry via pkg/tracing/otel).
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/tracing"
//
//	// Install OpenTelemetry backend
//	import "github.com/shing1211/futuapi4go/pkg/tracing/otel"
//	tracing.SetTracer(otel.NewTracer("my-service"))
//
//	// Start spans anywhere in your code
//	ctx, span := tracing.StartSpan(ctx, "operation_name")
//	defer span.End()
//
// 简体中文:
// tracing 包提供富途 OpenD SDK 的分布式追踪核心接口。
// 定义了 Tracer 和 Span 抽象，内置空实现（no-op），
// 用户可以接入任意后端（例如通过 pkg/tracing/otel 使用 OpenTelemetry）。
//
// 繁體中文:
// tracing 包提供富途 OpenD SDK 的分佈式追踪核心接口。
// 定義了 Tracer 和 Span 抽象，內置空實現（no-op），
// 用戶可以接入任意後端（例如通過 pkg/tracing/otel 使用 OpenTelemetry）。
package tracing
