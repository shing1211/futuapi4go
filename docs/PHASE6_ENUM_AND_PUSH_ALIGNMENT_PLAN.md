# Phase 6: Enum Alignment, Push Completeness & Client Defects

*Generated: 2026-09-17 | Status: COMPLETE (implemented in v0.19.0)*

## Executive Summary

Work started from a baseline run of `go test ./...`, which was **already failing on
`main`** before any change in this phase: `internal/client` hung for the full test
timeout and `pkg/tracing` panicked. Three independent client defects were found and
fixed, then the audit continued into the hand-rolled enums that mirror the generated
protobuf enums and into the push parsers.

Findings: **3 client defects** (2 hangs, 1 panic), **1 wrong enum value that silently
subscribed to the wrong stream**, **8 missing enum values across two enums**,
**1 mapping that existed only as an unexported duplicate**, **1 case of a shipped
API that could never have worked**, and **1 proto field discarded by a push parser**.

The through-line is the same failure mode in four places: a parameter, value or field
that is accepted and then not honoured — which fails quietly rather than loudly.

---

## Step 1: Client deadlocks and panic (P0)

### 1.1 `Conn.ReadResponse` deadlocked instead of timing out

**File:** `internal/client/conn.go`

```go
timer := time.NewTimer(timeout)
stopped := timer.Stop()          // stops the timer before it can ever fire
defer func() {
    if !stopped { <-timer.C }
}()

select {
case pkt := <-ch:                // normal path
case <-timer.C:                  // unreachable: the timer was stopped above
}
```

`timer.Stop()` returns `true` when the timer had not yet fired, so `stopped` was
always `true` and the drain branch never ran — but the timer was now stopped, so
`timer.C` could never deliver. Any call where no response arrived blocked forever.

**Impact:** five request paths use it — `internal/client/client.go:662, 849, 986,
1397, 1503` — including every API-timeout path. A slow or unresponsive OpenD hung
the caller's goroutine permanently instead of returning `read response: i/o timeout`.
`TestConnNilConnection` exercised this and hung the whole package.

**Fix:** adopt the `defer timer.Stop()` shape already used by `ReadResponseContext`,
which was correct.

### 1.2 `tracing.SetTracer` panicked

**File:** `pkg/tracing/tracing.go`

```go
var defaultTracer atomic.Value
func init()  { defaultTracer.Store(NoopTracer{}) }   // concrete type NoopTracer
func SetTracer(t Tracer) { defaultTracer.Store(t) }  // store of *customTracer
```

`atomic.Value` requires every store to share one concrete type. Storing a value of
type `NoopTracer` and then a different concrete type under the `Tracer` interface
panics: `sync/atomic: store of inconsistently typed value into Value`.

**Impact:** a public API panicked on first use. `TestSetTracer` panicked in the
baseline run. This was Phase 5 step 4.2 ("SetTracer() Race") — the right primitive
reached for with the wrong shape.

**Fix:** `atomic.Pointer[Tracer]`, which stores exactly one concrete type by
construction. `GetTracer` falls back to `NoopTracer{}` rather than dereferencing a
nil pointer.

### 1.3 `ClientPool.Get` ignored context cancellation

**File:** `internal/client/pool.go`

`Get` waited on a `sync.Cond`, which returns only on `Signal`/`Broadcast`. Context
cancellation therefore had no effect: the `ctx.Done()` check at the top of the loop
was unreachable while blocked, so the error the code itself documents —

```go
NewError(CodePoolExhausted, "pool exhausted: context timed out waiting for available connection")
```

— could never be returned. With the pool at `MaxSize` and no `Put` in flight, `Get`
blocked forever regardless of the context, including one with a deadline.

**Impact:** `TestPoolMaxSizeLimit` blocked for the full test timeout, and both
concurrency tests exercise the same exhaustion path (10 goroutines, `MaxSize` 5).

**Fix:** the pool wakes waiters through a broadcast channel (`notify`) that every
state change closes and replaces under `mu`, so `Get` waits in a `select` over the
caller's context, the pool's own context, and that channel. The lock is released
while waiting so `Put`, `Remove` and `Close` can progress. `sync.Cond` is gone;
`wakeLocked()` replaces `Signal`/`Broadcast`. `newClientLocked` still runs under the
lock, preserving the existing dial-while-holding-`mu` behaviour rather than changing
it silently in this phase.

**Note:** `TestPoolMaxSizeLimit` passed `context.Background()` and then asserted an
error on exhaustion — unsatisfiable under a waiting contract, and unsatisfiable by
the test alone, since a bounded context would still have hung before the fix. The
test now uses a bounded context, as its sibling concurrency tests already assumed.

---

## Step 2: `SubType_OrderBookOdd` was the wrong value (P0)

**File:** `pkg/constant/constant.go`

`SubType_OrderBookOdd` was `18`. The wire (`Qot_Common.proto:216-220`) assigns:

| Value | Wire name | Wire meaning |
|-------|-----------|--------------|
| 18 | `SubType_KL_10Min` | 10-minute bars |
| 22 | `SubType_OrderBook_Odd` | 碎股摆盘 — the odd-lot order book |

So subscribing to the odd-lot order book requested 10-minute bars. This is a
wrong-value-in-the-right-shape failure: the request is well formed, the response is
plausible, and nothing errors.

**Fix:** the constant is `22`, and the enum is now listed in wire order so the
correspondence is visible rather than implied.

---

## Step 3: Interval coverage (P1)

The wire defines four intervals neither hand-rolled enum carried:

| Wire | `constant.SubType` | `constant.KLType` |
|------|--------------------|-------------------|
| `KL_10Min` = 18 / `KLType_10Min` = 12 | `SubType_K_10Min` = 18 | `KLType_K_10Min` = 12 |
| `KL_120Min` = 19 / `KLType_120Min` = 13 | `SubType_K_120Min` = 19 | `KLType_K_120Min` = 13 |
| `KL_180Min` = 20 / `KLType_180Min` = 14 | `SubType_K_180Min` = 20 | `KLType_K_180Min` = 14 |
| `KL_240Min` = 21 / `KLType_240Min` = 15 | `SubType_K_240Min` = 21 | `KLType_K_240Min` = 15 |

They were unreachable through the typed API. `KLType.IsValid`, `KLType.String`,
`SubType.String`, `SubType.IsValid` and `IsKLType` now cover them, and
`SubType.IsValid` delegates to `IsKLType` for the K types so there is one list of
K-line subtypes rather than two that can disagree. `qot.SubType` gained the same
five values, so the enum adapters subscribe with can name them.

`client`'s `SubType_*` and `KLType_*` re-exports were also incomplete: they omitted
`KL_3Min`, `KL_Quarter` and `KL_Year`.

---

## Step 4: The mapping was private, so it was duplicated (P1)

`KLType` and `SubType` number **differently** and can never be converted by a
numeric cast:

| Interval | `KLType` | `SubType` |
|----------|----------|-----------|
| 1 minute | 1 | 11 |
| day | 2 | 6 |
| week | 3 | *undefined* |
| month | 4 | 13 |
| year | 5 | 16 |
| 5 minutes | 6 | 7 |

The correct mapping lived only in `pkg/push/chan`'s **unexported** `klTypeToSubType`.
`pkg/history.Streamer`, unable to reach it, hand-rolled its own with a raw cast —
and got every entry wrong (1 minute requested quotes, day requested the order book,
week requested an undefined value, month requested ticks, year requested the
time-share stream).

**Fix:** `constant.KLType.ToSubType()` and `constant.SubType.ToKLType()`
(`pkg/constant/convert.go`) are now the single expression of that correspondence.
`pkg/push/chan` is refactored onto them and its private copy is deleted.

---

## Step 5: `pkg/history.Streamer` was removed (P1)

Beyond the mapping defect above, `Start` registered a handler for protoID `3005`
(`Qot_UpdateBasicQot`, the *quote* push) rather than `3007` (`Qot_UpdateKL`), with an
empty handler body under a comment reading "Parse and dispatch to appropriate
handlers".

Nothing referenced it: an exhaustive search of `futuapi4go` and `futuapi4go-demo`
found `NewStreamer`, `OnKLine` and `Start` only at their definitions.

**Fix:** removed `StreamKLineRequest`, `Streamer`, `NewStreamer`, `OnKLine` and
`Start` (~46 lines). The `Downloader` surface is untouched and keeps its tests.

---

## Step 6: `push.UpdateOrderBook` discarded the book type (P1)

**File:** `pkg/push/qot_push.go`

The wire sends `orderBookType` (`Qot_UpdateOrderBook.proto:19`, field 9) on every
order-book push, distinguishing the whole-lot book (0) from the odd-lot book (1).
Neither the parsed struct nor the parser carried it.

**Impact:** one `SubType_OrderBook` subscription serves both books — `Qot_Sub.C2S`
has no book selector, so selecting odd-lot means subscribing with
`SubType_OrderBook_Odd` (22), and the push type is the only signal of which book
arrived. Discarding it made the two indistinguishable to a consumer. The books hold
different levels, so guessing is not a safe fallback.

**Fix:** `OrderBookType int32` on `push.UpdateOrderBook`, populated with
`util.ProtoInt32` (not a `GetXxx()` getter, per AGENTS.md proto safety), plus the
matching `constant.OrderBookType` enum with `String` and `IsValid`.

---

## Step 7: Drift guard (P1)

Nothing linked the hand-rolled enums to the generated ones, which is how a wrong
value survived. `pkg/constant/wire_drift_test.go` adds:

- `TestSubTypeMatchesWire` — all 22 SubType values against `qotcommon.SubType`
- `TestKLTypeMatchesWire` — all 16 KLType values against `qotcommon.KLType`
- `TestOrderBookTypeMatchesWire` — both values
- `TestToSubTypeRoundTrip` — every interval through `ToSubType`/`ToKLType`
- `TestToSubTypeRejectsNonKLine` — quotes, the books, ticks, time-share and broker queue
- `TestIsValidAgreesWithIsKLType` — the two validity checks cannot drift

Pairs are compared **by value, never by name**: the generated enum spells quarter
`Qurater` (an upstream Futu typo preserved faithfully in `Qot_Common.proto:213`),
while `constant` spells it `Quarter`.

---

## Official Wire Values (source of truth: `api/proto/Qot_Common.proto`)

```protobuf
enum SubType {
  SubType_None = 0;  SubType_Basic = 1;      SubType_OrderBook = 2;
  SubType_Ticker = 4; SubType_RT = 5;         SubType_KL_Day = 6;
  SubType_KL_5Min = 7; SubType_KL_15Min = 8;  SubType_KL_30Min = 9;
  SubType_KL_60Min = 10; SubType_KL_1Min = 11; SubType_KL_Week = 12;
  SubType_KL_Month = 13; SubType_Broker = 14;  SubType_KL_Qurater = 15;
  SubType_KL_Year = 16;  SubType_KL_3Min = 17; SubType_KL_10Min = 18;
  SubType_KL_120Min = 19; SubType_KL_180Min = 20; SubType_KL_240Min = 21;
  SubType_OrderBook_Odd = 22;
}
enum KLType { ... KLType_3Min = 10; KLType_Quarter = 11; KLType_10Min = 12;
              KLType_120Min = 13; KLType_180Min = 14; KLType_240Min = 15; }
enum OrderBookType { OrderBookType_Normal = 0; OrderBookType_Odd = 1; }
```

`Qot_Sub.C2S` has **no** `orderBookType` field: the odd-lot book is requested purely
via `subTypeList` containing `SubType_OrderBook_Odd`, while `Qot_GetOrderBook.C2S`
selects it by parameter. The push then reports which book each message carries.

---

## Risk Assessment

| Step | Risk | Reason |
|------|------|--------|
| 1.1 ReadResponse | **HIGH** | On five production request paths; changes timeout behaviour |
| 1.2 SetTracer | **MEDIUM** | Public API; changes the storage type |
| 1.3 Pool.Get | **MEDIUM** | Concurrency change; lock released while waiting |
| 2. SubType_OrderBookOdd | **HIGH** | Wire value; was referenced nowhere, so no live behaviour change |
| 3-4. Enum coverage and mapping | **LOW** | Additive; existing values unchanged |
| 5. Streamer removal | **LOW** | No references anywhere |
| 6. OrderBookType | **LOW** | Additive field |
| 7. Drift guard | **LOW** | Tests only |

---

## Verification Checklist

- [x] `go build ./...` passes
- [x] `go vet ./...` passes
- [x] `go test ./... -race -count=1` green — was red before this phase
- [x] `internal/client` completes in ~4s, previously hung for the full timeout
- [x] `pkg/tracing` passes 9/9, previously panicked
- [x] `SubType_OrderBookOdd` is 22, matching `qotcommon.SubType_SubType_OrderBook_Odd`
- [x] Every `constant` enum value matches its `qotcommon` counterpart (drift test)
- [x] No reference to `Streamer`/`NewStreamer`/`StreamKLineRequest` remains
- [x] `push.UpdateOrderBook` carries `OrderBookType`
- [x] CHANGELOG updated under `[Unreleased]`
- [x] Tagged `v0.19.0` and pushed to origin/main

---

## Known outstanding

- `internal/testutil/mock` is **intermittently flaky** under `-race` with repeated
  runs (`TestCustomHandler`, `TestKeepAlive`). Reproduced at the pre-change base
  commit with no local modifications, so it predates this phase and is unrelated to
  production code. Not addressed here.
- Several files are not `gofmt`-clean on `main` (`pkg/constant/constant.go`'s ProtoID
  block, `pkg/constant/tostring.go`'s `Int32` helpers, `pkg/push/push_test.go`'s
  import order and line 944). Changes in this phase were kept formatting-neutral;
  a repo-wide `gofmt` pass is left as a separate concern.
- `newClientLocked` performs the TCP dial and `InitConnect` handshake while holding
  the pool mutex (Phase 5 step 5.3 described this as fixed; it is not). Unchanged
  here to avoid conflating it with the `Get` fix.
