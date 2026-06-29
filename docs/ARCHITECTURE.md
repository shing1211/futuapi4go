# futuapi4go Architecture

> **Version:** v0.14.0 | **Futu Protocol:** v10.8.6808 | **Updated:** 2026-06-29

---

## 1. Overview

futuapi4go is a Go SDK for [Futu OpenD](https://www.futunn.com/en/overview) — a TCP-based trading and market data gateway. All communication uses Protocol Buffers over raw TCP sockets (no HTTP/REST/JSON).

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
- All protobuf fields accessed via direct nil-checks (no GetXxx helpers)
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
cli.OnQuote(func(q *push.UpdateBasicQot) { fmt.Println(q.CurPrice) })
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
| `breaker/` | Circuit breaker — trips after 5 failures in 10s, half-open probes |
| `ratelimit/` | Per-protoID rate limiter — burst of 10, refill 100ms |
| `retry/` | Exponential backoff — base 500ms, max 30s, jitter |
| `metrics/` | Request latency, success/failure counts, reconnect count |
| `tracing/otel/` | OpenTelemetry integration (opt-in) |
| `cache/` | LRU + TTL cache for kline data |
| `degradation/` | Graceful degradation under load |

### 2.6 Protobuf Definitions (`pkg/pb/`)

167 generated protobuf files matching Futu OpenD v10.8.6808 protocol.

```
pkg/pb/
├── common/           (RetType, TrdEnv, TrdMarket, etc.)
├── initconnect/      (C2S, S2C — connection handshake)
├── keepalive/        (Ping/pong)
├── qot/              (79 market data protos: GetBasicQot, GetKL, Subscribe, etc.)
├── trd/              (Trading protos: PlaceOrder, CancelOrder, GetOrderList, etc.)
├── sys/              (System protos: GetGlobalState, GetUserInfo)
└── getglobalstate/  (Server state response)
```

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
    ConnectWithRSA(addr, rsaPublicKeyPEM)        internal/client/client.go:558
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

**Key files:** `internal/client/client.go:558-715`, `internal/client/conn.go:105-120`

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
    client.requestInternal()                     internal/client/client.go:1080
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

**Key files:** `client/quote_api.go`, `internal/client/client.go:1080`, `internal/client/conn.go:178-221`

### Flow 3: PlaceOrder (Trading)

```
Application
    cli.Trade().PlaceOrder(ctx, trdEnv, market, code, side, qty, price)
         │
         ▼
    trd.PlaceOrder()                              pkg/trd/place_order.go
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

**Key files:** `pkg/qot/subscribe.go`, `pkg/push/qot_push.go`, `internal/client/client.go:672-685`

### Flow 5: Reconnection & Keep-Alive

```
readLoop()                                        internal/client/client.go:795
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
  reconnect()                                     internal/client/client.go:840
     │
     ├─ Check reconnecting flag (atomic CAS)
     ├─ Backoff: 3s → 6s → 12s → 30s (max)
     ├─ conn.Dial(addr) again
     ├─ re-init AES session key
     └─ Restart readLoop + keepAliveLoop

keepAliveLoop(interval)                           internal/client/client.go:734
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
│   ├── pb/                    Generated Protocol Buffer code (78 files)
│   │   ├── common/            Shared enums (RetType, TrdEnv, Market, etc.)
│   │   ├── initconnect/      InitConnect handshake
│   │   ├── keepalive/        Keep-alive ping
│   │   ├── qot/              79 market data protos
│   │   └── trd/              Trading protos
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
│   └── proto/                .proto source files (167 protos, Futu v10.8.6808)
│
├── test/                      Integration tests, benchmarks, fixtures
│   ├── integration/          Live OpenD tests (requires running OpenD)
│   ├── qot_api/              Market data unit tests
│   ├── trd_api/              Trading unit tests
│   ├── util/                 Mock server for testing
│   ├── benchmark/            Performance benchmarks
│   └── fixtures/             Test fixtures (HSI symbol data)
│
├── docs/                      Documentation
│   └── CHANGELOG.md          Release history
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

    subgraph GeneratedPB["pkg/pb/ — Generated Protobuf (78 files)"]
        PB1[common/<br/>RetType TrdEnv Market]
        PB2[initconnect/<br/>Handshake]
        PB3[qot/ 5001-5999<br/>79 Market Data Protos]
        PB4[trd/ 2201-2299<br/>Trading Protos]
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
            R2[Rate Limiter<br/>10 burst 100ms refill]
            R3[Retry<br/>500ms base 30s max]
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

| SDK Version | Proto Version | Notable Changes |
|-------------|---------------|-----------------|
| v0.14.0 | v10.8.6808 | Latest — 167 protos, 56 new v10.8 APIs (search, indicators, options analytics, rankings, institutional, chain, heatmap, market fundamentals) |
| v0.9.0 | v10.5.6508 | Latest — 78 protos |
| v0.5.7 | v10.5.6508 | Upgrade from v10.4.6408 |
| v0.5.0 | v10.4.6408 | Context as first param, typed enums |
| v0.4.0 | v10.3.5808 | Initial stable release |

---

## 8. Key Design Principles

1. **Binary over TCP** — No HTTP/REST/JSON. Pure Protobuf serialization.
2. **Context as first param** — All public APIs accept `context.Context` as first argument.
3. **No GetXxx helpers** — Direct nil-checks on proto fields: `if field != nil { val = *field }`.
4. **Concurrent reads** — `readLoop()` goroutine reads packets; response dispatched by SerialNo.
5. **Graceful degradation** — Circuit breaker, rate limiter, retry, and cache all opt-in.
6. **Thread-safe** — All shared state protected by `sync.Mutex` or `sync/atomic`.
7. **No goroutine leaks** — All goroutines have exit via `done` channel or `WaitGroup`.