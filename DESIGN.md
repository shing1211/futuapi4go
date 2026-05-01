# futuapi4go Design Document

> **Version:** v0.5.4 | **Last Updated:** 2026-05-02

---

## 1. Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                     Application                          │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   client/Client                           │
│         (High-level wrappers - RECOMMENDED)              │
└─────────────────────────────────────────────────────────────┘
          │                         │                        │
          ▼                         ▼                        ▼
┌─────────────────┐   ┌─────────────────┐   ┌─────────────────┐
│     pkg/qot/     │   │     pkg/trd/    │   │     pkg/sys/    │
│  Market Data    │   │    Trading     │   │    System      │
│   APIs          │   │    APIs        │   │    APIs        │
└─────────────────┘   └─────────────────┘   └─────────────────┘
          │                         │                        │
          └──────────────────────┼────────────────────────┘
                                 ▼
┌─────────────────────────────────────────────────────────────┐
│              internal/client/Client                        │
│     Connection management, reconnection, keep-alive     │
└─────────────────────────────────────────────────────────────┘
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────┐
│               internal/client/Conn                        │
│        TCP I/O, packet framing, buffered I/O            │
└─────────────────────────────────────────────────────────────┘
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────┐
│                    Futu OpenD                             │
│                  (TCP socket)                            │
└─────────────────────────────────────────────────────────────┘
```

**Key Constraint:** All communication is via Protocol Buffers over TCP. No JSON by default.

---

## 2. Key Design Decisions

### 2.1 Protocol Communication

| Decision | Rationale |
|----------|-----------|
| Binary over TCP | Performance, low latency |
| Protobuf serialization | Type safety, schema evolution |
| Custom 46-byte header | Magic "FT" + ProtoID + SerialNo + BodyLen |
| No async/await | Go-native concurrency via goroutines |

**Packet Format:**
```
┌──────────┬─────────┬─────────┬──────────┬─────────────┐
│ Magic(2) │ ProtoID │ SerialNo│ BodyLen │   Body    │
│   "FT"   │ 4 bytes │ 4 bytes │ 4 bytes │  N bytes  │
└──────────┴─────────┴─────────┴──────────┴─────────────┘
```

---

### 2.2 Package Structure

| Package | Responsibility | Public API |
|---------|----------------|------------|
| `client/` | High-level wrappers | Recommended entry point |
| `pkg/qot/` | Market data | `GetBasicQot()`, `GetKL()`, etc. |
| `pkg/trd/` | Trading | `PlaceOrder()`, `GetPositionList()` |
| `pkg/sys/` | System | `GetGlobalState()`, `GetUserInfo()` |
| `pkg/push/` | Push parsers | `ParseUpdateBasicQot()` |
| `pkg/push/chan/` | Channel push | `SubscribeQuote()` |
| `pkg/breaker/` | Circuit breaker | `New()`, `Do()` |
| `pkg/logger/` | Logging | `New()`, `Info()`, etc. |
| `pkg/constant/` | Constants | Typed enums |
| `pkg/util/` | Utilities | `ParseCode()`, `FormatCode()` |
| `internal/client/` | Core TCP | Connection, packet I/O |

---

### 2.3 Error Handling

```go
type FutuError struct {
    Code    ErrorCode
    Message string
    Category ErrorCategory  // API, Connection, Timeout, Trading
    Recovery string        // Suggestion for recovery
}

// Usage
fe, ok := constant.AsFutuError(err)
if ok && fe.Category == constant.ErrorCategoryAPI {
    // handle API error
}
```

**Error Categories:**
| Category | Description | Recovery |
|----------|-------------|----------|
| API | Server returned error | Check RetType, retry |
| Connection | TCP/socket error | Reconnect |
| Timeout | Request timeout | Retry with backoff |
| Trading | Order rejected | Check order params |

---

### 2.4 Context Usage

All APIs accept `context.Context` as FIRST parameter:

```go
// Good
resp, err := qot.GetBasicQot(ctx, cli, securities)

// Bad (won't compile in v0.3.0+)
resp, err := qot.GetBasicQot(cli, securities)
```

**Helpers:**
```go
ctx, cancel := cli.WithTimeout(5 * time.Second)
defer cancel()

ctx, cancel := cli.WithDeadline(time.Now().Add(30 * time.Second))
defer cancel()
```

---

### 2.5 Typed Enums

All constants use typed types (no raw int32 casts):

```go
// Before (v0.2.x)
client.GetQuote(cli, int32(constant.Market_US), "NVDA")

// After (v0.5.x)
client.GetQuote(ctx, cli, constant.Market_US, "NVDA")
```

**Typed Enum Conversion:**
```go
market := constant.Market_HK
protoValue := market.Int32()  // returns int32(1)
```

---

## 3. Connection Lifecycle

```
┌──────────┐     Connect()      ┌──────────┐
│  New()    │ ───────────────▶  │  Created │
└──────────┘                  └──────────┘
                                     │
                                     ▼
                              ┌──────────┐
                              │ Connect  │
                              │   ()     │
                              └──────────┘
                                     │
         ┌───────────��───────────────┼───────────────────────────┐
         │                           │                           │
         ▼                           ▼                           ▼
┌──────────────┐           ┌──────────────┐            ┌──────────────┐
│ InitConnect │           │   Ready for  │            │ Error during │
│  handshake  │           │     API     │            │   connect    │
└──────────────┘           └──────────────┘            └──────────────┘
         │                           │                           │
         │                           │                           ▼
         │                           │                  ┌──────────┐
         │                           │                  │  Error    │
         │                           │                  │ returned  │
         │                           │                  └──────────┘
         ▼                           │
┌──────────────┐                   │
│   Close()     │ ◀──────────────────┘
│  (drain,     │
│   cleanup)   │
└──────────────┘
```

**On Connect, OpenD returns:**
- `connID` — connection ID
- `loginUserID` — Futu/NiuNiu user ID
- `aesKey` — AES encryption key
- `serverVer` — OpenD version
- `keepAliveInterval` — heartbeat interval

---

## 4. API Design Patterns

### 4.1 Market Data APIs

**High-level (recommended):**
```go
quote, err := client.GetQuote(ctx, cli, constant.Market_HK, "00700")
klines, err := client.GetKLines(ctx, cli, constant.Market_HK, "00700", constant.KLType_K_Day, 100)
orderBook, err := client.GetOrderBook(ctx, cli, constant.Market_HK, "00700", 10)
```

**Low-level:**
```go
resp, err := qot.GetBasicQot(ctx, cli, securities)
resp, err := qot.GetKL(ctx, cli, req)
```

---

### 4.2 Trading APIs

```go
// Place order with typed constants
result, err := client.PlaceOrder(ctx, cli,
    accID,
    constant.TrdMarket_HK,      // trading market
    "00700",                    // code
    constant.TrdSide_Buy,        // side
    constant.OrderType_Normal,   // order type
    350.0,                     // price
    100,                       // quantity
)
```

**OrderBuilder pattern:**
```go
order := trd.NewOrder(accID, constant.TrdMarket_HK, constant.TrdEnv_Simulate).
    Buy("00700", 100).
    At(350.0).
    Build()
```

---

### 4.3 Push APIs

**Channel-based (recommended):**
```go
ch := make(chan *push.UpdateBasicQot, 100)
stop := chanpkg.SubscribeQuote(cli, constant.Market_HK, "00700", ch)
defer stop()

for q := range ch {
    fmt.Printf("Price: %.2f\n", q.CurPrice)
}
```

**Handler-based:**
```go
cli.RegisterHandler(constant.ProtoID_Qot_UpdateBasicQot, func(pid uint32, body []byte) {
    q, _ := push.ParseUpdateBasicQot(body)
    fmt.Printf("Price: %.2f\n", q.CurPrice)
})
```

---

## 5. Performance Optimizations

| Feature | Implementation | Impact |
|---------|---------------|--------|
| Buffered I/O | 64KB bufio.Reader in conn.go | Reduced syscalls |
| Zero-allocation | sync.Pool in alloc.go | GC reduction |
| Pool O(1) lookup | clientIndex map in pool.go | < 1μs lookup |
| Rate limiting | Token bucket in pkg/ratelimit/ | API protection |
| Batch subscribe | SubscribeSymbols() | Single round-trip |

---

## 6. Security Model

| Layer | Responsibility | Implementation |
|-------|---------------|----------------|
| **OpenD** | Authentication | Password, 2FA |
| **futuapi4go** | Safe protobuf | No deserialization vulnerabilities |
| **futuapi4go** | TCP management | Keep-alive, reconnection |
| **futuapi4go** | Credential handling | SensitiveString type |
| **User app** | Credential security | Environment variables |
| **User app** | Trading safeguards | Validate before trade |

**Sensitive Data:**
```go
type SensitiveString string

func (s SensitiveString) String() string {
    return "***"
}

// Usage: Password redacted in all fmt output
req := &trd.UnlockTradeRequest{
    PwdMD5: constant.SensitiveString("actual_password"),
}
fmt.Printf("%v", req) // Prints: {PwdMD5: ***}
```

---

## 7. Thread Safety

| Component | Protection |
|-----------|------------|
| Connection pool | `sync.RWMutex` |
| Client state | `sync.Mutex` |
| Push handlers | `sync.RWMutex` |
| Subscription map | `sync.RWMutex` |

**Pattern:**
```go
func (c *Client) EnsureConnected() error {
    c.mu.Lock()
    defer c.mu.Unlock()
    if c.conn == nil {
        return ErrNotConnected
    }
    return nil
}
```

---

## 8. Extensibility Points

### 8.1 Custom Push Handlers
```go
cli.RegisterHandler(constant.ProtoID_Qot_UpdateBasicQot, 
    func(protoID uint32, body []byte) {
        // Custom handling
    })
```

### 8.2 Connection Pool
```go
pool := futuapi.NewClientPool(config)
cli, _ := pool.Get(ctx, PoolTypeMarketData)
defer pool.Put(cli)
```

### 8.3 Circuit Breaker
```go
cb := breaker.New(
    breaker.WithThreshold(5),
    breaker.WithCooldown(30*time.Second),
)

result, err := cb.Do(func() (interface{}, error) {
    return client.PlaceOrder(...)
})
```

### 8.4 Rate Limiter
```go
limiter := rate.NewLimiter(rate.WithLimit(100)) // 100 req/sec
cli.SetRateLimiter(limiter)
```

### 8.5 Custom Logger
```go
cli.SetLogger(log.New(os.Stderr, "", 0))
// or use pkg/logger for structured logging
```

---

## 9. Dependencies

### Direct (go.mod)
```
google.golang.org/protobuf v1.x.x
```

### Generated (pkg/pb/)
- 78 protobuf message types (v10.4.6408)
- All under `github.com/shing1211/futuapi4go/pkg/pb/`

---

## 10. Version Compatibility

| Component | Version |
|------------|---------|
| Go | 1.26+ |
| OpenD | v10.4.6408 (recommended) |
| Protobuf | proto3 |

---

## 11. Quick Reference

| Operation | Code |
|------------|------|
| Connect | `cli.Connect("127.0.0.1:11111")` |
| Get quote | `client.GetQuote(ctx, cli, Market_HK, "00700")` |
| Place order | `client.PlaceOrder(ctx, cli, accID, TrdMarket_HK, code, TrdSide_Buy, OrderType_Normal, price, qty)` |
| Subscribe | `client.Subscribe(ctx, cli, Market_HK, "00700", []SubType{SubType_Quote})` |
| Close | `cli.Close()` |

See [README.md](README.md) for complete API reference.