# build-compliance Specification

## Purpose
The repository's build gates exist to catch regressions before they ship, but a gate that a push can bypass does not protect anything. This capability states what "green" must mean for `main` and for a release tag, so that a failing build is a blocked merge rather than a red badge nobody reads.

## Requirements

### Requirement: Main is gofmt-clean

Every Go source file tracked in the repository MUST be formatted with `gofmt`. Formatting MUST be checked, not assumed.

#### Scenario: Misaligned struct fields reach main

- **WHEN** a commit introduces Go code whose struct field alignment does not match `gofmt` output
- **THEN** the `build & test` workflow's Format check step fails, listing every non-gofmt-clean file
- **AND** the commit does not reach `main`, because the required status checks gate the merge

#### Scenario: Repository is verified before release

- **WHEN** a release is prepared
- **THEN** `gofmt -l .` produces no output
- **AND** the tree is gofmt-clean before the tag is created

### Requirement: Required status checks gate merges into main

`main` MUST be configured with the required status checks `build & test`, `docs guards`, `scan`, and `analyze`. Merges into `main` MUST be refused until all four pass. A push that skips the gate MUST be treated as a defect, not a shortcut.

#### Scenario: A change is delivered to main

- **WHEN** a change is delivered to `main`
- **THEN** it arrives through a pull request rather than a direct push
- **AND** the pull request is merged only after all four required status checks report success

#### Scenario: A direct push is attempted

- **WHEN** a commit is pushed directly to `main` instead of merged through a pull request
- **THEN** the four required status checks do not gate it, and `enforce_admins` being false permits the bypass
- **AND** this is treated as a gap to close, not as an accepted workflow

### Requirement: Release tags are cut from verified-green commits

A release tag MUST be created only from a commit whose local verification suite passed. The tag MUST NOT point at a commit that is known-failing.

#### Scenario: A release is cut

- **WHEN** a release version is tagged and pushed
- **THEN** the tagged commit passed all six local gates beforehand: `gofmt -l .` empty, `go build ./...`, `go vet ./...`, `go test -race -count=1 ./...`, `govulncheck ./...`, and `make docs-check`
- **AND** the `CHANGELOG.md` contains a section for that version, so `release.yml` can extract curated release notes

#### Scenario: A gate fails at release time

- **WHEN** any of the six gates fails
- **THEN** the release is held rather than tagged
- **AND** the failure is resolved in its own commit, so the defect is traceable to the change that introduced it

### Requirement: Release notes derive from the changelog

Release notes published by the tag-triggered workflow MUST be extracted from the `CHANGELOG.md` section for that version, and MUST NOT be generated ad hoc from the commit list.

#### Scenario: A tag triggers the release workflow

- **WHEN** a `v*` tag is pushed
- **THEN** the workflow locates the `## [x.y.z]` section in `CHANGELOG.md` for that version
- **AND** uses that section as the release body
- **AND** falls back to generated notes only when no such section exists
