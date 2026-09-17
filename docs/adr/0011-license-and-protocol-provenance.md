# ADR 0011 — Apache-2.0 license; protos derived from published definitions

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

The SDK ships generated protobuf bindings (184 packages) that originate from
Futu's protocol definitions, alongside original Go code. The licensing and
provenance of both need to be unambiguous.

## Decision

- The project is licensed **Apache-2.0**. The `LICENSE` file is the **verbatim**
  Apache-2.0 text (the appendix placeholder is not edited, because altering the
  body breaks license detection and the terms themselves).
- The copyright holder is declared in `NOTICE` and in per-file headers.
- `api/proto/` contains protocol definitions derived from Futu's **publicly
  published** OpenAPI documentation; `pkg/pb/` is generated from them. The SDK
  contains no proprietary Futu source code.
- Trademarks ("Futu", "moomoo", …) belong to Futu Holdings; the project is not
  affiliated or endorsed (see `DISCLAIMER.md`).

## Consequences

- Because a modified license body defeats automated detection (pkg.go.dev showed
  `License: UNKNOWN` until it was restored), `LICENSE` must stay verbatim.
- Third-party dependencies and their licenses are listed in `NOTICE`.
- Protocol updates replace `api/proto/` and regenerate `pkg/pb/`; the
  provenance statement continues to hold.
