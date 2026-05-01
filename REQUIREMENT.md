# futuapi4go Requirements

## Project Overview

| Attribute | Value |
|----------|-------|
| **Project Name** | futuapi4go |
| **Type** | Go SDK (Software Development Kit) |
| **Version** | v0.5.4 |
| **Purpose** | Type-safe, production-ready Go SDK for Futu OpenAPI — market data, trading, real-time push |
| **Target Users** | Algo traders, quantitative developers, trading bot developers |
| **Platform** | Go 1.26+ (Linux, macOS, Windows) |
| **License** | Apache 2.0 |

---

## 1. Functional Requirements

### 1.1 Connection Management (FR1)

| ID | Requirement | Priority | Acceptance Criteria |
|----|-------------|----------|---------------------|
| FR1.1 | Connect to Futu OpenD via TCP socket | P0 | `Connect("127.0.0.1:11111")` returns nil error on success |
| FR1.2 | Handle InitConnect handshake and AES key exchange | P0 | Connection receives connID, serverVer, aesKey after handshake |
| FR1.3 | Implement auto-reconnect on connection loss | P1 | Client reconnects within configured interval on disconnect |
| FR1.4 | Keep-alive heartbeat mechanism | P0 | Ping/pong maintains connection; detects dead connections within keepAliveInterval |
| FR1.5 | Graceful connection close with proper cleanup | P0 | `Close()` drains goroutines, closes socket, returns nil |

**Dependencies:** Standard library only

---

### 1.2 Market Data (FR2)

| ID | Requirement | Priority | ProtoID | Acceptance Criteria |
|----|-------------|----------|--------|---------------------|
| FR2.1 | Get real-time quotes | P0 | 3004 | `GetBasicQot()` returns CurPrice, Volume, High, Low, Open |
| FR2.2 | Retrieve K-line data | P0 | 3006 | `GetKL()` returns Open, High, Low, Close, Volume |
| FR2.3 | Fetch order book depth | P0 | 3012 | `GetOrderBook()` returns bid/ask price levels |
| FR2.4 | Get tick-by-tick trades | P0 | 3010 | `GetTicker()` returns Time, Price, Volume, Direction |
| FR2.5 | Fetch intraday time-share | P0 | 3008 | `GetRT()` returns Price, AvgPrice, Volume per minute |
| FR2.6 | Get broker queue | P0 | 3014 | `GetBroker()` returns broker Name, Position, Volume |
| FR2.7 | Retrieve static security info | P0 | 3002 | `GetStaticInfo()` returns Name, LotSize, ListDate |
| FR2.8 | Get market state | P0 | 3016 | `GetMarketState()` returns Open/Closed/PreMarket state |
| FR2.9 | Fetch capital flow | P1 | 3018 | `GetCapitalFlow()` returns InFlow, MainInFlow per category |
| FR2.10 | Fetch capital distribution | P1 | 3020 | `GetCapitalDistribution()` returns Large/Medium/Small breakdown |
| FR2.11 | Plate operations | P1 | 3022/3024 | `GetPlateSet()` returns industry/region/concept plates |
| FR2.12 | Option chain and expiration | P1 | 3032 | `GetOptionChain()` returns Call/Put options with strikes |
| FR2.13 | Warrant data | P1 | 3034 | `GetWarrant()` returns warrant data by issuer |
| FR2.14 | Historical K-line download | P0 | 3104 | `RequestHistoryKL()` supports pagination via NextReqKey |
| FR2.15 | Historical K-line at time points | P1 | 3106 | `GetHistoryKLPoints()` returns K-line at specific times |
| FR2.16 | Subscription quota query | P1 | 3108 | `GetSubInfo()` returns quota usage |

---

### 1.3 Trading (FR3)

| ID | Requirement | Priority | ProtoID | Acceptance Criteria |
|----|-------------|----------|--------|---------------------|
| FR3.1 | List trading accounts | P0 | 5002 | `GetAccList()` returns AccID, AccType, TrdEnv |
| FR3.2 | Unlock trading with password | P0 | 5004 | `UnlockTrade()` returns nil on success |
| FR3.3 | Query account funds | P0 | 5006 | `GetFunds()` returns Power, Cash, TotalAssets |
| FR3.4 | Query positions | P0 | 5010 | `GetPositionList()` returns Code, Qty, CostPrice, PnL |
| FR3.5 | Place buy/sell orders | P0 | 5001 | `PlaceOrder()` returns OrderID, OrderIDEx |
| FR3.6 | Modify or cancel orders | P0 | 5003 | `ModifyOrder()` returns updated OrderID/OrderIDEx |
| FR3.7 | Cancel all orders | P0 | 5005 | `CancelAllOrder()` cancels all open orders |
| FR3.8 | Query order list | P0 | 5008 | `GetOrderList()` returns today's orders |
| FR3.9 | Query historical orders | P1 | 5012 | `GetHistoryOrderList()` with date range |
| FR3.10 | Query order fills | P0 | 5014 | `GetOrderFillList()` returns FillID, Price, Qty |
| FR3.11 | Query historical fills | P1 | 5016 | `GetHistoryOrderFillList()` with date range |
| FR3.12 | Get max tradable quantities | P1 | 5020 | `GetMaxTrdQtys()` returns MaxCashBuy, MaxSell |
| FR3.13 | Get order fees | P1 | 5022 | `GetOrderFee()` returns Commission, StampDuty |
| FR3.14 | Get margin ratios | P1 | 5024 | `GetMarginRatio()` for short selling |
| FR3.15 | Get account trading info | P1 | 5026 | `GetAccTradingInfo()` returns margin requirements |
| FR3.16 | Get flow summary | P1 | 5030 | `GetFlowSummary()` returns cash flow entries |
| FR3.17 | Reconfirm order | P1 | 5032 | `ReconfirmOrder()` for risky operations |

---

### 1.4 System APIs (FR4)

| ID | Requirement | Priority | ProtoID | Acceptance Criteria |
|----|-------------|----------|--------|---------------------|
| FR4.1 | Get global state | P0 | 1002 | `GetGlobalState()` returns QotLogined, TrdLogined, ServerVer |
| FR4.2 | Get user info | P0 | 1003 | `GetUserInfo()` returns UserID, NickName, ApiLevel |
| FR4.3 | Get quota usage | P1 | 1010 | `GetUsedQuota()` returns used subscription/HistoryKL quota |
| FR4.4 | Get delay statistics | P2 | 1004 | `GetDelayStatistics()` returns latency (known proto2 issue) |

---

### 1.5 Real-Time Push (FR5)

| ID | Requirement | Priority | ProtoID | Acceptance Criteria |
|----|-------------|----------|--------|---------------------|
| FR5.1 | Subscribe to updates | P0 | 3001 | `Subscribe()` enables real-time data |
| FR5.2 | Unsubscribe | P0 | 3001 | `Unsubscribe()` stops updates |
| FR5.3 | Parse quote push | P0 | 3005 | `ParseUpdateBasicQot()` decodes price updates |
| FR5.4 | Parse K-line push | P0 | 3007 | `ParseUpdateKL()` decodes K-line updates |
| FR5.5 | Parse order book push | P0 | 3013 | `ParseUpdateOrderBook()` decodes depth |
| FR5.6 | Parse ticker push | P0 | 3011 | `ParseUpdateTicker()` decodes trades |
| FR5.7 | Parse RT push | P0 | 3009 | `ParseUpdateRT()` decodes time-share |
| FR5.8 | Parse broker push | P0 | 3015 | `ParseUpdateBroker()` decodes broker queue |
| FR5.9 | Parse order update push | P0 | 2208 | `ParseUpdateOrder()` decodes order status |
| FR5.10 | Parse order fill push | P0 | 2218 | `ParseUpdateOrderFill()` decodes fills |
| FR5.11 | Channel-based delivery | P0 | - | `chanpkg.SubscribeQuote()` returns channel |
| FR5.12 | Batch subscribe | P1 | - | `SubscribeSymbols()` subscribes multiple codes |

---

### 1.6 User Security / Watchlist (FR6)

| ID | Requirement | Priority | ProtoID | Acceptance Criteria |
|----|-------------|----------|--------|---------------------|
| FR6.1 | List security groups | P1 | 2802 | `GetUserSecurityGroup()` returns group names |
| FR6.2 | Get securities in group | P1 | 2804 | `GetUserSecurity()` returns StaticInfo list |
| FR6.3 | Modify group | P1 | 2806 | `ModifyUserSecurity()` adds/removes securities |

---

### 1.7 Price Alerts (FR7)

| ID | Requirement | Priority | ProtoID | Acceptance Criteria |
|----|-------------|----------|--------|---------------------|
| FR7.1 | Get price reminders | P1 | 3102 | `GetPriceReminder()` returns alert list |
| FR7.2 | Set price alerts | P1 | 3104 | `SetPriceReminder()` creates/updates/deletes alerts |

---

## 2. Non-Functional Requirements

### 2.1 Performance

| ID | Requirement | Target | Evidence |
|----|-------------|--------|-----------|
| NR1 | Buffered I/O | 64KB buffer | `bufio.Reader` in conn.go |
| NR2 | Connection pool | Multiple clients | `pkg/pool.go` |
| NR3 | Zero-allocation path | Hot request path | `sync.Pool` in alloc.go |
| NR4 | Pool O(1) lookup | < 1μs | `clientIndex` map |
| NR5 | Request latency | < 10ms typical | Benchmarks in bench_test.go |

---

### 2.2 Reliability

| ID | Requirement | Implementation |
|----|-------------|----------------|
| NR6 | Circuit breaker | `pkg/breaker/` |
| NR7 | Auto-reconnect | Exponential backoff in client.go |
| NR8 | Rate limiting | Token bucket in pkg/ratelimit/ |
| NR9 | Retry with backoff | Exponential retry in pkg/retry/ |

---

### 2.3 Security

| ID | Requirement | Implementation |
|----|-------------|----------------|
| NR10 | Sensitive data protection | `SensitiveString` type redacts in fmt |
| NR11 | No credential storage | Password passed per-call, not stored |
| NR12 | TLS support | `WithTLS()` option |

---

### 2.4 Observability

| ID | Requirement | Implementation |
|----|-------------|----------------|
| NR13 | Structured logging | `pkg/logger/` (text + JSON) |
| NR14 | Prometheus metrics | `metrics.RecordAPICall()` |
| NR15 | Error tracing | FutuError with stack trace |

---

### 2.5 Compatibility

| ID | Requirement | Implementation |
|----|-------------|----------------|
| NR16 | Go version | Go 1.26+ |
| NR17 | OpenD version | v10.4.6408 recommended |

---

## 3. Constraints

- **No WebSocket** — TCP only (ENHANCEMENT_PLAN.md has WebSocket as future)
- **No authentication storage** — passwords passed per-call
- **No order caching** — real-time only
- **No market data persistence** — user responsible for storage

---

## 4. Out of Scope (Future)

| Feature | Phase | Notes |
|---------|-------|-------|
| WebSocket transport | D1 | WebSocket in ENHANCEMENT_PLAN.md |
| TWAP/VWAP/IS | A | Execution algorithms |
| VaR/Greeks | B | Risk engine |
| Strategy framework | C | Event-driven |
| Backtesting | C2 | Historical replay |

---

## 5. Dependencies

### Direct Dependencies

- `google.golang.org/protobuf` — Proto serialization
- Standard library (net, context, sync, etc.)

### Generated Dependencies

- `pkg/pb/*` — Protobuf-generated types (local module)

---

## 6. Quality Standards

- All public APIs require `context.Context` as first parameter
- Input validation at entry (nil checks, boundary checks)
- Error wrapping via `wrapError()` helper
- Goroutine leak prevention (done channels, WaitGroups)
- Race detection tests pass: `go test -race ./...`