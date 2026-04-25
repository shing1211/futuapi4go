# Changelog

All notable changes to this project are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- **All API functions now accept `context.Context` as first parameter** — enables request cancellation, timeouts, and deadline propagation across all `pkg/qot`, `pkg/trd`, and `pkg/sys` functions
- **`AGENTS.md` completely rewritten** — comprehensive operational guide with session workflow, phase gates, code review checklist, and troubleshooting
- **`IMPLEMENTATION_PLAN.md` updated with 24-item roadmap** — full-spectrum quality enhancement plan across 5 phases

### Added

- **`wrapError` helper** — standardized error messages across all API functions (`%s failed: retType=%d, retMsg=%s`)

### Fixed

- **Nil pointer guards** — all list iteration loops now check for nil elements before dereferencing
- **Input validation** — all API functions now validate required fields before sending requests

## [0.1.0] - 2026-04-25

### Added

- **Futu OpenAPI proto v10.4.6408 upgrade** — full proto regeneration from v10.2 → v10.4
- **`Extra` field in `GetUserInfoResponse`** — passthrough information field for future API extensibility
- **`ShQotRight` / `SzQotRight` fields in `GetUserInfoResponse`** — Shanghai/Shenzhen market quote permissions

### Changed

- **`UserAttribution_MM` comment updated** — from "美国" to "MooMoo" (branding alignment)
- **`QotRight_Level3` enum removed** — deprecated Level3 quote right removed from protos

### Fixed

- All proto `go_package` paths now correctly reference project module path (`github.com/shing1211/futuapi4go/pkg/pb/...`)

### Maintenance

- 78 proto files regenerated with proper package directory structure
- `go build ./...` and `go vet ./...` both pass cleanly

## [0.0.7] - 2026-04-24

### Added

- **`chanpkg.SubscribeKLines(cli, market, code, map[KLType]func(*UpdateKL))`** — subscribe to multiple K-line periods with type-safe per-period callbacks; replaces both the map-of-channels and callback variants


- **`constant.KLType` enum values** — were scrambled (SubType values used instead of KLType values); OpenD sends `KlType=6` for 5min, `KlType=1` for 1min, etc. — constants now match proto wire values

## [0.0.6] - 2026-04-23 — Feature Parity Achieved

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
