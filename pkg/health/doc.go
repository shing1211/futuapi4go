// Package health provides health check utilities for monitoring the
// connection status and responsiveness of Futu OpenD. Use it to implement
// liveness and readiness probes in production deployments.
//
// Usage:
//
//	checker := health.New(client)
//	status := checker.Check(ctx)
//	// status.Connected, status.Latency, status.LastError
//
// 简体中文:
// health 包提供健康检查工具，用于监控富途 OpenD 的连接状态和响应能力。
// 可在生产部署中用于实现存活性（liveness）和就绪性（readiness）探针。
//
// 繁體中文:
// health 包提供健康檢查工具，用於監控富途 OpenD 的連接狀態和響應能力。
// 可在生產部署中用於實現存活性（liveness）和就緒性（readiness）探針。
package health
