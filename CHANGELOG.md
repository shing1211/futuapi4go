# Changelog

All notable changes to this project are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- `Plate.PlateType` field added to `pkg/qot/plate.go` and `client/types.go` — was missing from the SDK wrapper layers despite being present in the proto `PlateInfo`.

## [0.8.2] - 2026-05-17

### Changed

- `GetTradeDate` (`pkg/qot/trade_date.go`) now uses proper typed `GetTradeDateRequest`/`GetTradeDateResponse` structs instead of raw proto types. `client/quote_api.go` `GetTradeDates` updated to use new wrapper.

- `SubscribeKLines`/`SubscribeKLine` (`pkg/push/chan/chan.go`) now accept `constant.Market` and `[]constant.KLType` instead of raw `int32`/`[]int32`. `klTypeToSubType` now returns an error on unknown KLType instead of silently defaulting.

- `GetFunds` (`pkg/trd/position.go`) validation: split combined nil/zero check into separate checks with distinct error messages.

- Input validation added to `GetBasicQot` (empty security list), `GetOrderBook` (nil Security), `GetTicker` (nil Security), `GetRT` (nil Security), `GetKL` (nil Security, ReqNum > 0).

- `Funds.IsPDT` and `Funds.PDTSeq` fields now have GoDoc comments explaining PDT meaning.

- Removed unused `ProtoID_GetHistoryKLPoints` constant from `pkg/qot/kline.go` and unused `ProtoID_GetMarketSnapshot` alias from `pkg/qot/quote.go`.

### Fixed

- `klTypeToSubType` returns error on unknown KLType instead of silently mapping to `SubType_K_1Min`.

## [0.8.1] - 2026-05-17

### Changed

- `RT` struct (`pkg/qot/market_data.go`, `client/types.go`) now exposes all 9 `TimeShare` proto fields: `Time`, `Minute`, `IsBlank`, `Price`, `LastClosePrice`/`LastClose`, `AvgPrice`/`AvgPrice`, `Volume`, `Turnover`, `Timestamp`. Previously `Minute`, `IsBlank`, and `Timestamp` were omitted.

- `BasicQot` struct (`pkg/qot/quote.go`) now exposes all 25 proto fields. Added: `ListTime`, `PriceSpread`, `DarkStatus`, `ListTimestamp`, `UpdateTimestamp`, `SecStatus`, `OptionExData`, `PreMarket`, `AfterMarket`, `FutureExData`, `Overnight`. Nested optional messages are exposed as raw proto pointers (`*qotcommon.Xxx`).

- `Ticker` struct (`pkg/qot/market_data.go`) now includes `PushDataType` field.

- `Broker` struct (`pkg/qot/market_data.go`) now includes `OrderID` field (SF market support).

- `KLine` struct (`pkg/qot/quote.go`, `pkg/qot/kline.go`) now includes `TurnoverRate` field.

- `Position` struct (`pkg/trd/position.go`) now includes `PositionSide` field.

- `Funds` struct (`pkg/trd/position.go`) now includes `SecuritiesAssets`, `FundAssets`, `BondAssets` fields.

- `UpdateBasicQot` struct (`pkg/push/qot_push.go`) now includes `IsSuspended`, `LastClosePrice`, `UpdateTime`, `UpdateTimestamp` fields.

- `Quote` type (`client/types.go`) now includes `IsSuspended`, `SecStatus` fields.

- `Position` type (`client/types.go`) now includes `PositionSide` field.

- `Funds` type (`client/types.go`) now includes `SecuritiesAssets`, `FundAssets`, `BondAssets` fields.

- `KLine` type (`client/types.go`) now includes `IsBlank` field.

- `MaxTrdQtysInfo` type (`client/types.go`) now includes `LongRequiredIM`, `ShortRequiredIM` fields.

- `MarginRatioInfo` type (`client/types.go`) now includes `ShortPoolRemain`, `AlertLongRatio`, `AlertShortRatio`, `McmLongRatio`, `McmShortRatio`, `MmLongRatio`, `MmShortRatio` fields.

- `PushTicker` type (`client/types.go`) now includes `Sequence`, `Dir`, `RecvTime`, `Type`, `TypeSign` fields.

- `PushRT` type (`client/types.go`) now includes `Minute`, `IsBlank`, `Timestamp` fields.

- `BrokerItem` type (`client/types.go`) now includes `Name` field.

## [0.8.0] - 2026-05-17

### Added

- **Rate limiter wired** — `ratelimit.ProtoLimiter` now actually invoked in `requestInternal`/`requestContextInternal` (previously dead code).
- **Retry wired** — `retry.Config` wraps `Request()`/`RequestContext()` with exponential backoff and recoverable error filtering.
- **WebSocket auto-reconnect** — `reconnect()` now handles WS connections (previously TCP-only).
- **SkillWrapAPI wrappers** — `GetTechnicalUnusual()`, `GetFinancialUnusual()`, `GetDerivativeUnusual()` in `pkg/sys/skill_wrap.go`.
- **Qot_GetTradeDate wrapper** — `GetTradeDate()` (proto 3225) in `pkg/qot/trade_date.go`.
- **Connection state machine** — `ConnState` enum (`Disconnected`→`Connecting`→`Connected`→`Reconnecting`→`Closing`), `State()` method, `OnStateChange` callback.
- **Graceful Shutdown** — `Shutdown(timeout)` drains in-flight requests before closing, `ErrClientClosing` sentinel.
- **Fluent API completion** — 30+ new client-level convenience wrappers covering remaining qot/trd/sys functions.
- **OpenTelemetry metrics bridge** — `pkg/tracing/otel/metrics.go` with 8 sync instruments + 3 observable gauges.
- **K-Line data cache** — `pkg/cache/kl_cache.go` with LRU eviction + TTL expiry.
- **Structured slog integration** — `logInfo`/`logWarn`/`logError` route through `SlogLogger` when configured.
- **Order pre-flight validation** — `pkg/trd/validation.go` with market-open, buying-power, and max-qty checks.
- **Audit/compliance logging** — `pkg/trd/audit.go` with `AuditLogger` for `PlaceOrder`, `ModifyOrder`, `ReconfirmOrder`.
- **Verification tests** — 37 new tests across rate limiter, retry, breaker, state machine, shutdown, KL cache, validation, audit, OTel metrics, slog integration.

### Changed

- **ENHANCEMENT_PLAN.md** — Stripped to only core SDK work (Phases A–F removed to separate repo).
- **docs/USAGE.md** — Added 7 new sections: connection state machine, graceful shutdown, K-line cache, order validation, audit logging, OTel metrics, structured logging.

## [0.7.0] - 2026-05-16

### Added

- **OpenTelemetry tracing** — New `pkg/tracing/otel/` package (opt-in adapter) plus wiring in `internal/client/client.go` for `requestContextInternal`, `ConnectWithRSA`, `connectWebSocket`, `reconnect`, `Close`, and both TCP/WS push handlers. 19 tests pass across `pkg/tracing` and `pkg/tracing/otel`.
- **`.goreleaser.yaml`** — Release automation config with empty builds (pure library), changelog, and GitHub release publisher.
- **`BoolAttr`** helper added to `pkg/tracing/tracing.go`.
- **Trilingual package documentation** — Created `doc.go` for all 22 public packages with English, Simplified Chinese (zh-CN), and Traditional Chinese (zh-TW) documentation.

### Changed

- **README.md** — Complete rewrite: added TOC, feature highlights, categorized features, troubleshooting, and contributing section. Reduced from 330 to 280 lines with cleaner structure.
- **docs/USAGE.md** — Rewritten with advanced patterns (Fluent API, order builder, circuit breaker, auto-paginated K-lines), troubleshooting section, and parity between English and Chinese (Traditional) sections.
- **Makefile** — Added `release` target (goreleaser) and phony declarations.
- **AGENTS.md** — Updated with goreleaser release process.

## [0.6.2] - 2026-05-16

### Added

- **SHA1 validation test** — `internal/testutil/mock/sha1_test.go` with two
  tests that empirically prove: (1) the SDK sends SHA1(ciphertext), (2) both
  SHA1(plaintext) and SHA1(ciphertext) are accepted by a spec-compliant server
  (OpenD is lenient), (3) the official Python SDK
  (`futu/common/utils.py:_joint_head()`) computes SHA1(plaintext) and works
  correctly with OpenD in production.

### Fixed

- **`WritePacketEncrypted` comments** — Updated to accurately state the SHA1
  situation: OpenD accepts both SHA1(plaintext) per the official Python SDK
  and SHA1(ciphertext) per our v0.5.15 testing. We keep SHA1(ciphertext) for
  historical compatibility.

## [0.6.1] - 2026-05-16

### Fixed

- **`WritePacketEncrypted` parameter name** — Renamed `plaintextSHA1` to
  `encryptedBodySHA1` in `ConnInterface`, `Conn.WritePacketEncrypted`, and
  `wsConn.WritePacketEncrypted`. Updated comments to accurately reflect that
  the parameter receives SHA1 of the encrypted body.

## [0.6.0] - 2026-05-16

### Added

- **Phase 1: File Splits** — `client/client.go` (3,393L → 349L) split into 9 focused files; `pkg/qot/quote.go` (2,767L → 238L) split into 14 files; `pkg/trd/trade.go` (1,695L → 171L) split into 7 files.
- **Phase 2a: Util Package** — `pkg/util/{price,date,security,crypto,json}.go` providing PricePrecision, FormatPrice, NewSecurity, MD5Hex, ToJSON, ToCSV, etc. 88 util tests.
- **Phase 2b: JSON Tags + Slice Methods** — `json:"camelCase"` tags on all ~280 fields across 64 structs in `client/types.go`; 27 slice types with `ToJSON()`, `ToCSV()`, `Filter()` in `client/slice_methods.go`.
- **Phase 3a: Mock OpenD Server** — `internal/testutil/mock/` with FT-protocol, RSA/AES encryption, handler registry, request logging, `fillNilPointers`. Replaces `test/util/mock_server.go`.
- **Phase 3b: Typed Push Callbacks** — `OnQuote()`, `OnOrder()`, `OnOrderFill()`, `OnKLine()`, `OnOrderBook()`, `OnTicker()`, `OnRT()`, `OnBroker()`, `OnPriceReminder()`, `OnTrdNotify()` — chainable callbacks on `*Client`.
- **Phase 3b: `WithEnvConfig()` option** — reads `FUTU_OPEND_ADDR`, `FUTU_RSA_PUBLIC_KEY`, `FUTU_RSA_PRIVATE_KEY`, `FUTU_ENCRYPT`, `FUTU_LOG_LEVEL` from environment. PEM values support both file paths and inline strings.
- **Phase 4: Convenience Re-export** — `pkg/futuapi` with `NewClient()`, `NewClientFromEnv()` and re-exported constants.
- **Documentation** — `docs/USAGE.md` (bilingual EN/CN), `client/example_test.go` (runnable examples), updated `ENHANCEMENT_PLAN.md`.

### Changed

- **Client struct** now holds a `*callbackState` for typed push callbacks. `WithTradeEnv`/`WithTradeMarket` remain copy-safe.
- **FixupResponse** in mock server corrected protoID mappings (1002→GetGlobalState, 1004→KeepAlive).
- **FixupResponse** now fills nil pointers even for unregistered protoIDs (custom handlers).
- **Mock server** now only AES-decrypts when the client explicitly requests encryption via `PacketEncAlgo != -1`.

### Fixed

- **AES ECB encrypts/decrypts only 16 bytes** — `block.Encrypt`/`Decrypt` processes exactly one AES block. Any protobuf body >16 bytes is silently corrupted. Fixed with a loop: `for i := 0; i < len(padded); i += bs`.
- **Operator precedence in `inferSecMarket`** — `&&` binds tighter than `||`. Code ≤3 chars with `.SZ` suffix panics with index-out-of-bounds. Fixed with parentheses.
- **isEncrypt unprotected reads/writes** — 8 locations accessed `isEncrypt` (int32) without atomic operations. Fixed with `atomic.LoadInt32`/`StoreInt32`.
- **aesKey read without mutex** — `EncryptRequestBody`/`DecryptResponseBody` read `c.aesKey` without lock protection. Fixed with new `getAESKey()` mutex accessor.
- **Nil safety: GetOrderBook** — Iterated nil `ob` or `d` items from proto lists. Fixed with nil guards.
- **Nil safety: GetFunds** — `s2c.GetFunds()` could be nil, panics on first field access. Fixed with nil check.
- **Push handler deregistration race** — `RegisterHandler(protoID, nil)` writes nil to handler map while concurrent push dispatch may read it. Fixed: handler checks a closed stop channel instead.
- **wrapError in pkg/qot/quote.go** — 3 functions used `fmt.Errorf` instead of `wrapError`: GetBasicQot, GetKL, GetOrderBook nil checks. Fixed.
- **GetStaticInfo overly strict validation** — `req.Market == 0` rejected requests even when `securityList` was provided. Proto says `securityList` takes precedence. Fixed to only require market when no securities provided.
- **RSA private key PEM warning** — `RSAEncrypt` accepted private key PEM (extracted public key internally). Added runtime `logf` warning that this is for testing/backward compat only.
- **WithContext shared ClientOptions** — Deep-copied `ClientOptions` to prevent caller mutations from affecting original client.
- **MockServer protocol handshake** — `MockHandler` signature changed to `func(req []byte) (proto.Message, error)`. Server now calls `fillNilPointers` + `proto.Marshal` on the returned message, enabling proto2 required-field auto-fill. 17 handler registrations updated across 3 test files.

### Added (pre-v0.6.0)

- **`WithEncryption(enable)` option** — opt-in FTAES_ECB encryption for all packets after InitConnect (matching Python SDK's `SysConfig.enable_proto_encrypt(True)`). When enabled, the InitConnect response is RSA-decrypted to extract the AES key, and all subsequent communication uses FTAES_ECB encryption with that key.
- **`WithRSAPrivateKey(pem)` option** — sets the RSA private key PEM for decrypting InitConnect responses. Required when `WithEncryption(true)` is used. The public key is extracted automatically, so `WithRSAPublicKey` is optional when a private key PEM is provided.
- **`RSADecrypt()` function** — RSA decryption using PKCS#1 v1.5 (matching the custom padding scheme used by `RSAEncrypt`). Decrypts the InitConnect response body when FTAES encryption is negotiated.
- **57 new unit tests** — Added tests for: AES ECB round-trip and edge cases (pkg/metrics, pkg/ratelimit, pkg/retry, pkg/degradation, pkg/health, pkg/history, pkg/tracing, pkg/breaker).
- **Push handler test coverage** — Mock push handler race detection, goroutine leak tests.
- **fillNilPointers helper** — Reflection-based auto-filler for nil proto2 pointer fields in mock server.
- **`OrderBuilder.WithTrailType(t)`** — sets the trailing stop type (Ratio or Amount) on the fluent OrderBuilder.
- **`OrderBuilder.WithTrailValue(v)`** — sets the trailing stop value (percentage or dollar amount) on the fluent OrderBuilder.
- **`OrderBuilder.WithSpread(s)`** — sets the trailing stop spread on the fluent OrderBuilder.

### Fixed (pre-v0.6.0)

- **RSA connections default to no FTAES encryption** — The SDK previously forced `isEncrypt=1` whenever an RSA public key was configured, causing the server to receive undecryptable FTAES-encrypted GetGlobalState requests.

## [0.5.16] - 2026-05-15

### Fixed

- **Lost response when timer fires before packet arrival** — `ReadResponseContext` used a bare `select` where `timer.C` and `ctx.Done()` could fire before the packet arrived at the dispatch channel. When the timer won the race, the buffered packet was discarded. Fix: each "terminal" case (timeout, ctx-cancelled) now checks the channel with a non-blocking receive before returning.

## [0.5.15] - 2026-05-15

### Fixed

- **AES request SHA1 mismatch** — When sending encrypted API requests (non-InitConnect), the packet header SHA1 was computed over the **plaintext** body, but Futu OpenD verifies it against the **encrypted** body. Every encrypted API call failed with "The packet body SHA1 signature is incorrect." Fix: compute `sha1.Sum(encBody)` instead of `sha1.Sum(body)`.

## [0.5.14] - 2026-05-15

### Fixed

- **FTAES encrypt: empty payload produces invalid ciphertext** — `ftaesEncrypt` set `padLen=0` for empty plaintext (0%16=0 → padLen=16, then wrongly reset to 0), producing 16-byte ciphertext without a 16-byte trailer. Fix: only apply the "already aligned" shortcut when plaintext length > 0.
- **FTAES decrypt: panic on unencrypted responses** — `DecryptResponseBody` tried to AES-decrypt all non-InitConnect responses when `isEncrypt=1`, but Futu OpenD does not encrypt response bodies even when RSA key exchange was used. A 59-byte unencrypted response caused `crypto/aes: input not full block`. Fix: validate FTAES ciphertext format (encrypted portion must be n*16 bytes) before decrypting; return plaintext if format is invalid.

## [0.5.13] - 2026-05-14

### Fixed

- **AES ECB encrypts/decrypts only 16 bytes** — `block.Encrypt`/`Decrypt` processes exactly one AES block. Any protobuf body >16 bytes is silently corrupted. Fixed with a loop: `for i := 0; i < len(padded); i += bs`.
- **Operator precedence in `inferSecMarket`** — `&&` binds tighter than `||`. Code ≤3 chars with `.SZ` suffix panics with index-out-of-bounds. Fixed with parentheses.
- **isEncrypt unprotected reads/writes** — 8 locations accessed `isEncrypt` (int32) without atomic operations. Fixed with `atomic.LoadInt32`/`StoreInt32`.
- **aesKey read without mutex** — `EncryptRequestBody`/`DecryptResponseBody` read `c.aesKey` without lock protection. Fixed with new `getAESKey()` mutex accessor.
- **Nil safety: GetOrderBook** — Iterated nil `ob` or `d` items from proto lists. Fixed with nil guards.
- **Nil safety: GetFunds** — `s2c.GetFunds()` could be nil, panics on first field access. Fixed with nil check.
- **Push handler deregistration race** — `RegisterHandler(protoID, nil)` writes nil to handler map while concurrent push dispatch may read it. Fixed: handler checks a closed stop channel instead.
- **wrapError in pkg/qot/quote.go** — 3 functions used `fmt.Errorf` instead of `wrapError`: GetBasicQot, GetKL, GetOrderBook nil checks. Fixed.
- **GetStaticInfo overly strict validation** — `req.Market == 0` rejected requests even when `securityList` was provided. Proto says `securityList` takes precedence. Fixed to only require market when no securities provided.
- **RSA private key PEM warning** — `RSAEncrypt` accepted private key PEM (extracted public key internally). Added runtime `logf` warning that this is for testing/backward compat only.
- **WithContext shared ClientOptions** — Deep-copied `ClientOptions` to prevent caller mutations from affecting original client.
- **MockServer protocol handshake** — `MockHandler` signature changed to `func(req []byte) (proto.Message, error)`. Server now calls `fillNilPointers` + `proto.Marshal` on the returned message, enabling proto2 required-field auto-fill. 17 handler registrations updated across 3 test files.

### Changed

- **All Subscribe helpers accept context** — `Subscribe`, `SubscribeRehab`, `SubscribeNoData`, `Unsubscribe`, `UnsubscribeRehab`, `UnsubscribeAll`, `RegQotPush`, `UnregQotPush` now accept `ctx context.Context` as first parameter.
- **Test infrastructure timeout** — `WithAPITimeout` in test client increased from 5s to 30s to handle race detector overhead.
- **TestTradingWorkflow_Complete** — Added `ClearRequests()` before workflow to avoid counting registration-phase requests.

### Added

- **57 new unit tests** — Added tests for: AES ECB round-trip and edge cases (pkg/metrics, pkg/ratelimit, pkg/retry, pkg/degradation, pkg/health, pkg/history, pkg/tracing, pkg/breaker).
- **Push handler test coverage** — Mock push handler race detection, goroutine leak tests.
- **fillNilPointers helper** — Reflection-based auto-filler for nil proto2 pointer fields in mock server.

## [0.5.7] - 2026-05-11

### Changed

- **Futu API v10.5.6508 upgrade** — Updated proto files, regenerated Go pb files, updated clientVer from 10100 to 1005, updated README badge

## [0.5.2] - 2026-04-28

### Added

- **GetHistoryKLPoints Wrapper** — `pkg/qot/history_kl_points.go` provides historical K-line data at specific time points using ProtoID 3106
- **UsedQuota Wrapper** — `pkg/sys/system.go` GetUsedQuota() returns used subscription and historical K-line quota using ProtoID 1010
- **NoDataMode Typed Enum** — pkg/qot: NoDataMode_Null, Forward, Backward constants
- **DataStatus Typed Enum** — pkg/qot: DataStatus_Null, Current, Previous, Back constants
- **Fluent API** — client.Quote(), client.Trade(), client.System() for cleaner API access

### Changed

- Updated README with v0.5.2 examples
- Updated ENHANCEMENT_PLAN.md as completed

## [0.5.1] - 2026-04-28

### Added

- **Enhanced Error System** — FutuError now includes Category, Recovery fields, FullMessage(), CodeString(), Is() for errors.Is() compatibility
- **Error Predicates** — IsServerError(), IsAPIError(), IsConnectionError(), IsTimeoutError(), IsTradingError()
- **Error Bridge Functions** — CategoryOf(), RecoveryHint() with unwrap traversal for both FutuError and internal/client.Error
- **Circuit Breaker Integration** — Optional breaker on Client via SetBreaker()/GetBreaker(), auto-wraps Request/RequestContext for business protos
- **Input Validation Standalones** — ValidateAccID(), ValidateCode(), ValidateQty(), ValidatePrice(), ValidateRemark() with boundary checks
- **OrderBuilder AutoDetectMarket** — AutoDetectMarket() method using util.DetectTradingMarkets(), WithSecMarket() method
- **OrderBuilder Build Validation** — Build() now returns (*PlaceOrderRequest, error) with code/qty/market validation
- **Convenience Wrappers** — GetTodayFills(), GetTodayOrders() for today's fill/order queries
- **HistoryKLine Iterator** — TotalFetched(), PageCount() methods, context cancellation check in Next()
- **Typed Enum Int32()** — All typed enums now have Int32() method for implicit proto field conversion
- **TrdMarket.Prefix()** — Returns short market codes ("HK", "US", "CN", "SG", etc.)
- **Client.GetReconnectCount()** — Returns reconnect attempt count from metrics

### Changed

- **wrapError() enhanced** — Both pkg/trd and pkg/qot now map RetType to proper ErrorCode with category/recovery
- **FutuError struct** — Added Category (ErrorCategory) and Recovery (string) fields
- **Prometheus Metrics Integration** — recordRequest/recordReconnect/recordPush now also call metrics.RecordAPICall/RecordReconnect/RecordPushMessage; Connect/Close call RecordConnection/RecordDisconnect/RecordOpenDUp
- **Push handler enhanced** — Records Prometheus push message with ProtoID label
- **Client Public API** — WithTradeEnv/GetTradeEnv now use constant.TrdEnv instead of int32; WithTradeMarket/GetTradeMarket use constant.TrdMarket
- **Connection Pool** — Added Stats() returning map[PoolType]PoolStats and GetPoolType() for O(1) lookup
- **Convenience Wrappers** — Added GetAccountFunds() combining GetFunds in one call
- **TLS Connection Support** — WithTLS() option, Conn.SetTLSConfig(), tls.DialWithDialer with mTLS support
- **Rate Limiter** — pkg/ratelimit: token bucket limiter with ProtoLimiter (global + per-ProtoID), reject/wait modes
- **Retry with Backoff** — pkg/retry: exponential backoff 2^n*base+jitter, ErrorCategory-based recoverability, context-aware
- **Health Checker** — pkg/health: CheckFunc registry, /healthz+/readyz HTTP endpoints, healthy/degraded/unhealthy states
- **New Error Codes** — CodeTLSHandshakeFailed, CodeRateLimited, CodeRetryExhausted, CodeConfigInvalid
- **Client Rate Limiter/Retry** — SetRateLimiter(), SetRetryConfig(), WithRateLimiter(), WithRetryConfig() options
- **Performance Benchmarks** — bench_test.go: BenchmarkNextSerialNo, BenchmarkRecordRequest, BenchmarkPush, BenchmarkPoolGetPut
- **Integration Test Framework** — //go:build integration tag, FUTU_OPEND_ADDR env var, skipWithoutOpenD() helper

### Fixed

- **TestHistoryKLineIteratorNilClient** — Corrected expected behavior for nil client case

## [0.5.0] - 2026-04-27

### Added

- **P6-8: Graceful Shutdown Helpers** — `WaitForSignal()`, `CloseOnSignal()` for signal handling

### Changed

- **P2-1: Type Safety** — 26+ wrapper functions now use typed constants (Market, KLType, WarrantSortField, WarrantType, OptionType, IndexOptionType, HolderCategory, PriceReminderType, PriceReminderFreq, PriceReminderOp, TrdMarket, SubType, RehabType, CapitalFlowPeriodType, ReferenceType, Issuer, WarrantStatus, TrdCashFlowDirection)
  - `GetPlateSet`, `GetIpoList`, `GetFlowSummary`, `GetCapitalFlow`, `GetCapitalDistribution`
  - `RequestHistoryKL`, `RequestHistoryKLWithLimit`, `GetReference`, `GetOwnerPlate`
  - `GetPlateSecurity`, `GetOptionExpirationDate`, `ModifyUserSecurity`, `RequestTradeDate`
  - `GetOptionChain`, `StockFilter`, `GetWarrant`, `Unsubscribe`, `RegQotPush`
  - `SetPriceReminder`, `GetPriceReminder`, `ReconfirmOrder`
  - `GetHoldingChangeList`, `RequestRehab`

### Documentation

- **P6-9: Examples Overhaul** — Updated 17 demo examples to use typed constants
- Updated MIGRATION_GUIDE.md with typed enum examples

## [0.3.1] - 2026-04-26

### Added

- **P2-2: Zero-allocation path** — `sync.Pool` buffer management in internal/client/alloc.go
- **P2-3: Pool O(1) lookup** — `clientIndex` map for fast connection retrieval
- **P2-4: Historical data downloader** — `pkg/history` package for batch data retrieval
- **P2-7: Structured logging** — `internal/client/slog.go` with slog support
- **WebSocket transport infrastructure** — `internal/client/ws.go` (incomplete, not yet working)

### Changed

- **P1-5: Buffered I/O** — Conn now uses `bufio.Reader` for packet reads

## [0.3.0] - 2026-04-25

### Breaking Changes (v0.3.0)

- **P6-1: Context-required API** — All functions now require context as first parameter
- **P6-2: Typed Market Constants** — No more `int32(constant.Market_US)` casts needed

### Added

- **P6-1: Context helpers** — Added `WithTimeout()`, `WithDeadline()` to Client
- **P6-2: Typed Market** — `constant.Market` type for all market parameters
- **P6-3: Enhanced Error Codes** — Added 20+ error codes with predicates
- **P6-5: Bounded Push Channels** — Added buffer size constants & helpers
- **P6-6: Market Detection** — Detect warrants, CBBC, futures from code patterns

### Already Existed

- **P6-4: Configurable Timeouts** — Client.WithTimeout(), WithDeadline()
- **P6-7: Retry Logic** — MaxRetries, ReconnectInterval, ReconnectBackoff

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