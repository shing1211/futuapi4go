// Package retry provides configurable retry logic with exponential backoff
// for transient failures in Futu OpenD operations. It supports jitter,
// max retries, and context-based cancellation.
//
// Usage:
//
//	cfg := retry.NewConfig()
//	err := retry.Do(ctx, cfg, func() error {
//	    return cli.GetSecuritySnapshot(ctx, "HK.00700", &snap)
//	})
//
// 简体中文:
// retry 包提供可配置的重试逻辑，带指数退避策略，用于富途 OpenD 操作的瞬时故障恢复。
// 支持抖动（jitter）、最大重试次数和基于 context 的取消。
//
// 繁體中文:
// retry 包提供可配置的重試邏輯，帶指數退避策略，用於富途 OpenD 操作的瞬時故障恢復。
// 支持抖動（jitter）、最大重試次數和基於 context 的取消。
package retry
