# Error Handling & Error Codes

All SDK errors returned from API calls are `*constant.FutuError`, which carries
a code, a human-readable message, the originating function, an optional wrapped
cause, a category, and a recovery hint.

```go
type FutuError struct {
    Code     ErrorCode
    Message  string
    Func     string
    Err      error          // wrapped inner error, if any
    Category ErrorCategory
    Recovery string
}
```

## Inspecting an error

```go
import (
    "errors"
    "github.com/shing1211/futuapi4go/pkg/constant"
)

if err != nil {
    // errors.As gives the structured error.
    if fe, ok := constant.AsFutuError(err); ok {
        log.Printf("code=%s (%d) category=%s: %s",
            fe.CodeString(), fe.Code, fe.Category, fe.Message)
        log.Printf("hint: %s", fe.Recovery)
    }

    // Or use the convenience predicates.
    switch {
    case constant.IsInsufficientBalance(err):
        // top up / reduce size
    case constant.IsMarketClosed(err):
        // wait for the session
    case constant.IsSubscriptionError(err):
        // subscribe first
    case constant.IsTimeoutError(err):
        // retry with backoff
    }
}
```

- `constant.AsFutuError(err) (*FutuError, bool)`
- `constant.CategoryOf(err) ErrorCategory`
- `constant.RecoveryHint(err) string`
- `FutuError.Unwrap()` / `FutuError.Is(target)` — `errors.Is`/`errors.As` work
- Predicates: `IsSuccess`, `IsInvalidParams`, `IsTimeout`/`IsTimeoutError`,
  `IsDisconnected`, `IsNetworkError`, `IsServerBusy`, `IsServerError`,
  `IsAPIError`, `IsAccountError`, `IsInsufficientBalance`, `IsMarketClosed`,
  `IsOrderRejected`, `IsSubscriptionError`, `IsConnectionError`,
  `IsTradingError`

## Categories

| Category | Constant | Meaning |
|----------|----------|---------|
| `api` | `constant.CategoryAPI` | Server returned an application error |
| `connection` | `constant.CategoryConnection` | TCP/socket problem |
| `timeout` | `constant.CategoryTimeout` | Request timed out |
| `account` | `constant.CategoryAccount` | Account/authentication problem |
| `trading` | `constant.CategoryTrading` | Order rejected or trading unavailable |
| `subscribe` | `constant.CategorySubscribe` | Subscription state problem |
| `unknown` | `constant.CategoryUnknown` | Unclassified |

## Codes

| Value | Constant | Category | Recovery hint |
|------:|----------|----------|---------------|
| `0` | `ErrCodeSuccess` | api | — |
| `-1` | `ErrCodeInvalidParams` | api | Check function parameters for validity |
| `-100` | `ErrCodeTimeout` | timeout | Increase timeout or check network connectivity |
| `-101` | `ErrCodeNetworkError` | connection | Check network connection and retry |
| `-102` | `ErrCodeProtocolErr` | connection | Protocol mismatch — reconnect |
| `-103` | `ErrCodeServerBusy` | api | Server busy — retry after a delay |
| `-200` | `ErrCodeDisconnected` | connection | Reconnect (`Connect()`) |
| `-201` | `ErrCodeAccNotFound` | account | Verify account ID and trading category |
| `-202` | `ErrCodeAccDisabled` | account | Account disabled — contact broker |
| `-203` | `ErrCodeAccLocked` | account | Unlock trading first |
| `-204` | `ErrCodeAccAuthFail` | account | Verify trading password |
| `-301` | `ErrCodeInsufficientBalance` | trading | Check available buying power |
| `-302` | `ErrCodeMarketClosed` | trading | Wait for market to open |
| `-303` | `ErrCodeOrderRejected` | trading | Check order parameters and market rules |
| `-304` | `ErrCodePriceOutOfRange` | trading | Adjust price within the allowed range |
| `-305` | `ErrCodeQtyTooLarge` | trading | Reduce order quantity |
| `-306` | `ErrCodeTradingDisabled` | trading | Trading not enabled for this account |
| `-307` | `ErrCodeInvalidSecurity` | trading | Verify stock code format |
| `-308` | `ErrCodeNoPermission` | trading | Check API subscription level |
| `-400` | `ErrCodeUnknown` | unknown | Inspect the raw code/message |
| `-401` | `ErrCodeAlreadySubbed` | subscribe | Already subscribed |
| `-402` | `ErrCodeNotSubbed` | subscribe | Subscribe to the data feed first |

The authoritative definitions live in
[`pkg/constant/errors.go`](../pkg/constant/errors.go); this table mirrors them.

## Retry guidance

Retry is safe for **read** operations and transient connection/timeout/busy
errors. **Never** auto-retry order-mutation calls (`PlaceOrder`, `ModifyOrder`,
`CancelOrder`, `ReconfirmOrder`) — see
[CONTRIBUTING.md](../CONTRIBUTING.md) and the trading-safety notes in
[DESIGN.md](../DESIGN.md).
