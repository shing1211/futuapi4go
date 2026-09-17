# futuapi4go Architecture

> **Version:** v0.19.2 | **Futu Protocol:** v10.10.7008 | **Updated:** 2026-09-17
>
> See [VERSION_MAP.md](VERSION_MAP.md) for the authoritative protocol/tag mapping.

---

## 1. Overview

futuapi4go is a Go SDK for [Futu OpenD](https://www.futunn.com/en/overview) — a TCP-based trading and market data gateway. It speaks Protocol Buffers over raw TCP sockets, with an optional WebSocket transport (`ConnectWS`/`ConnectWSS`) for TLS/proxy environments. There is no HTTP/REST/JSON API. The SDK is a pure library — there is no `main` package or binary.

```
Application
  └── client/Client         (Public wrapper API — RECOMMENDED)
       ├── pkg/qot/         (Market data: quotes, klines, orderbook, brokers)
       ├── pkg/trd/         (Trading: orders, positions, account info)
       ├── pkg/sys/         (System: global state, user info, health)
       ├── pkg/push/        (Push notification parsers)
       └── internal/client/  (Connection management, reconnection, keep-alive)
            ├── Conn         (TCP I/O, packet framing, buffering)
            ├── pool.go      (Connection pool for multi-account)
            ├── rsa.go       (RSA encryption for InitConnect)
            └── aes.go       (AES session encryption)
                 └── Futu OpenD (TCP socket)
```

**Design Constraints:**
- Protocol Buffers over TCP — no JSON
- Context passed as first parameter to all public APIs
- Protobuf fields are read through nil-safe helpers (`util.ProtoStr`, `util.ProtoInt32`, `util.ProtoUint64`, …) for explicit nil-safety
- Thread-safe connection management with automatic reconnection
- Typed error handling via `FutuError` with recovery suggestions

---

## 2. Functional Areas

### 2.1 Connection Layer (`internal/client/`)

The core networking stack that handles all communication with Futu OpenD.

| File | Responsibility |
|------|----------------|
| `client.go` | Main `Client` struct — connection lifecycle, serial numbers, request dispatch, reconnect loop, keep-alive heartbeat, handler registry, metrics |
| `conn.go` | `Conn` struct — raw TCP socket, packet read/write with 44-byte header (Magic="FT", ProtoID, SerialNo, BodyLen, SHA1), dispatch table for responses |
| `ws.go` | WebSocket transport — `wsConn` for TLS/ws connections, separate read/write loops |
| `pool.go` | `Pool` struct — thread-safe pool of `Client` instances keyed by account/trade env |
| `rsa.go` | `RSAEncrypt` — PKCS1v15 RSA encryption, chunked for key sizes > body (same padding as Futu Python/C++ SDKs) |
| `aes.go` | `AESEncrypt`/`AESDecrypt` — AES-128-CBC session encryption after InitConnect |
| `errors.go` | `FutuError` type, error codes (`Code*`), `IsConnectionError()` detection, `wrapError()` helper |
| `slog.go` | Structured logging integration |
| `alloc.go` | `sync.Pool` for hot-path buffer allocations |

**Connection Lifecycle:**
```
1. New() → creates Client with options
2. Connect(addr) → TCP dial → InitConnect → AES key exchange
3. readLoop() goroutine → reads packets → dispatches to handlers
4. keepAliveLoop() goroutine → ping every 30s
5. On disconnect → reconnect() → backoff retry with MaxRetries
```

### 2.2 Market Data (`pkg/qot/`)

All market data APIs and real-time subscription support.

| File | Purpose |
|------|---------|
| `get_basic_qot.go` | `GetBasicQot()` — snapshot quote (last price, high/low, volume) |
| `get_kl.go` | `GetKLines()` — historical klines, `RequestHistoryKL()` |
| `get_order_book.go` | `GetOrderBook()` — bid/ask levels |
| `get_ticker.go` | `GetTicker()` — tick data |
| `get_broker.go` | `GetBroker()` — broker queue |
| `get_stock_filter.go` | `StockFilter()` — scan by criteria |
| `subscribe.go` | `Subscribe()` — register for real-time push |
| `qot_push.go` | Push notification parsers: `UpdateBasicQot`, `UpdateKL`, `UpdateOrderBook`, `UpdateTicker`, `UpdateBroker` |

**Subscription Model:**
```go
// Channel-based (recommended)
ch := make(chan *push.UpdateBasicQot, 100)
stop, err := chanpkg.SubscribeQuote(ctx, cli, market, code, ch)
defer stop()
for q := range ch { fmt.Println(q.CurPrice) }

// Callback-based
cli.OnQuote(func(q *client.PushQuote) error { fmt.Println(q.CurPrice); return nil })
```

### 2.3 Trading (`pkg/trd/`)

Trading APIs with pre-flight validation and simulate/real env switching.

| File | Purpose |
|------|---------|
| `place_order.go` | `PlaceOrder()` with pre-flight checks (market hours, price sanity, lot size) |
| `cancel_order.go` | `CancelOrder()` |
| `get_order_list.go` | `GetOrderList()` — today's orders |
| `get_order_fill_list.go` | `GetOrderFillList()` — today's fills |
| `get_history_order_list.go` | `GetHistoryOrderList()` — historical orders |
| `get_history_order_fill_list.go` | `GetHistoryOrderFillList()` — historical fills |
| `get_position_list.go` | `GetPositionList()` — current positions |
| `get_acc_list.go` | `GetAccList()` — all accounts |
| `get_trade_date.go` | `RequestTradeDate()` — trade dates for security |
| `get_trade_fee.go` | `GetOrderFee()` |
| `get_margin_ratio.go` | `GetMarginRatio()` |
| `trd_push.go` | Push parsers: `UpdateOrder`, `UpdateFill`, `UpdatePosition` |

**Trade Environment:**
```go
cli := client.New().WithTradeEnv(constant.TrdEnv_Simulate) // default
cli := client.New().WithTradeEnv(constant.TrdEnv_Real)      // live trading
```

### 2.4 System (`pkg/sys/`)

System-level APIs for connection state and health checks.

| File | Purpose |
|------|---------|
| `get_global_state.go` | `GetGlobalState()` — server time, market status, connection info |
| `get_user_info.go` | `GetUserInfo()` — user ID, account list |
| `keep_alive.go` | `KeepAlive()` ping |
| `health.go` | Health checks for monitoring |

### 2.5 Resilience (`pkg/`)

Cross-cutting concerns wired into every API call.

| Package | Purpose |
|---------|---------|
| `breaker/` | Circuit breaker — opens after 5 consecutive failures, 30s cooldown, half-open probe |
| `ratelimit/` | Per-protoID token bucket — caller-supplied rate/capacity (opt-in, no defaults) |
| `retry/` | Exponential backoff — base 500ms, max 10s, 3 attempts |
| `metrics/` | Request latency, success/failure counts, reconnect count |
| `tracing/otel/` | OpenTelemetry integration (opt-in) |
| `cache/` | LRU + TTL cache for kline data |
| `degradation/` | Graceful degradation under load |

### 2.6 Protobuf Definitions (`pkg/pb/`)

**184** generated protobuf packages matching Futu OpenD **v10.10.7008** (Futu
protocol). There is **one Go package per `.proto` file**, named after the
lowercased file name — e.g. `api/proto/Qot_GetBasicQot.proto` →
`pkg/pb/qotgetbasicqot/`. Shared message types live in their own `*common`
packages.

```
pkg/pb/
├── common/           shared wire types  (Common.proto)
├── qotcommon/        Qot shared types   (Qot_Common.proto)
├── trdcommon/        Trd shared types   (Trd_Common.proto)
├── initconnect/      connection handshake
├── keepalive/        ping/pong
├── notify/           server notifications
├── getglobalstate/ · getuserinfo/ · usedquota/ · verification/ · skillwrapapi/ · getdelaystatistics/
├── qotgetbasicqot/ · qotgetkl/ · qotsub/ · qotgetorderbook/ · …   (151 Qot_*.proto → 151 packages)
└── trdplaceorder/ · trdmodifyorder/ · trdgetorderlist/ · …       (22 Trd_*.proto → 22 packages)
```

Regenerated from `api/proto/` via `scripts/regen-all-protos.sh`. See
[VERSION_MAP.md](VERSION_MAP.md) for the protocol ↔ SDK-tag mapping.

---

## 3. Key Execution Flows

### Flow 1: Connect (RSA Encrypted)

```
Application
    client.New(client.WithRSAPublicKey(pem))
         │
         ▼
    Connect(addr)
         │
         ▼
    ConnectWithRSA(addr, rsaPublicKeyPEM)        internal/client/client.go
         │
         ├─ conn.Dial(addr)                      TCP dial (30s timeout)
         │
         ├─ proto.Marshal(InitConnect.Request)    Build C2S packet
         │
         ├─ RSAEncrypt(rsaPublicKeyPEM, body)     PKCS1v15 chunked encryption
         │   └─ packetEncAlgo = 0 (FTAES_ECB)
         │
         ├─ conn.WritePacket(1001, serialNo, body) Write 44-byte header + body
         │
         ├─ spawn readLoop() goroutine           Concurrent packet reader
         │
         └─ conn.ReadResponse(serialNo, 30s)     Wait for S2C response
                   │
                   ▼
            InitConnect.Response
              ├─ RetType == 0 (success)
              ├─ connID, loginUserID, aesKey
              └─ keepAliveInterval
```

**Key files:** `internal/client/client.go`, `internal/client/conn.go`

### Flow 2: GetQuote (Market Data)

```
Application
    cli.GetQuote(ctx, "HK.00700")
         │
         ▼
    client.GetQuote()                            client/quote_api.go
         │
         ▼
    cli.RequestContext(ctx, ProtoID_GetBasicQot, req, &rsp)
         │
         ├─ circuitBreaker.Check()               Skip if open
         ├─ rateLimiter.Acquire()                Wait if exceeded
         │
         ▼
    client.requestInternal()                     internal/client/client.go
         │
         ├─ serialNo := nextSerialNo()           Atomic counter
         ├─ proto.Marshal(req)                   Serialize to bytes
         ├─ conn.WritePacket(protoID, serialNo, body)
         │       │
         │       └─ header[0:2] = "FT"           Magic bytes
         │           header[2:6] = ProtoID       e.g. 5001 for GetBasicQot
         │           header[8:12] = SerialNo    Correlation ID
         │           header[12:16] = BodyLen
         │           header[16:36] = SHA1(body)  Integrity check
         │
         ▼
    conn.ReadResponse(serialNo, timeout)
         │   └── readLoop goroutine reads packets concurrently
         │       readOne() → verify magic + SHA1 → Dispatch(serialNo)
         │
         ▼
    proto.Unmarshal(resp.Body, &rsp)
         │
         ▼
    Return *qot.GetBasicQot.Response (or error)
```

**Key files:** `client/quote_api.go`, `internal/client/client.go`, `internal/client/conn.go`

### Flow 3: PlaceOrder (Trading)

```
Application
    cli.Trade().PlaceOrder(ctx, req *trd.PlaceOrderRequest)
         │
         ▼
    trd.PlaceOrder()                              pkg/trd/orders.go
         │
         ├─ ValidateOrderParams()                Pre-flight checks
         │   ├─ market hours check
         │   ├─ price sanity (limit ≤ 10x last)
         │   ├─ lot size validation
         │   └─ qty > 0
         │
         ├─ Build trd.C2S.PlaceOrderRequest
         │
         ▼
    trd.RequestContext(ctx, ProtoID_PlaceOrder, req, &rsp)
         │
         ├─ Verify CanSendProto(2206)            Trade proto available?
         ├─ Apply TrdEnv (simulate vs real)
         │
         ▼
    Same requestInternal() flow as GetQuote
         │
         ▼
    rsp.S2C.RetType == 0?
      ├─ Yes: OrderID returned, push will arrive on UpdateOrder
      └─ No:  Return FutuError with RetMsg
```

**Key files:** `pkg/trd/place_order.go`, `pkg/trd/trd_push.go`

### Flow 4: Subscribe Real-Time Push

```
Application
    cli.Quote().Subscribe(ctx, market, codes, []ProtoID{5001})
         │
         ▼
    qot.Subscribe()                              pkg/qot/subscribe.go
         │
         ▼
    cli.RequestContext(ctx, ProtoID_SubQot, req, &rsp)
         │
         ▼
    SubscribeACK received
         │
         ▼
    OpenD pushes UpdateBasicQot packets asynchronously
         │
         ▼
    readLoop() goroutine
         │
         ├─ conn.readOne()                       Reads 44-byte header
         │
         ├─ Dispatch(pkt)                        Finds handler by ProtoID
         │       │
         │       └─ pushHandler(pkt)             Registered in ConnectWithRSA
         │
         ▼
    cli.pushHandler(pkt)
         │
         ├─ metrics.RecordPushMessage()
         ├─ proto.Unmarshal(body, &pushMsg)
         │
         ▼
    UpdateBasicQot → chanpkg.SubscribeQuote() writes to channel
         │
         ▼
    Application reads from channel (or callback fires)
```

**Key files:** `pkg/qot/subscribe.go`, `pkg/push/qot_push.go`, `internal/client/client.go`

### Flow 5: Reconnection & Keep-Alive

```
readLoop()                                        internal/client/client.go
     │
     ├─ conn.readOne()                           Blocking read (no deadline)
     │
     ├─ pkt, err := c.conn.readOne()
     │       └─ timeout = 0 (infinite wait)
     │
     ▼
  Error received (e.g., TCP reset)?
     │
     ├─ atomic.StoreInt32(&c.connected, 0)
     ├─ logWarn("connection lost: %v")
     │
     ▼
  reconnect()                                     internal/client/client.go
     │
     ├─ Check reconnecting flag (atomic CAS)
     ├─ Backoff: 3s → 6s → 12s → 30s (max)
     ├─ conn.Dial(addr) again
     ├─ re-init AES session key
     └─ Restart readLoop + keepAliveLoop

keepAliveLoop(interval)                           internal/client/client.go
     │
     ├─ Every 30s (default):
     ├─ conn.WritePacket(ProtoID_KeepAlive, serialNo, body)
     └─ ReadResponse(serialNo, 10s) — failure triggers reconnect
```

---

## 4. Directory Structure

```
futuapi4go/
├── client/                    Public API wrappers (RECOMMENDED entry point)
│   ├── client.go              Main Client type, Connect/Close, high-level methods
│   ├── quote_api.go           Market data: GetQuote, GetKLines, Subscribe, etc.
│   ├── trade_api.go           Trading: PlaceOrder, CancelOrder, GetPositionList, etc.
│   ├── system_api.go          System: GetGlobalState, GetUserInfo, KeepAlive, etc.
│   ├── fluent_api.go          Fluent API: cli.Quote().GetKLines(), cli.Trade().PlaceOrder()
│   ├── types.go               Shared request/response types
│   ├── push_callbacks.go      Callback-based push handlers
│   └── push.go                Push registration helpers
│
├── pkg/                       Business logic & protobuf-generated code
│   ├── qot/                   Market data API implementations
│   │   ├── get_basic_qot.go
│   │   ├── get_kl.go
│   │   ├── get_order_book.go
│   │   ├── get_ticker.go
│   │   ├── get_broker.go
│   │   ├── get_stock_filter.go
│   │   ├── subscribe.go
│   │   └── qot_push.go        Push parsers
│   │
│   ├── trd/                   Trading API implementations
│   │   ├── place_order.go
│   │   ├── cancel_order.go
│   │   ├── get_order_list.go
│   │   ├── get_order_fill_list.go
│   │   ├── get_history_order_list.go
│   │   ├── get_position_list.go
│   │   ├── trd_push.go        Push parsers (UpdateOrder, UpdateFill)
│   │   └── queries.go
│   │
│   ├── sys/                   System APIs
│   │   ├── get_global_state.go
│   │   └── get_user_info.go
│   │
│   ├── push/                  Push notification parsers
│   │   ├── qot_push.go        Market data pushes
│   │   └── trd_push.go        Trade pushes
│   │
│   ├── pb/                    Generated protobuf — 184 packages (one per .proto)
│   │   ├── common/ qotcommon/ trdcommon/    Shared wire types
│   │   ├── initconnect/ keepalive/ notify/   Protocol
│   │   ├── qotgetbasicqot/ qotgetkl/ …       151 Qot packages
│   │   └── trdplaceorder/ trdmodifyorder/ …  22 Trd packages
│   │
│   ├── constant/              Typed enums, error codes, constants
│   ├── breaker/               Circuit breaker (circuitbreaker pattern)
│   ├── ratelimit/             Per-protoID rate limiter (token bucket)
│   ├── retry/                 Exponential backoff with jitter
│   ├── metrics/               Request latency, success/failure tracking
│   ├── cache/                 LRU + TTL cache for hot data
│   ├── logger/                Structured logging
│   ├── degradation/           Graceful degradation under load
│   ├── tracing/otel/          OpenTelemetry spans (opt-in)
│   └── futuapi/               Public-facing API (NewClientFromEnv)
│
├── internal/                  Private implementation details
│   └── client/
│       ├── client.go         Core Client (connection lifecycle, serial numbers, dispatch)
│       ├── conn.go           TCP socket, packet I/O, 44-byte header
│       ├── ws.go             WebSocket transport
│       ├── pool.go           Connection pool
│       ├── rsa.go            RSA PKCS1v15 encryption for InitConnect
│       ├── aes.go            AES-128-CBC session encryption
│       ├── errors.go         FutuError, error codes, wrapError()
│       ├── slog.go           Structured logging
│       └── alloc.go          sync.Pool for buffer recycling
│
├── api/                       Protocol definitions
│   └── proto/                .proto source files (184 protos, Futu v10.10.7008)
│
├── test/                      Integration tests, benchmarks, fixtures
│   ├── integration/          Live OpenD tests (requires running OpenD)
│   ├── qot_api/              Market data unit tests
│   ├── trd_api/              Trading unit tests
│   ├── util/                 Mock server for testing
│   ├── benchmark/            Performance benchmarks
│   └── fixtures/             Test fixtures (HSI symbol data)
│
├── docs/                      Documentation (ARCHITECTURE, USAGE, VERSION_MAP, index.html, plans)
│   └── CHANGELOG.md           (note: the CHANGELOG lives at the repo root)
│
├── client/client_test.go     Unit tests for client
└── Makefile                  build, test, release targets
```

---

## 5. Mermaid Architecture Diagram

```mermaid
%%{init: {'theme': 'base', 'themeVariables': { 'fontSize': '14px'}}}%%
flowchart TB
    subgraph Application["Application Layer"]
        A[("User Code")]
    end

    subgraph PublicAPI["client/ — Public API (Recommended Entry Point)"]
        B[Client<br/>Connect · Close<br/>Quote() · Trade() · Sys()]
        B1[Fluent API<br/>cli.Quote().GetKLines()<br/>cli.Trade().PlaceOrder()]
    end

    subgraph BusinessLogic["pkg/ — Business Logic"]
        subgraph Qot["pkg/qot/ — Market Data"]
            Q1[get_basic_qot.go<br/>GetQuote]
            Q2[get_kl.go<br/>GetKLines RequestHistoryKL]
            Q3[get_order_book.go<br/>GetOrderBook]
            Q4[subscribe.go<br/>Subscribe]
            Q5[qot_push.go<br/>UpdateKL UpdateOrderBook]
        end

        subgraph Trd["pkg/trd/ — Trading"]
            T1[place_order.go<br/>PlaceOrder pre-flight]
            T2[cancel_order.go<br/>CancelOrder]
            T3[get_order_list.go<br/>GetOrderList]
            T4[get_position_list.go<br/>GetPositionList]
            T5[trd_push.go<br/>UpdateOrder UpdateFill]
        end

        subgraph Sys["pkg/sys/ — System"]
            S1[get_global_state.go<br/>GetGlobalState]
            S2[get_user_info.go<br/>GetUserInfo]
        end

        subgraph CrossCutting["Cross-Cutting"]
            C1[breaker/<br/>Circuit Breaker]
            C2[ratelimit/<br/>Rate Limiter]
            C3[retry/<br/>Exponential Backoff]
            C4[metrics/<br/>Latency Tracking]
        end
    end

    subgraph GeneratedPB["pkg/pb/ — Generated Protobuf (184 packages)"]
        PB1[common/ qotcommon/ trdcommon/<br/>Shared wire types]
        PB2[initconnect/ keepalive/ notify/<br/>Protocol]
        PB3[qot*/<br/>151 Qot packages]
        PB4[trd*/<br/>22 Trd packages]
    end

    subgraph Core["internal/client/ — Core TCP Stack"]
        CL[Client<br/>RequestContext<br/>readLoop keepAlive<br/>reconnect]
        CN[Conn<br/>Dial WritePacket<br/>readOne Dispatch]
        WS[ws.go<br/>WebSocket TLS]
        RS[rsa.go<br/>RSAEncrypt InitConnect]
        AS[aes.go<br/>AES Session Encrypt]
        ER[errors.go<br/>FutuError wrapError]
    end

    subgraph OpenD["Futu OpenD"]
        OD[("TCP Socket<br/>Protobuf")]

        subgraph Resiliences["Resilience"]
            R1[Circuit Breaker<br/>5 failures → open]
            R2[Rate Limiter<br/>opt-in token bucket]
            R3[Retry<br/>500ms base 10s max]
        end
    end

    A --> B
    A --> B1
    B --> Qot
    B --> Trd
    B --> Sys

    B1 --> Q1 & Q2 & Q3 & Q4 & T1 & T2 & T3 & T4 & S1 & S2

    Q1 & Q2 & Q3 & Q4 & T1 & T2 & T3 & T4 & S1 & S2
        --> CL
        --> CN
        --> RS
        --> AS
        --> OD

    CL --> CL
    CL --> R1 & R2 & R3
    CN --> WS

    Q5 & T5 --> CL

    style OD fill:#e1f5fe
    style Core fill:#fff3e0
    style BusinessLogic fill:#f1f8e9
    style PublicAPI fill:#fce4ec
```

---

## 6. Package Map (Public API)

| Use Case | Package | Key Functions |
|----------|---------|---------------|
| Connect to OpenD | `client/` | `New()`, `Connect()`, `ConnectWS()`, `Close()` |
| Quote | `client/` | `GetQuote()`, `GetKLines()`, `GetOrderBook()`, `GetTicker()` |
| Subscribe Push | `pkg/push/chan/` | `SubscribeQuote()`, `SubscribeKLine()`, `SubscribeOrderBook()` |
| Trade | `pkg/trd/` | `PlaceOrder()`, `CancelOrder()`, `GetOrderList()`, `GetPositionList()` |
| System | `pkg/sys/` | `GetGlobalState()`, `GetUserInfo()`, `KeepAlive()` |
| Errors | `pkg/constant/` | `AsFutuError()`, `ErrorCategory*` constants |
| Circuit Breaker | `pkg/breaker/` | `New()`, `Do()` |
| Rate Limiter | `pkg/ratelimit/` | `New()`, `Acquire()` |
| Retry | `pkg/retry/` | `New()`, `Do()` |

---

## 7. Protocol Version History

The authoritative table lives in [VERSION_MAP.md](VERSION_MAP.md). Summary:

| SDK Version | Proto Version | Notable Changes |
|-------------|---------------|-----------------|
| v0.19.2 | v10.10.7008 | Latest — LICENSE/module-doc/release-workflow housekeeping |
| v0.19.0 | v10.10.7008 | `clientVer` 1100, `SubType` enum aligned to the wire, CI (build/vet/race/gofmt/docs) |
| v0.18.0 | v10.10.7008 | Human-readable opt-in packet logging, injectable `slog` logger, `handshakeClientVer` constant |
| v0.16.0 | v10.10.7008 | Protocol upgrade (184 protos) |
| v0.15.0 | v10.9.6908 | 184 protos, 17 Event Contract / Prediction Market APIs |
| v0.14.0 | v10.8.6808 | 167 protos, 56 new v10.8 APIs |
| v0.11.0 | v10.6.6608 | 104 protos |
| v0.9.0 | v10.5.6508 | 78 protos |
| v0.5.0 | v10.4.6408 | Context as first param, typed enums |

---

## 8. Key Design Principles

1. **Binary over TCP** — No HTTP/REST/JSON. Pure Protobuf serialization.
2. **Context as first param** — All public APIs accept `context.Context` as first argument.
3. **Nil-safe proto access** — New code reads proto fields through `util.ProtoStr` / `util.ProtoInt32` / `util.ProtoUint64` helpers instead of dereferencing pointers directly.
4. **Concurrent reads** — `readLoop()` goroutine reads packets; response dispatched by SerialNo.
5. **Graceful degradation** — Circuit breaker, rate limiter, retry, and cache all opt-in.
6. **Thread-safe** — All shared state protected by `sync.Mutex` or `sync/atomic`.
7. **No goroutine leaks** — All goroutines have exit via `done` channel or `WaitGroup`.