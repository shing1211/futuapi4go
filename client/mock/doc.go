// Package mock provides a simulated Futu OpenD mock server for integration testing.
// Use it to test client code without connecting to a real OpenD instance.
//
// Usage:
//
//	srv := mock.NewServer()
//	go srv.Start(":0")
//	defer srv.Stop()
//	addr := srv.Addr()
//	cli.Connect(addr)
//
// 简体中文:
// mock 包提供了模拟的富途 OpenD 服务器，用于集成测试。
// 无需连接真实的 OpenD 即可测试客户端代码。
//
// 繁體中文:
// mock 包提供了模擬的富途 OpenD 伺服器，用於集成測試。
// 無需連接真實的 OpenD 即可測試客戶端代碼。
package mock
