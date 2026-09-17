# Contributing to futuapi4go

Thank you for your interest in contributing!

- [Code of Conduct](./CODE_OF_CONDUCT.md)
- [Security Policy](./SECURITY.md)
- [Governance](./GOVERNANCE.md)
- [Support](./SUPPORT.md)
- [Developer guide](./AGENTS.md)

---

## Getting Started

### Prerequisites

- Go 1.26+
- A running [Futu OpenD](https://www.futunn.com/en/overview) instance (for
  integration tests; unit tests use an in-process mock)
- `protoc` / the repo scripts, only if you change `.proto` files

### Setup

```bash
git clone https://github.com/shing1211/futuapi4go
cd futuapi4go

go build ./...        # compile
go vet ./...          # lint
go test -race ./...   # full test suite with the race detector
```

---

## Repository Layout

```
.
├── client/          # Public wrapper API — recommended entry point
├── pkg/             # Business logic (qot, trd, sys, push, ...)
├── internal/        # Connection, TCP I/O, packet framing
├── api/proto/       # Futu protocol definitions
├── scripts/         # Proto regeneration + checks
├── test/            # Integration tests (build tag: integration)
└── docs/            # Architecture, usage, version map
```

---

## How to Contribute

### 1. Fork and Branch

```bash
git checkout -b feat/your-feature-name
# or
git checkout -b fix/your-bug-fix-name
```

Branch naming: `feat/`, `fix/`, `docs/`, `test/`, `chore/` prefixes.

### 2. Make Changes

- **New API**: confirm the proto in `api/proto/`, run
  `./scripts/regen-all-protos.sh`, then add the wrapper in `pkg/qot/` or
  `pkg/trd/` (context first, input validation at entry, `wrapError()`, nil
  guards on list iteration). Add a public helper in `client/client.go` if it
  simplifies usage.
- **Bug fix**: add a failing test first, then fix.
- **Docs**: update the relevant `.md` files and GoDoc comments.

### 3. Run Checks

```bash
make check          # gofmt + go vet + build

# Or the full gate:
go build ./... && go vet ./... && go test -race ./...
```

### 4. Commit

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
feat(qot): add GetOptionChain wrapper
fix(trd): handle nil order response
docs(readme): add push example
test(sys): cover reconnect path
```

DCO sign-off (`git commit -s`) is welcome but not required.

### 5. Pull Request

- Fill out the pull-request template.
- Reference the related issue (if any).
- For API changes: include the ProtoID / method name.
- Ensure CI is green.

---

## Code Style

- **Formatting**: `gofmt` (see `make fmt`).
- **Context**: all public functions take `context.Context` as the first arg.
- **Proto safety**: never call generated `GetXxx()` accessors on proto messages;
  use direct nil checks (`if msg.Field != nil`).
- **Error handling**: never swallow errors with `_`; use `wrapError()` for proto
  errors. `FutuError` supports `Unwrap()`.
- **Orders**: **never** auto-retry `PlaceOrder`, `ModifyOrder`, or `CancelOrder`.
- **Secrets**: wrap sensitive fields (e.g. `PwdMD5`) in `SensitiveString`.
- **Concurrency**: protect shared state with a mutex; every goroutine needs an
  exit path.

See [AGENTS.md](./AGENTS.md) for the full review checklist.

---

## Testing Guidelines

| Test Type | Location | Needs OpenD |
|-----------|----------|-------------|
| Unit tests | `*_test.go` alongside source | No (mock server) |
| Integration tests | `test/integration/` (`-tags=integration`) | Yes |

Always run with `-race`. See the [README](./README.md) "Build & Test" section.

---

## Reporting Issues

Bug reports welcome. Please include:

- Go version (`go version`)
- futuapi4go version (git commit or tag)
- Futu OpenD / protocol version (see [docs/VERSION_MAP.md](./docs/VERSION_MAP.md))
- Simulate or real environment
- Minimal reproduction case
- Full error output, with secrets redacted

---

## Translations

Translations of the [README](./README.md) are welcome and follow
[TRANSLATING.md](./TRANSLATING.md):

- English is canonical; translations are best-effort.
- Add/update the language switcher in **every** `README*.md`.
- Keep the translation banner and its `Last synced:` commit current.
- Do **not** translate legal text (`LICENSE`, `DISCLAIMER.md`).
- Run `make docs-check`.

Current languages: English, 简体中文, 繁體中文, 日本語, 한국어, Español.

---

## License

By contributing, you agree that your contributions are licensed under the
[Apache License 2.0](./LICENSE).
