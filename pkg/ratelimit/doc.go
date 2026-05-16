// Package ratelimit provides API rate limiting utilities to prevent exceeding
// Futu OpenD's request rate limits. It uses a token bucket algorithm with
// per-protocol-ID rate tracking.
//
// Usage:
//
//	limiter := ratelimit.New(ratelimit.WithRate(10, time.Second))
//	if limiter.Allow(protoID) {
//	    // send request
//	}
//
// 简体中文:
// ratelimit 包提供 API 频率限制工具，防止超出富途 OpenD 的请求速率限制。
// 使用令牌桶算法，支持按协议 ID 分别限速。
//
// 繁體中文:
// ratelimit 包提供 API 頻率限制工具，防止超出富途 OpenD 的請求速率限制。
// 使用令牌桶算法，支持按協議 ID 分別限速。
package ratelimit
