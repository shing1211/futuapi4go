# ADR 0005 — One Go package per `.proto` file

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

Futu publishes 184 `.proto` files. They could be grouped into a few logical Go
packages (`qot`, `trd`, `sys`, …), or generated one package per file.

## Decision

Generate **one Go package per `.proto` file**, named after the lowercased file
name. For example:

- `api/proto/Qot_GetBasicQot.proto` → `pkg/pb/qotgetbasicqot`
- `api/proto/Trd_PlaceOrder.proto` → `pkg/pb/trdplaceorder`
- shared types: `common`, `qotcommon`, `trdcommon`

Generated via `scripts/regen-all-protos.sh`; never edited by hand.

## Consequences

- Import paths are predictable and collision-free, and a proto change touches
  exactly one generated package.
- It is easy to tell which protocol a working tree carries by counting
  `pkg/pb/**/*.pb.go` (see [ADR-0010](0010-version-map-authority.md)).
- The import list for a file that uses many APIs is long; hand-written wrappers
  in `pkg/{qot,trd,sys}` and `client/` keep this out of user code.
