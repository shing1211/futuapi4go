# ADR 0008 — Wrap secrets in `constant.SensitiveString`

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

The SDK handles secrets: the RSA private key PEM (for `InitConnect`) and the
trading-password MD5 (`PwdMD5`). If these are stored as ordinary strings they
can leak into logs or error output via `fmt` formatting of a struct.

## Decision

Sensitive fields use the `constant.SensitiveString` type, whose `String()`,
`GoString()`, and `Format()` methods all yield `[REDACTED]`. The real value is
obtained explicitly with `Raw()`. Applies to:

- `trd.UnlockTradeRequest.PwdMD5` / `trd.Account.PwdMD5`
- `internal.ClientOptions.RSAPrivateKey` (via `WithRSAPrivateKey`)

## Consequences

- Accidentally printing an options or request struct does not reveal the
  secret.
- Code that needs the underlying value must call `.Raw()`, which makes the
  disclosure explicit and greppable.
- See [DESIGN.md §6](../../DESIGN.md#6-security-model).
