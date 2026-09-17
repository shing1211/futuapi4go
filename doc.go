// Package futuapi4go is a typed Go SDK for the Futu OpenD / OpenAPI protocol.
//
// It wraps the protobuf-over-TCP Qot (market data), Trd (trading), and Sys
// services in idiomatic Go, and adds a connection state machine with
// auto-reconnect, real-time push via channels or typed callbacks, rate limiting,
// circuit breaking and retry, optional OpenTelemetry tracing and metrics, a
// K-line LRU cache, and order pre-flight validation.
//
// The high-level entry point is the client package:
//
//	import "github.com/shing1211/futuapi4go/client"
//
//	cli := client.New(client.WithEnvConfig())
//	if err := cli.Connect("127.0.0.1:11111"); err != nil {
//		log.Fatal(err)
//	}
//	defer cli.Close()
//
// See the client package and the repository README for details.
//
// 简体中文:
// futuapi4go 是富途 OpenD / OpenAPI 协议的 Go 语言 SDK，封装了行情（Qot）、
// 交易（Trd）和系统（Sys）服务，并提供断线重连、实时推送、限流、熔断重试、
// OpenTelemetry 可观测性、K 线缓存和下单前校验等能力。
//
// 繁體中文:
// futuapi4go 是富途 OpenD / OpenAPI 協定的 Go 語言 SDK，封裝了行情（Qot）、
// 交易（Trd）與系統（Sys）服務，並提供斷線重連、即時推送、限流、熔斷重試、
// OpenTelemetry 可觀測性、K 線快取與下單前校驗等能力。
package futuapi4go
