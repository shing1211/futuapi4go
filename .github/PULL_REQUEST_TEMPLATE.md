<!--
Copyright 2026 shing1211
SPDX-License-Identifier: Apache-2.0
-->
## Summary

<!-- What does this PR change, and why? -->

## Related issue / endpoint

<!-- Link the issue, and (for API changes) the Futu ProtoID / method. -->

## Type of change

- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation
- [ ] Refactor / internal

## Checklist

- [ ] `go build ./...` passes
- [ ] `gofmt -l .` is clean (`make fmt-fix`)
- [ ] `go vet ./...` passes
- [ ] `go test -race ./...` passes
- [ ] `make docs-check` passes (README translations in sync)
- [ ] New/changed public functions have GoDoc comments
- [ ] `CHANGELOG.md` updated under `[Unreleased]`
- [ ] `docs/VERSION_MAP.md` updated (for proto/protocol changes)
- [ ] No dangling doc links
- [ ] No secrets (keys, passwords, tokens) committed
- [ ] Proto changes regenerated via `./scripts/regen-all-protos.sh`
- [ ] Trading operations are **not** auto-retried
- [ ] Commits are DCO-signed (`git commit -s`)
