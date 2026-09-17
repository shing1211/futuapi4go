# ADR 0002 — `context.Context` is the first parameter

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

Every API call performs network I/O against OpenD and can block or be
cancelled. Callers need a uniform way to set deadlines, cancel work, and
propagate trace context.

## Decision

Every public function that performs I/O takes a `context.Context` as its
**first** parameter, before the client and request. Examples:

```go
func GetQuote(ctx context.Context, c *Client, market constant.Market, code string) (*Quote, error)
func PlaceOrder(ctx context.Context, c *Client, accID uint64, ...) (*PlaceOrderResult, error)
```

## Consequences

- Cancellation and deadlines propagate to the request/response wait.
- Callers must supply a context (use `context.Background()` for simple cases).
- New APIs that omit the context are considered non-conforming.
