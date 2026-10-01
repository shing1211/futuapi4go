# ADR 0012 — CI-only dependency: secret scanning in the pipeline

- **Status:** Accepted
- **Date:** 2026-10-01

## Context

The repository had no credential scanner. `govulncheck` (its own workflow)
queries the vulnerability database for known CVEs and `gosec` analyses source for
insecure patterns, but neither looks for a token or private key that a human
pasted into a commit. A leaked credential can therefore enter the repository
through a pull request and reach every clone, and the SDK is public.

The gap is narrow but concrete, and this repository already commits material a
naive scanner flags as secret:

- `pkg/util/crypto_test.go` exercises the MD5 helper with fixed input/output
  pairs. `password123` paired with
  `482c811da5d5b4bc6d497ffa98491e38` reads as `generic-api-key` because the two
  strings sit together in a composite literal; the second is a published digest
  and the first is only ever hashed.
- `pkg/pb/` is generated from the Futu OpenAPI `.proto` definitions. It contains
  no credentials, but it is regenerated rather than hand-edited, so a finding
  there cannot be fixed by editing the file.

Scanning without an allowlist for these would fail every run and train reviewers
to ignore the job, which is worse than not having one.

## Decision

Add **`gitleaks` as a CI-only dependency, pinned by version, with an explicit
allowlist committed to the repository.**

- It runs as a pinned GitHub Action in a dedicated `secrets` job and scans the
  full history of the pushed ref, not only the diff, so a credential that landed
  in an earlier commit is still reported.
- The action is pinned to an exact release tag (`v2.3.9`). `@latest` and the
  floating `@v2` major tag are rejected for the same reason `govulncheck` is
  pinned: an upstream release must not change a security result without a commit.
- Configuration is committed as `.gitleaks.toml` and **allowlists by path**,
  never by disabling a rule class. Each path carries a comment naming why its
  content is not a credential, so a new entry is a visible decision rather than
  a silent suppression.
- The job receives `GITHUB_TOKEN`. gitleaks-action v2 requires it on
  `pull_request` events because it calls the GitHub API to report a finding;
  without it the job aborts in a few seconds having scanned nothing.
- It has **no effect on the Go module**: nothing in `go.mod`, and no Go package
  imports it, so consumers see no new dependency.

## Consequences

- A credential committed to any branch is reported by CI, including in history.
- False positives on the cases above are suppressed by path, and the allowlist is
  reviewable in a normal pull request.
- The job adds a third-party Action to the CI trust boundary. It runs under the
  repository's default `contents: read` permission and does not write to the
  repository, so its blast radius is a scan result, not a code change.
- Detection is a backstop, not a control; it does not replace keeping secrets out
  of the repository in the first place.