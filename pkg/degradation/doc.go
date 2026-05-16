// Package degradation provides graceful degradation utilities for the Futu OpenD SDK.
// When the connection to OpenD is interrupted, degradation strategies allow the
// application to continue serving with cached or fallback data.
//
// Usage:
//
//	degrader := degradation.New(degradation.WithFallback(data))
//	result, err := degrader.Execute(ctx, primaryOp)
//
// 简体中文:
// degradation 包提供富途 OpenD SDK 的优雅降级工具。
// 当与 OpenD 的连接中断时，降级策略允许应用使用缓存或备用数据继续服务。
//
// 繁體中文:
// degradation 包提供富途 OpenD SDK 的優雅降級工具。
// 當與 OpenD 的連接中斷時，降級策略允許應用使用緩存或備用數據繼續服務。
package degradation
