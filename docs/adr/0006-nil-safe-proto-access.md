# ADR 0006 — Nil-safe protobuf access via `util.Proto*` helpers

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

Generated protobuf message fields are pointers. Dereferencing them directly
panics on a nil field, and protobuf getters (`GetXxx()`) hide nil values behind
zero values, which makes nil-safety hard to audit.

## Decision

Hand-written code reads proto fields through explicit nil-safe helpers in
`pkg/util`, e.g. `util.ProtoStr`, `util.ProtoInt32`, `util.ProtoUint64`,
`util.ProtoFloat64`, `util.ProtoBool`. Wrapper code also guards nested messages
(e.g. `if s2c == nil || s2c.Header == nil { return wrapError(...) }`) rather
than chaining getters.

## Consequences

- Nil-handling is explicit and auditable; a missing field yields a typed zero
  instead of a panic.
- New code should prefer the helpers. Legacy wrappers still contain some
  generated `GetXxx()` calls; converting them is incremental and tracked in
  [IMPROVEMENT_PLAN.md](../IMPROVEMENT_PLAN.md).
