# 🚀 futuapi4go: The Resilient Open-Source Go SDK for Futu OpenAPI

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" alt="Language">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=for-the-badge" alt="License">
  <img src="https://img.shields.io/badge/Version-0.8.0-blue?style=for-the-badge" alt="Version">
  <img src="https://img.shields.io/badge/Status-Stable-brightgreen?style=for-the-badge" alt="Status">
</p>

---

## ✨ Introduction: Your Gateway to Futu Trading Data

`futuapi4go` is a powerful, high-performance, and resilient Go SDK engineered to provide developers with seamless access to the comprehensive suite of trading and market data provided by the Futu OpenAPI. We've built this library from the ground up using modern Go concurrency patterns (goroutines, channels) and robust error handling to ensure it performs reliably even in challenging real-world network environments.

**We are building this project with an Open Source mindset.** Our commitment is to create a developer experience that is intuitive, reliable, and easy to contribute to.

## 💡 Key Features
*   **Unified API:** Single interface for Market Data (quotes, KLines), Trading Operations (`PlaceOrder`, `GetPositionList`), and System Status checks.
*   **Robust Connection Handling:** Features a sophisticated connection pool with automatic reconnection logic and configurable health checks.
*   **Asynchronous Push Notifications:** Easily subscribe to real-time market data and order updates using structured push handlers.
*   **Advanced Market Tools:** Access specialized functions like `RequestHistoryKL` (with auto-pagination), `GetOptionChain`, and `GetSecuritySnapshot`.

## ⚡️ Getting Started: A Quick Dive

### Prerequisites
Before you begin, ensure you have a basic understanding of the Futu OpenAPI requirements, including RSA public key generation for secure connection.

1.  **Installation:**
    ```bash
    go get github.com/shing1211/futuapi4go
    ```
2.  **Connection Example (Simplified):**
    *(Note: Full implementation requires robust error checking.)*
    ```go
    package main

    import (
        "fmt"
        "github.com/shing1211/futuapi4go/client"
        "github.com/shing1211/futuapi4go/pkg/qot"
    )

    func main() {
        // Replace with your actual RSA public key PEM string
        publicKeyPEM := "YOUR_RSA_PUBLIC_KEY_HERE" 
        
        cli := client.New() // Use New() for default options
        defer cli.Close()

        fmt.Println("Attempting connection...")
        if err := cli.ConnectWithRSA("127.0.0.1:11111", publicKeyPEM); err != nil {
            panic(fmt.Sprintf("FATAL CONNECTION ERROR: %v", err))
        }

        // Example Market Data Request (HSI)
        marketCode := 2 // HK Market Code
        symbolCode := "HSImain" 
        
        quote, err := qot.GetBasicQot(cli, marketCode, symbolCode)
        if err != nil {
            fmt.Printf("Failed to get quote: %v\n", err)
        } else {
            fmt.Printf("📈 Quote for %s: Price=%.2f | Volume=%d\n", quote.Symbol, quote.Price, quote.Volume)
        }
    }
    ```

## 🏗️ Architectural Overview (For Developers)

The SDK follows a layered architecture designed for modularity and testability:

`Application Logic` $\rightarrow$ `client/Client` (Public Interface) $\rightarrow$ `pkg/*` (Business Logic Wrappers, e.g., `qot`, `trd`) $\rightarrow$ `internal/client/Conn` (Low-Level Network Handling) $\rightarrow$ `Futu OpenD Gateway`.

**Key Components:**
*   **`ClientPool` (`internal/client/pool.go`):** Manages and reuses underlying connections for efficiency, minimizing connection overhead per request.
*   **`Connection` (`internal/client/conn.go`):** Handles the raw TCP socket communication, packet serialization/deserialization, and the primary read loop.
*   **Push Handlers (`pkg/push/*.go`):** Dedicated parsers to translate complex binary push notifications into clean Go structs for immediate use by the application layer.

## 🛠️ Contributing & Development Standards (Open Source Focus)

We welcome contributions! Whether it's a bug fix, a feature addition, or documentation improvement, your help is invaluable.

1.  **Code Style:** Adhere to standard Golang conventions and maintain consistency with existing patterns in `futuapi4go`.
2.  **Testing:** All new features must be accompanied by unit tests (`*_test.go`). Utilize the provided mock server utilities for integration testing where possible.
3.  **Process:** Follow the Git workflow (branching, committing) and submit a Pull Request with a clear description of the changes.

For detailed contribution guidelines, please refer to [`CONTRIBUTING.md`](./CONTRIBUTING.md). For questions, feel free to open an issue!

---
### 🛡️ Legal & Disclaimer
**WARNING: TRADING FINANCIAL INSTRUMENTS CARRIES SIGNIFICANT RISK.** The `futuapi4go` library is a software utility only and does not provide financial advice. By using this SDK, you acknowledge that all trading decisions are made at your own risk. We provide this software "as is."

*Trademark Notice: Futu Holdings Limited trademarks are used descriptively to indicate the protocol implementation. This project is independent of Futu.*
