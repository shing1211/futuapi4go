# Changelog

All notable changes to this project are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.6] - 2026-04-25

### Added (Phase 5 Polish)

- **P5-1: Pagination Iterator** — Add `NewHistoryKLineIterator()` in pkg/qot/iterator.go

### Completed (Previously Existed)

- **P5-2: Unified Client** — client.New() already provides unified API
- **P5-3: GoDoc** — All packages already documented
- **P5-4: ProtoID naming** — Already standardized
- **P5-5: Examples** — README.md with HK, US examples

## [0.2.5] - 2026-04-25

### Completed (Previously Existed)

- **P4-1: Mock Server** — test/util/mock_server.go already implements InitConnect + handlers
- **P4-2: Edge Case Tests** — 46 tests in internal/client, 38 in pkg/trd
- **P4-3: Docker Integration** — futuopend Docker image available

### Added (P4-4)

- **Order validation helpers** — Added LotSize(market), PriceTick(market) in pkg/constant/validation.go

### Changed

- **P3-4: TLS** — Skipped: RSA+AES encryption already sufficient for non-localhost connections

## [0.2.4] - 2026-04-25

### Added

- **P1-6: Input validation** — Added validation to key trading functions (GetFunds, GetPositionList)
- **P1-7: Proto nil checks** — Already handled (nil guards exist in loops, proto3 uses zero values)

### Fixed

- **Validation errors use FutuError** — Consistent error type with error codes

## [0.2.3] - 2026-04-25

### Added

- **P1-5: Buffered I/O** (`internal/client/conn.go`) — Added 64KB bufio.Reader for reduced syscalls
- **P3-3: sync.Pool placeholder** (`pkg/trd/trade.go`) — Added pool definitions for future optimization

## [0.2.2] - 2026-04-25

### Added (Phase 3 Infrastructure)

- **P3-1: FutuError type** (`pkg/constant/errors.go`) — programmatic error handling:
  ```go
  if constant.IsTimeout(err) { /* handle timeout */ }
  fe, ok := constant.AsFutuError(err)
  ```
  Error codes: `ErrCodeSuccess`, `ErrCodeInvalidParams`, `ErrCodeTimeout`, `ErrCodeDisconnected`, `ErrCodeUnknown`
- Helper predicates: `IsTimeout()`, `IsDisconnected()`, `IsInvalidParams()`, `IsSuccess()`, `AsFutuError()`

## [0.2.1] - 2026-04-25

### Added (Phase 2 Ease of Use)

- **P2-2: OrderBuilder** (`pkg/trd/builder.go`) — fluent builder pattern for orders:
  ```go
  trd.NewOrder(accID, market, env).Buy("00700", 100).At(350.5).Build()
  ```
- **P2-3: Convenience wrappers** (`pkg/trd/convenience.go`) — one-liner functions:
  - `QuickBuy()`, `QuickSell()`, `QuickMarketBuy()`, `QuickMarketSell()`
  - `CancelAllOrders()`, `GetPositions()`
- **P2-4: DetectTradingMarkets** (`pkg/util/code.go`) — auto-detect TrdMarket/TrdSecMarket from code

## [0.2.0] - 2026-04-25

### Changed

- **P2-1: Typed enums for all trading API parameters** — all `pkg/trd` request structs now use typed enum types (`constant.TrdMarket`, `constant.TrdEnv`, `constant.TrdSide`, `constant.OrderType`, `constant.ModifyOrderOp`, `constant.TrdCategory`) instead of raw `int32` for compile-time type safety
- **All API functions now accept `context.Context` as first parameter** — enables request cancellation, timeouts, and deadline propagation across all `pkg/qot`, `pkg/trd`, and `pkg/sys` functions
- **`AGENTS.md` completely rewritten** — comprehensive operational guide with session workflow, phase gates, code review checklist, and troubleshooting
- **`IMPLEMENTATION_PLAN.md` updated with 24-item roadmap** — full-spectrum quality enhancement plan across 5 phases

### Added

- **`wrapError` helper** — standardized error messages across all API functions (`%s failed: retType=%d, retMsg=%s`)
- **Race detection tests for connection pool** — `TestPoolConcurrentAccess` and `TestPoolConcurrentGetPutRemove`
- **Packet validation tests** — `TestConnWritePacketEmptyBody` and `TestConnWritePacketBodyTooBig`

### Fixed

- **Nil pointer guards** — all list iteration loops now check for nil elements before dereferencing
- **Input validation** — all API functions now validate required fields before sending requests
- **Packet length overflow check** — `WritePacket()` now validates body size before casting to `uint32` (prevents silent overflow)
- **Empty packet rejection** — `WritePacket()` now rejects empty bodies with `CodeInvalidPacket` error

### Security

- **Connection pool mutex protection verified** — all `ClientPool` methods properly protected with `sync.RWMutex`
- **Sensitive data logging protection** — `UnlockTradeRequest.PwdMD5` now uses `constant.SensitiveString` type which redacts itself in all `fmt` output formats (`%s`, `%v`, `%+v`, `%#v`), preventing accidental password exposure in logs

## [0.0.6] - 2026-04-24

### Added

- **`chanpkg.SubscribeKLines(cli, market, code, map[KLType]func(*UpdateKL))`** — subscribe to multiple K-line periods with type-safe per-period callbacks; replaces both the map-of-channels and callback variants
- **`constant.KLType` enum values** — were scrambled (SubType values used instead of KLType values); OpenD sends `KlType=6` for 5min, `KlType=1` for 1min, etc. — constants now match proto wire values

## [0.0.5] - 2026-04-23 — Feature Parity Achieved

### Added

- **`GetLoginUserID() uint64`** — returns the Futu/NiuNiu user ID logged into OpenD
- **`IsEncrypt() bool`** — returns whether the connection uses AES encryption
- **`CanSendProto(protoID uint32) bool`** — checks if a proto can be sent based on connection state
- **`pkg/breaker`** — circuit breaker pattern for resilient trading
- **`pkg/logger`** — structured leveled logging (text + JSON, Debug/Info/Warn/Error)
- **`pkg/push/chan`** — channel-based push delivery (goroutine-safe, buffered channels)
- **`pkg/util`** — code parsing (`ParseCode`, `FormatCode`, market helpers)
- **`pkg/constant`** — Python-style `String()` methods on all enum types
- **`GetAccountInfo`** — full account info with multi-currency cash (`CashInfoList`) and per-market assets (`MarketInfoList`)
- **`GetFlowSummary`** — account cash flow entries (equivalent to Python's `get_acc_cash_flow`)
- **`GetAccTradingInfo`** — max tradable quantities + margin info (equivalent to `acctradinginfo_query`)
- Extended `Funds` struct with 16 new fields: `CashInfoList`, `MarketInfoList`, `MarginCallMargin`, `IsPDT`, `PDTSeq`, `BeginningDTBP`, `DtCallAmount`, `DtStatus`, `RemainingDTBP`

### Fixed

- `GetDelayStatistics` / `GetTradeDate` — documented as known proto2 wire-format issues; calls skipped gracefully in demo

### Changed

- `go.mod`: `go 1.26.1`, published to `proxy.golang.org` as `v0.0.6`
- `client/client.go`: `GetLoginUserID`, `IsEncrypt`, `CanSendProto` wrappers
- `internal/client/client.go`: `loginUserID`, `isEncrypt` fields stored on connect; new methods added

### Tests

- Unit tests for `pkg/util`, `pkg/constant`, `pkg/logger`, `pkg/breaker` — all pass

## [0.0.5] - 2026-04-21

### Added

- Context-aware request cancellation (`RequestContext()`, `ReadResponseContext()`)
- Waitable connection pool with `context.Context` support
- Push notification handler API

### Changed

- All API functions now accept `context.Context` as first parameter
- `ClientPool.Get()` now requires `context.Context`

## [0.0.4] - 2026-04-19

### Added

- Full proto field mapping audit — 100% field coverage across all 59 wrapper functions
- Proto generation pipeline

## [0.0.3] - 2026-04-18

### Fixed

- Push notification parsers now correctly unmarshal into `Response` wrapper then extract `S2C` (22/22 tests pass)

## [0.0.2] - 2026-04-18

### Fixed

- Push parsers unmarshal directly into `S2C` (matching OpenD push body format)
- `logf()` nil logger panic — eager initialization with `log.Default()`
- Connection state race — `connected bool` → `int32` with atomic operations

## [0.0.1] - 2026-04-12

### Added

- Push notification handler API with 11 handlers
- 100% proto field coverage — all 59 wrapper functions fully mapped
- Automatic pagination for `RequestHistoryKL` via `NextReqKey`