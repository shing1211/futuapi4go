# Architecture Decision Records

This directory records the significant, durable decisions behind futuapi4go.
Each record captures the context, the decision, and its consequences so future
changes don't silently reverse a deliberate choice.

Format: [MADR](https://adr.github.io/madr/)-style, one file per decision.

| ADR | Decision | Status |
|-----|----------|--------|
| [0001](0001-pure-library.md) | Ship a pure library (no binaries) | Accepted |
| [0002](0002-context-first.md) | `context.Context` is the first parameter | Accepted |
| [0003](0003-no-auto-retry-orders.md) | Never auto-retry order mutations | Accepted |
| [0004](0004-protobuf-over-tcp.md) | Protobuf over TCP, optional WebSocket, no HTTP/JSON | Accepted |
| [0005](0005-one-package-per-proto.md) | One Go package per `.proto` file | Accepted |
| [0006](0006-nil-safe-proto-access.md) | Nil-safe proto access via `util.Proto*` helpers | Accepted |
| [0007](0007-public-client-wrapper.md) | Public `client` wrapper over the internal client | Accepted |
| [0008](0008-sensitive-string.md) | Wrap secrets in `constant.SensitiveString` | Accepted |
| [0009](0009-release-from-changelog.md) | Releases are published by `gh` from the CHANGELOG | Accepted |
| [0010](0010-version-map-authority.md) | `docs/VERSION_MAP.md` is the version authority | Accepted |
| [0011](0011-license-and-protocol-provenance.md) | Apache-2.0; protos derived from published definitions | Accepted |

## Adding a record

Copy the structure of an existing record, use the next number, and add it to
the table above. Decisions that change a prior ADR should mark the old one
`Superseded by ADR-NNNN` rather than deleting it.
