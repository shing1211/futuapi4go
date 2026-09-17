# ADR 0010 — `docs/VERSION_MAP.md` is the version authority

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

The SDK's own tags (`v0.x.y`) do not encode the Futu OpenD protocol version.
The protocol advances independently and several files could plausibly carry
"the current version": `go.mod`, `CHANGELOG.md`, `internal/client/version.go`,
the README badges, `docs/IMPLEMENTATION_COMPLETE.md`.

## Decision

`docs/VERSION_MAP.md` is the **single source of truth** for the mapping between
SDK tags and the Futu OpenD protocol (`protos` count, `clientVer`, dates). Other
documents should link to it rather than restating numbers, and must not
duplicate the mapping table.

## Consequences

- Stale numbers in other docs are bugs; when they disagree with VERSION_MAP.md,
  VERSION_MAP.md wins.
- A protocol upgrade updates the mapping table and `handshakeClientVer` in one
  place (see the "Adding a protocol upgrade" section of that file).
- The README badge and other docs should be checked against it on release; a
  doc-drift CI check is a candidate improvement.
