# Futu API for Golang (Updated)

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=for-the-badge" alt="License">
  <img src="https://img.shields.io/badge/Version-0.6.1-blue?style=for-the-badge" alt="Version">
  <img src="https://img.shields.io/badge/Status-Beta-orange?style=for-the-badge" alt="Status">
</p>

<p align="center">
  <strong>Professional-Grade Go SDK for Futu OpenAPI</strong><br>
  Robust and resilient trading and market data interface, enhanced with modern Go best practices.
 </p>

---

## ⭐ What's New in v0.6.1

*   **Bug Fixes:** Push notification parse functions now correctly unmarshal `S2C` directly (matching OpenD push body format). Fixed nil logger panic and connection state race condition.
*   **Stability:** Core unit tests all pass. See [docs/CHANGELOG.md](docs/CHANGELOG.md) for full release notes.

---

## 🚀 Quick Start

### 1. Connect and Get Quote

```go
import (
    "github.com/shing1211/futuapi4go/client"
    "github.com/shing1211/futuapi4go/pkg/qot"
    "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
)

cli := client.New()
defer cli.Close()

if err := cli.ConnectWithRSA("127.0.0.1:11111", publicKeyPEM); err != nil {
    panic(err)
}
defer cli.Close()

quote, err := qot.GetBasicQot(cli, &qotcommon.Security{Market: 2, Code: &"HSImain"})
if err != nil {
    panic(err)
}
fmt.Printf("HSI: %.2f\n", quote.CurPrice)
```

---

## 🛠️ Architecture Deep Dive (For Developers)
The library's architecture has been fundamentally redesigned to adhere to modern Go patterns:

**High-Level Flow:** `Your Application` $\rightarrow$ `client/Client` $\rightarrow$ `internal/client/Conn` $\rightarrow$ `Futu OpenD Gateway`.

## 🔗 API Reference & Documentation
For detailed function signatures and type definitions, please consult:
*   [docs/API_REFERENCE.md](docs/API_REFERENCE.md) — Full function signatures
*   [docs/DEVELOPER.md](docs/DEVELOPER.md) — Architecture and internal structure
*   [docs/TESTING.md](docs/TESTING.md) — Testing guide
*   [ROADMAP.md](ROADMAP.md) — Project roadmap

---

## ⚠️ Disclaimer

**futuapi4go** is a software library only. Trading financial instruments carries significant risk
of financial loss. Past performance does not guarantee future results. This software is provided
"as is" without warranty. Users are solely responsible for any trading decisions and their
outcomes. This library does not constitute financial advice.

## ™ Trademark Notice

**"Futu"**, **"moomoo"**, **"牛牛"**, and **"富途"** are trademarks of **Futu Holdings Limited**.
This project is an independent open-source project and is not affiliated with, endorsed by,
or connected to Futu Holdings Limited. Use of these trademarks in this project is purely
descriptive (the project implements a client for the Futu OpenD protocol) and does not imply
any official relationship or endorsement.

***