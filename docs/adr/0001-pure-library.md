# ADR 0001 — Ship a pure library, not binaries

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

futuapi4go is an SDK consumed by other Go programs. It has no command-line
surface of its own, and Futu's OpenD is the process that actually connects to
Futu's servers.

## Decision

The module is a **pure Go library**. There is no `main` package, no `cmd/`
directory, and no released binaries. `go build ./...` compiles only library
packages.

## Consequences

- Releases publish source and a GitHub release, not build artifacts.
- The release workflow does not build binaries (see
  [ADR-0009](0009-release-from-changelog.md)); goreleaser, which assumes
  binaries, is not used.
- Runnable examples live in the companion
  [`futuapi4go-demo`](https://github.com/shing1211/futuapi4go-demo) repository
  rather than in this module.
