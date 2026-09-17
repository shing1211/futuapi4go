# ADR 0009 — Releases are published by `gh` from the CHANGELOG

- **Status:** Accepted
- **Date:** 2026-09-17

## Context

The project needs tag-triggered GitHub releases with readable notes. goreleaser
was tried first but is a poor fit for a pure library (see
[ADR-0001](0001-pure-library.md)): it assumes binaries, and its auto-changelog
drops `ci:`/`docs:`/`chore:` commits — which dominate this repository's history.

## Decision

Pushing a `v*` tag triggers `.github/workflows/release.yml`, which:

1. Extracts the matching `## [x.y.z]` section from `CHANGELOG.md` (accepting
   both `[x.y.z]` and `[vx.y.z]` heading styles), and
2. runs `gh release create "$tag" --title "$tag" --notes-file notes.md --latest`,
   falling back to `--generate-notes` when no section exists.

`.goreleaser.yaml` was removed; `make release` remains a manual `gh` fallback.

## Consequences

- Release notes are the curated CHANGELOG entry, not a commit dump.
- The only requirement is the `gh` CLI, which is preinstalled on GitHub runners;
  maintainers do not need goreleaser installed.
- The CHANGELOG section is the single place to write user-facing release notes.
