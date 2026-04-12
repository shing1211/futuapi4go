# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=for-the-badge" alt="License">
  <img src="https://img.shields.io/badge/Version-0.4.1-blue?style=for-the-badge" alt="Version">
  <img src="https://img.shields.io/badge/Status-Production%20Ready-brightgreen?style=for-the-badge" alt="Status">
</p>

<p align="center">
  <strong>Production-Ready Go SDK for Futu OpenAPI</strong><br>
  Professional-grade trading and market data interface for quantitative developers
</p>

---

<p align="center">
  <a href="#features">Features</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage Examples</a> •
  <a href="#api-reference">API Reference</a> •
  <a href="#testing">Testing</a> •
  <a href="#documentation">Documentation</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

---

## ⭐ Why futuapi4go?

| Feature | Description |
|---------|-------------|
| **🚀 Production Ready** | 59 high-level wrapper functions + 74 low-level APIs, battle-tested in production |
| **📊 Real-Time Data** | Live quotes, K-lines (1min to month), order book, tick-by-tick, broker queue |
| **💼 Full Trading** | Place orders, modify/cancel, position management, account funds |
| **🔔 Push Notifications** | Real-time market data and order updates via WebSocket-like callbacks |
| **🎯 100% Go** | Pure Go implementation, no CGO dependencies, cross-platform compatible |
| **📦 Apache 2.0 Licensed** | Open source, free for commercial use, with patent protection |

---

## 🚀 Quick Start

### 1. Connect and Get Quote

```go
package main

import (
    "fmt"
    "log"

    "github.com/shing1211/futuapi4go/client"
)

func main() {
    // Create client
    cli := client.New()
    defer cli.Close()

    // Connect to OpenD
    if err := cli.Connect("127.0.0.1:11111"); err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }

    fmt.Printf("Connected! ConnID: %d\n", cli.GetConnID())

    // Subscribe to real-time data
    if err := client.Subscribe(cli, client.Market_HK_Future, "HSImain", 
        []int32{client.SubType_Basic, client.SubType_KL_1Min}); err != nil {
        log.Printf("Subscribe warning: %v", err)
    }

    // Get quote
    quote, err := client.GetQuote(cli, client.Market_HK_Future, "HSImain")
    if err != nil {
        log.Fatalf("GetQuote failed: %v", err)
    }

    fmt.Printf("HSImain Price: %.2f\n", quote.Price)
    fmt.Printf("Open: %.2f, High: %.2f, Low: %.2f\n", quote.Open, quote.High, quote.Low)
}
```

### 2. Get K-Lines (Historical Data)

```go
// Get 100 days of daily K-lines for a stock
klines, err := client.GetKLines(cli, client.Market_HK_Security, "00700", client.KLType_Day, 100)
if err != nil {
    log.Fatal(err)
}

for _, kl := range klines {
    fmt.Printf("%s: O=%.2f H=%.2f L=%.2f C=%.2f V=%d\n",
        kl.Time, kl.Open, kl.High, kl.Low, kl.Close, kl.Volume)
}
```

### 3. Place Order (Requires Trading Account)

```go
// Unlock trading first (MD5 hash of password)
if err := client.UnlockTrading(cli, "your_password_md5"); err != nil {
    log.Fatal(err)
}

// Place an order
result, err := client.PlaceOrder(cli,
    accID,                    // Account ID
    client.Market_HK_Security, // Market
    "00700",                   // Stock code
    client.Side_Buy,          // Buy side
    client.OrderType_Normal,  // Normal order
    350.00,                   // Price
    100)                      // Quantity
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Order placed! OrderID: %d\n", result.OrderID)
```

---

## 📦 Installation

```bash
go get github.com/shing1211/futuapi4go@latest
```

Or use the latest development version:

```bash
go get github.com/shing1211/futuapi4go@v0.4.1
```

### Requirements

- **Go**: 1.21 or higher
- **Futu OpenD**: Running on your machine (download from [moomoo](https://www.moomoo.com))

---

## 🏗️ Project Structure

```
futuapi4go/
├── client/                 # High-level wrapper APIs (59 functions)
│   └── client.go          # Main client with all wrappers
├── pkg/
│   ├── qot/              # Market data APIs (quote + market)
│   │   ├── quote.go      # Real-time quote functions
│   │   └── market.go     # Market info functions
│   ├── trd/              # Trading APIs
│   │   └── trade.go      # Order, position, account functions
│   ├── sys/              # System APIs
│   │   └── system.go     # Global state, user info
│   └── pb/               # Protobuf definitions (74 APIs)
│       ├── qotcommon/   # Common QOT types
│       └── trdcommon/   # Common trading types
├── internal/
│   └── client/           # Low-level TCP connection
│       ├── conn.go      # Connection handling
│       ├── client.go    # Main client
│       └── pool.go      # Connection pool
├── push/                 # Push notification handlers
├── test/                 # Integration tests
│   ├── integration/      # Integration tests
│   ├── qot_api/         # Market data API tests
│   └── trd_api/         # Trading API tests
├── cmd/examples/         # Example programs (29 examples)
│   ├── 01_market_data_basic/
│   ├── 02_market_data_advanced/
│   ├── 03_trading_operations/
│   ├── 04_push_subscriptions/
│   ├── 05_comprehensive_demo/
│   └── algo_*/          # Algorithm examples
└── docs/                 # Documentation
```

---

## 📊 Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Your Application                         │
└──────────────────────┬──────────────────────────────────────┘
                       │ calls high-level wrappers
        ┌──────────────▼──────────────┐
        │     client/ (59 APIs)       │  ← Start here!
        │  - GetQuote                 │
        │  - GetKLines                │
        │  - Subscribe                │
        │  - PlaceOrder               │
        │  - ...                      │
        └──────────────┬──────────────┘
                       │ uses low-level functions
        ┌──────────────▼──────────────┐
        │   pkg/qot, pkg/trd, pkg/sys │  ← Advanced users
        │   (74 protobuf APIs)        │
        └──────────────┬──────────────┘
                       │ wraps protobuf
        ┌──────────────▼──────────────┐
        │   pkg/pb/ (74 protobufs)    │  ← Protocol definitions
        └──────────────┬──────────────┘
                       │ serializes
        ┌──────────────▼──────────────┐
        │   internal/client/          │  ← TCP connection
        │   - Conn (TCP)              │
        │   - Packet I/O            │
        └──────────────┬──────────────┘
                       │ network
        ┌──────────────▼──────────────┐
        │      Futu OpenD             │  ← Local Gateway
        │   (127.0.0.1:11111)        │
        └─────────────────────────────┘
```

---

## 🔧 Configuration

### OpenD Connection

| Parameter | Default | Description |
|-----------|---------|-------------|
| Address | `127.0.0.1:11111` | OpenD TCP address |
| Timeout | `30s` | API call timeout |

```go
cli := client.New()

// Default connection (127.0.0.1:11111)
cli.Connect("")

// Custom address
cli.Connect("192.168.1.100:11111")
```

### Market Constants

```go
// Markets
client.Market_HK_Security   // Hong Kong Stocks
client.Market_HK_Future     // Hong Kong Futures
client.Market_US_Security   // US Stocks
client.Market_CNSH_Security // Shanghai Stocks
client.Market_CNSZ_Security // Shenzhen Stocks

// K-Line Types
client.KLType_1Min, client.KLType_5Min, client.KLType_15Min
client.KLType_30Min, client.KLType_60Min, client.KLType_Day
client.KLType_Week, client.KLType_Month

// Subscription Types
client.SubType_Basic      // Basic quote
client.SubType_KL_1Min    // 1-minute K-line
client.SubType_OrderBook  // Order book
client.SubType_Ticker     // Tick-by-tick
client.SubType_RT         // Real-time minute data
client.SubType_Broker     // Broker queue
```

---

## 📖 Usage Examples

### Example 1: Market Data Pipeline

See [`cmd/examples/01_market_data_basic/main.go`](cmd/examples/01_market_data_basic/main.go)

### Example 2: Real-Time Subscription

See [`cmd/examples/04_push_subscriptions/main.go`](cmd/examples/04_push_subscriptions/main.go)

### Example 3: Trading Operations

See [`cmd/examples/03_trading_operations/main.go`](cmd/examples/03_trading_operations/main.go)

### Example 4: Algorithm - SMA Crossover

See [`cmd/examples/algo_sma_crossover/main.go`](cmd/examples/algo_sma_crossover/main.go)

---

## 🧪 Testing

### Run All Tests

```bash
# From project root
go test ./...

# With verbose output
go test -v ./...

# Run specific test
go test -v ./client/... -run TestAllWrapperFunctionsExist
```

### Test Categories

| Category | Location | Command |
|----------|----------|---------|
| Unit Tests | `pkg/**/*_test.go` | `go test ./pkg/...` |
| Integration | `test/integration/` | `go test ./test/integration/...` |
| Examples | `cmd/examples/` | `go test ./test/examples/...` |
| Benchmarks | `test/benchmark/` | `go test -bench=. ./test/benchmark/...` |

---

## 📚 Documentation

| Document | Description |
|----------|-------------|
| [README.md](README.md) | This file |
| [docs/API_REFERENCE.md](docs/API_REFERENCE.md) | Complete API reference |
| [docs/USER_GUIDE.md](docs/USER_GUIDE.md) | User guide with tutorials |
| [docs/DEVELOPER.md](docs/DEVELOPER.md) | Developer guide |
| [docs/TESTING.md](docs/TESTING.md) | Testing guide |

---

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guide](docs/CONTRIBUTING.md).

### Development Setup

```bash
# Clone the repository
git clone https://github.com/shing1211/futuapi4go.git
cd futuapi4go

# Run tests
go test ./...

# Build examples
go build ./cmd/examples/...

# Run a specific example
go run cmd/examples/01_market_data_basic/main.go
```

---

## ⚠️ Disclaimer

**futuapi4go** is an open source project and is **not affiliated with**, **sponsored by**, or **endorsed by** Futu Holdings Limited (Futu), moomoo, or any of their subsidiaries.

This SDK implements a client for the Futu OpenD protocol. It is a clean-room implementation that does not contain any proprietary code from Futu. Users must comply with Futu's terms of service when using the OpenD API.

Use at your own risk. Trading involves financial risk. The authors assume no liability for any losses incurred while using this software.

---

## 📄 License

Apache License 2.0 - see [LICENSE](LICENSE) for details.

---

## 🙏 Acknowledgments

- [Futu OpenAPI](https://openapi.futunn.com/) - Official API documentation
- [py-futu-api](https://github.com/FutunnOpen/py-futu-api) - Python SDK reference
- All contributors and testers

---

<p align="center">
  <strong>⭐ Star us on GitHub if you find this useful!</strong>
</p>