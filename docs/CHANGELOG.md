# Changelog

All notable changes to this project are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.9.7] - 2026-04-24

### Added

- **`chanpkg.SubscribeKLines`** — subscribe to multiple K-line periods, returns `map[constant.KLType]chan *UpdateKL` + stop func
- **`chanpkg.SubscribeKLinesHandler`** — Python-style callback; single callback receives all periods, caller switches on `kl.KlType`

### Changed

- **`chanpkg.SubscribeKLines`** — unknown `KlType` (e.g., simulator sends `0`) routes to first subscribed type

### Fixed

- **simulator** — `handlePushKL` now sets `KlType=1` and `RehabType=0` in the S2C response

## [0.9.0] - 2026-04-23 — Feature Parity Achieved

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

- `go.mod`: `go 1.26.1`, published to `proxy.golang.org` as `v0.9.0`
- `client/client.go`: `GetLoginUserID`, `IsEncrypt`, `CanSendProto` wrappers
- `internal/client/client.go`: `loginUserID`, `isEncrypt` fields stored on connect; new methods added

### Tests

- Unit tests for `pkg/util`, `pkg/constant`, `pkg/logger`, `pkg/breaker` — all pass

## [0.8.0] - 2026-04-21

### Added

- Context-aware request cancellation (`RequestContext()`, `ReadResponseContext()`)
- Waitable connection pool with `context.Context` support
- Push notification handler API

### Changed

- All API functions now accept `context.Context` as first parameter
- `ClientPool.Get()` now requires `context.Context`

## [0.7.0] - 2026-04-19

### Added

- Full proto field mapping audit — 100% field coverage across all 59 wrapper functions
- Proto generation pipeline

## [0.6.2] - 2026-04-18

### Fixed

- Push notification parsers now correctly unmarshal into `Response` wrapper then extract `S2C` (22/22 tests pass)

## [0.6.1] - 2026-04-18

### Fixed

- Push parsers unmarshal directly into `S2C` (matching OpenD push body format)
- `logf()` nil logger panic — eager initialization with `log.Default()`
- Connection state race — `connected bool` → `int32` with atomic operations

## [0.6.0] - 2026-04-12

### Added

- Push notification handler API with 11 handlers
- 100% proto field coverage — all 59 wrapper functions fully mapped
- Automatic pagination for `RequestHistoryKL` via `NextReqKey`

## [0.4.1] - 2026-04-08

### Fixed

- Wrapper structs missing fields causing example compilation failures
- 20/20 example compile tests pass

## [0.4.0] - 2026-04-08

### Added

- Push notification support (serial matching to prevent push/consume collision)
- Client metrics collection (latency, success/failure rates, reconnect count)
- Health check with auto-reconnection
- Options trading APIs (GetOptionChain, GetOptionExpirationDate)

## [0.3.0] - 2026-04-07

### Added

- OpenD Simulator — full TCP server handling 70+ ProtoIDs with realistic mock responses
- Push notification support (7 Qot handlers, 3 Trd handlers)

## [0.2.0] - 2026-04-07

### Added

- All Qot market data APIs (37 APIs)
- All Trd trading APIs (14 APIs)
- All Sys system APIs (4 APIs)
- Protobuf definitions at v10.2.6208

## [0.1.0] - 2026-04-07

### Added

- Initial release — core client, InitConnect, basic protobuf definitions
