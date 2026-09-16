# Futu API Version Map

This SDK's own release tags (`v0.x.y`) do **not** encode the Futu OpenD protocol
version. This table is the authoritative mapping between them.

## Current

| | |
|---|---|
| SDK release | **v0.17.0** |
| Futu OpenD protocol | **v10.10.7008** |
| `clientVer` in handshake | **1100** (`internal/client/client.go`) |
| Generated proto files | **184** (`pkg/pb/**/*.pb.go`) |

## Mapping

| Futu OpenD protocol | protos | `clientVer` | SDK tags | released |
|---------------------|-------:|------------:|----------|----------|
| **v10.10.7008** | 184 | 1100 | `v0.16.0` … `v0.17.0` | 2026-08-14 … 2026-09-07 |
| v10.9.6908 | 184 | 1090 | `v0.15.0` … `v0.15.2` | 2026-08-04 … 2026-08-05 |
| v10.8.6808 | 167 | 1080 | `v0.14.0` … `v0.14.1` | 2026-06-29 |
| v10.7.6708 | 111 | 1076 | `v0.13.0` | 2026-06-05 |
| v10.6.6608 | 104 | 1066 | `v0.11.0` … `v0.12.0` | 2026-05-21 … 2026-05-22 |
| v10.5.6508 | 79 | 1005 | `v0.5.8` … `v0.10.0` | 2026-05-12 … 2026-05-20 |
| pre-v10.5.6508 | 78–79 | 10100 | `v0.0.1` … `v0.5.7` | 2026-04-24 … 2026-05-10 |

Protocol-upgrade points, as recorded in [CHANGELOG.md](../CHANGELOG.md):

| Introduced in | Upgrade |
|---------------|---------|
| `v0.16.0` | v10.9.6908 → **v10.10.7008** (184 protos) |
| `v0.15.0` | v10.8.6808 → v10.9.6908 (167 → 184 protos, +17 Event Contract APIs) |
| `v0.14.0` | v10.7.6708 → v10.8.6808 (111 → 167 protos) |
| `v0.13.0` | v10.6.6608 → v10.7.6708 (104 → 111 protos) |
| `v0.11.0` | v10.5.6508 → v10.6.6608 (79 → 104 protos) |
| `v0.5.7`/`v0.5.8` | → v10.5.6508 |

## Notes

**`clientVer` is not a protocol-version mirror.**

Futu's own Python SDK pins `CLIENT_VERSION = 300` in every release checked
(10.5.6508, 10.7.6708, 10.8.6808, 10.9.6908, 10.10.7008 — all 300). It is a
client-side identifier sent in the `InitConnect` handshake; OpenD does not
require it to track the protocol version.

This SDK instead advances `clientVer` alongside each protocol bump. That is a
**local convention**, kept because it makes the announced version traceable in
OpenD logs. Consequences:

- The value is *not* gated by OpenD, so a stale value does not break anything —
  but it silently misreports which protocol a client was built against.
- `v0.16.0` upgraded the protos to v10.10.7008 but left `clientVer` at 1090.
  Fixed in `v0.17.0`-era unreleased work: the value now lives in a single
  constant (`handshakeClientVer`) so the two handshake paths cannot drift apart
  again, and is set to 1100 for v10.10.7008.

**Proto count is a reliable fingerprint.** For a given protocol version the
number of generated `.pb.go` files is stable, so it is the quickest way to
confirm which protocol a working tree or a built artifact actually carries.

## Inspecting a build

```bash
grep -n 'handshakeClientVer' internal/client/client.go   # announced clientVer
ls pkg/pb/**/*.pb.go | wc -l                              # protocol fingerprint
git describe --tags                                       # SDK release
```

## Adding a protocol upgrade

1. Replace `api/proto/` and run `./scripts/regen-all-protos.sh`.
2. Bump `handshakeClientVer` in `internal/client/client.go` to the new value.
3. Add a row to the mapping table above.
4. Note the upgrade under `[Unreleased]` in [CHANGELOG.md](../CHANGELOG.md),
   including the new proto count.
