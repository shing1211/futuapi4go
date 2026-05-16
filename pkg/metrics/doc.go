// Package metrics provides client-side performance metrics collection for
// monitoring Futu OpenD connection health, request latency, error rates,
// and push message throughput.
//
// Usage:
//
//	m := client.Metrics()
//	fmt.Printf("requests: %d, errors: %d", m.TotalRequests, m.FailedReqs)
//
// 简体中文:
// metrics 包提供客户端性能指标收集，用于监控富途 OpenD 的连接健康、
// 请求延迟、错误率和推送消息吞吐量。
//
// 繁體中文:
// metrics 包提供客戶端性能指標收集，用於監控富途 OpenD 的連接健康、
// 請求延遲、錯誤率和推送消息吞吐量。
package metrics
