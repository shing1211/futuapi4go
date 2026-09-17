# ADR 0007 — Public `client` wrapper over the internal client

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

The networking core (connection lifecycle, dispatch, reconnect, crypto) should
not be part of the supported public API, but users want a stable, ergonomic
surface. Some low-level helpers in `pkg/*` are written against the core client
type.

## Decision

- The core client lives in `internal/client` (import path
  `github.com/shing1211/futuapi4go/internal/client`), which Go prevents external
  modules from importing.
- The public surface is the `client` package: `client.New(...)`,
  package-level helpers such as `client.GetQuote` / `client.PlaceOrder`, and the
  fluent `cli.Quote()` / `cli.Trade()` / `cli.Sys()` APIs.
- `pkg/futuapi` re-exports convenience constructors such as
  `NewClientFromEnv()`.
- `client.Client.Inner()` exposes the underlying internal client for the few
  `pkg/*` helpers that require it (e.g. `cache.NewKLCachedClient`,
  `trd.PlaceOrder`).

## Consequences

- External code cannot depend on the unstable core directly, so it can evolve.
- Documentation must consistently show the public `client` API; examples that
  import `internal/client` are wrong by construction.
- `Inner()` is the escape hatch for advanced/in-module helpers and should be
  used sparingly.
