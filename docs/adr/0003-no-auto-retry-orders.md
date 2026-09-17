# ADR 0003 — Never auto-retry order mutations

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

Order operations (`PlaceOrder`, `ModifyOrder`, `CancelOrder`, `ReconfirmOrder`)
are **not idempotent**. A request that times out may still have reached OpenD
and been executed. Automatically retrying it can place a duplicate order, double
a position, or cancel an order the caller did not intend to cancel.

## Decision

The SDK's retry/backoff machinery is **not** applied to order-mutation calls.
Retries are permitted for read-only calls (quotes, positions, account queries)
and for transient connection/timeout errors on non-mutating requests.

## Consequences

- Order calls may return a timeout even when the operation succeeded; callers
  must reconcile by querying order state (`GetOrderList`) before retrying.
- Contributors must not "helpfully" wrap order calls in a retry. This is
  enforced by review and stated in
  [CONTRIBUTING.md](../../CONTRIBUTING.md) and the pull-request template.
- Read paths still benefit from retry/backoff and the circuit breaker.
