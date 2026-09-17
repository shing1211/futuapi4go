# ADR 0004 — Protobuf over TCP; optional WebSocket; no HTTP/JSON

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

Futu OpenD speaks a binary protocol: a fixed header followed by a Protocol
Buffers message body. Some deployments need to traverse TLS proxies or
restrictive networks where a raw TCP socket is awkward.

## Decision

- The primary transport is **Protocol Buffers over a raw TCP socket**.
- An **optional WebSocket transport** (`ConnectWS` / `ConnectWSS`) is provided
  for TLS/proxy environments; it carries the same packet framing.
- The SDK does **not** offer an HTTP/REST/JSON API.

## Consequences

- Wire compatibility is with OpenD, not with any REST service; there is no
  OpenAPI schema to generate from.
- The packet framing is shared by both transports and includes a 20-byte
  SHA-1 body checksum (see
  [DESIGN.md](../../DESIGN.md#21-protocol-communication)).
- Users choosing WebSocket supply the WS secret key; TCP users use RSA/AES
  (see [ADR-0008](0008-sensitive-string.md)).
