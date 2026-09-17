# futuapi4go Design Document

> **Version:** v0.19.3 | **Last Updated:** 2026-09-17 | **Futu Protocol:** v10.10.7008 (184 protos)

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

**Key Constraint:** All communication is via Protocol Buffers over TCP (or the optional WebSocket transport). There is no JSON/HTTP API.

---

## 2. Key Design Decisions

### 2.1 Protocol Communication

| Decision | Rationale |
|----------|-----------|
| Binary over TCP | Performance, low latency |
| Protobuf serialization | Type safety, schema evolution |
| Custom 44-byte header (see `internal/client/conn.go`) | Magic "FT" + ProtoID + ProtoFmt/ProtoVer + SerialNo + BodyLen + SHA-1 + Reserved |
| No async/await | Go-native concurrency via goroutines |

**Packet Format** (44-byte header, little-endian, then the protobuf body):

```
┌────────┬────────┬────────┬────────┬─────────┬─────────┬───────────┬──────────┐
│ Magic  │ ProtoID│ProtoFmt│ProtoVer│ SerialNo│ BodyLen │ BodySHA1  │ Reserved │
│ 2 "FT" │ 4      │ 1      │ 1      │ 4       │ 4       │ 20        │ 8        │
└────────┴────────┴────────┴────────┴─────────┴─────────┴───────────┴──────────┘
```

The 20-byte `BodySHA1` is verified on every read (`readOne`) and written by
`WritePacket`; `ProtoFmt`/`ProtoVer` select the wire format (proto2/proto3).

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
| `pkg/ratelimit/` | Rate limiting | `NewLimiter()`, `NewProtoLimiter()` |
| `pkg/retry/` | Retry with backoff | `New()`, `Do()` |
| `pkg/metrics/` | Metrics | counters/histograms |
| `pkg/cache/` | K-line LRU+TTL cache | `NewKLCache()` |
| `pkg/degradation/` | Graceful degradation | manager |
| `pkg/health/` | Health checks | monitoring helpers |
| `pkg/history/` | Historical data helpers | pagination |
| `pkg/market/` | Market hours/session helpers | `IsOpen()` etc. |
| `pkg/option/` | Options analytics helpers | greeks, chains |
| `pkg/tracing/` + `pkg/tracing/otel/` | Tracing | opt-in OTel |
| `pkg/logger/` | Logging | `New()`, `Info()`, etc. |
| `pkg/constant/` | Constants | Typed enums, error codes |
| `pkg/util/` | Utilities | `ParseCode()`, `FormatCode()` |
| `pkg/futuapi/` | Public re-export | `NewClientFromEnv()` |
| `pkg/pb/` | Generated protobuf | 184 packages |
| `internal/client/` | Core TCP | Connection, packet I/O |

---

### 2.3 Error Handling

```go
type FutuError struct {
    Code     ErrorCode
    Message  string
    Category ErrorCategory  // connection, timeout, api, account, trading, subscribe, unknown
    Recovery string         // Suggestion for recovery
}

// Usage
fe, ok := constant.AsFutuError(err)
if ok && fe.Category == constant.CategoryAPI {
    // handle API error
}
```

**Error Categories** (see `pkg/constant/errors.go`):
| Category | Description | Recovery |
|----------|-------------|----------|
| API | Server returned an error | Check `RetType`/`RetMsg`, retry if transient |
| Connection | TCP/socket error | Reconnect |
| Timeout | Request timeout | Retry with backoff |
| Account | Account/authentication problem | Check account ID, unlock, permissions |
| Trading | Order rejected | Check order params |
| Subscribe | Subscription problem | Check subscription state |
| Unknown | Unclassified | Inspect the raw code/message |

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
    constant.TrdMarket_HK,        // trading market
    "00700",                      // code
    constant.TrdSide_Buy,         // side
    constant.OrderType_Normal,    // order type
    350.0,                        // price
    100,                          // quantity
    constant.TrdSecMarket_HK,     // security market
)
```

**OrderBuilder pattern** (`Build` returns the request and an error):
```go
req, err := trd.NewOrder(accID, constant.TrdMarket_HK, constant.TrdEnv_Simulate).
    Buy("00700", 100).
    At(350.0).
    Build()
if err != nil {
    return err
}
```

---

### 4.3 Push APIs

**Channel-based (recommended):**
```go
ch := make(chan *push.UpdateBasicQot, 100)
stop, err := chanpkg.SubscribeQuote(ctx, cli, constant.Market_HK, "00700", ch)
if err != nil {
    return err
}
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
    return "[REDACTED]"
}

// Usage: the value is redacted in all fmt output.
req := &trd.UnlockTradeRequest{
    PwdMD5: constant.SensitiveString("actual_password"),
}
fmt.Printf("%v", req) // Prints: {PwdMD5:[REDACTED]}
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
// The client pool lives in internal/client/pool.go and is intended for
// in-module use (it is not part of the public client API).
pool := futuapi.NewClientPool(cfg)
cli, err := pool.Get(ctx, futuapi.PoolTypeMarketData)
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
// pkg/ratelimit — token bucket. rate and capacity are caller-supplied.
limiter := ratelimit.NewProtoLimiter(100, 100, ratelimit.ModeWait) // 100 req/s, burst 100
// Wired into a client via the internal option (internal/client.WithRateLimiter);
// it is not yet re-exported on the public client package.
```

### 8.5 Custom Logger
```go
// Route SDK logs through a caller-supplied *slog.Logger.
cli, err := client.New(client.WithSlogLogger(slog.Default()))
```

---

## 9. Dependencies

### Direct (go.mod)
```
github.com/gorilla/websocket            v1.5.3    // WebSocket transport
github.com/prometheus/client_golang     v1.20.5   // metrics bridge
go.opentelemetry.io/otel                v1.43.0   // tracing
go.opentelemetry.io/otel/metric         v1.43.0
go.opentelemetry.io/otel/trace          v1.43.0
google.golang.org/protobuf              v1.36.11  // wire messages
```

### Generated (pkg/pb/)
- 184 generated protobuf packages (Futu protocol v10.10.7008)
- All under `github.com/shing1211/futuapi4go/pkg/pb/`

---

## 10. Version Compatibility

| Component | Version |
|------------|---------|
| Go | 1.26.6+ (`go.mod`) |
| OpenD | v10.10.7008 (matches the generated protos; see [docs/VERSION_MAP.md](docs/VERSION_MAP.md)) |
| Protobuf | proto2 + proto3 (mixed on the wire) |

---

## 11. Quick Reference

| Operation | Code |
|------------|------|
| Connect | `cli.Connect("127.0.0.1:11111")` |
| Get quote | `client.GetQuote(ctx, cli, Market_HK, "00700")` |
| Place order | `client.PlaceOrder(ctx, cli, accID, TrdMarket_HK, code, TrdSide_Buy, OrderType_Normal, price, qty, TrdSecMarket_HK)` |
| Subscribe | `client.Subscribe(ctx, cli, Market_HK, "00700", []constant.SubType{constant.SubType_Basic})` |
| Close | `cli.Close()` |

See [README.md](README.md) for complete API reference.