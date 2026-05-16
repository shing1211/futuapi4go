// Package breaker implements a circuit breaker pattern for preventing
// cascading failures when OpenD connections are unreliable. It tracks
// consecutive failures and opens the circuit when a threshold is exceeded.
//
// Usage:
//
//	b := breaker.New(breaker.WithThreshold(5))
//	if b.Allow() {
//	    err := riskyOperation()
//	    b.Report(err == nil)
//	}
//
// 简体中文:
// breaker 包实现了断路器模式，在 OpenD 连接不可靠时防止级联故障。
// 跟踪连续失败次数，超过阈值时断开电路。
//
// 繁體中文:
// breaker 包實現了斷路器模式，在 OpenD 連接不可靠時防止級聯故障。
// 跟踪連續失敗次數，超過阈值时斷開電路。
package breaker
