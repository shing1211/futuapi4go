# Design

## Context

Three facts shape this change.

**The gate is configured but unenforceable.** `main` carries four required status checks (`build & test`, `docs guards`, `scan`, `analyze`) with `enforce_admins: false` and no required pull-request reviews. A push by an admin therefore skips all four. GitHub reports the bypass on every such push. The v0.20.0 release commit `d435f9b` was pushed directly, so the failing Format check that followed could not block anything, and the next ten-plus commits inherited the red state.

**The defect is whitespace only.** `gofmt -l .` names three files, all touched by `d435f9b`. The misalignment is in the `// WebSocket` block of `ClientOptions`, where `WSSecretKey` and `WSReconnect` carry an extra space before their types. Because a comment line separates them from the `PushHandler` field below, `gofmt` treats them as their own alignment group and expects single spaces. There is no identifier, type, or behavior change.

**The dependency surface is smaller than it first appears.** `go list -m -u all` reports eleven outdated modules, but that command walks the *extended* module graph, including test-only dependencies of dependencies. Only four of the eleven appear in `go.mod`, and all four are indirect. The six direct requires are already current. This matters for scoping: a global replace of `0.20.0`-style version strings is unnecessary, and `go get -u ./...` is a small, bounded edit rather than a large transitive cascade.

## Goals / Non-Goals

**Goals:**

- Make `gofmt -l .` empty and keep the Format check green.
- Move the four indirect modules to their latest versions with all six gates verified locally.
- Cut `v0.21.0` from a commit that is known-green, with curated release notes.
- Deliver the change through a pull request, so the four required status checks gate the merge for the first time.

**Non-Goals:**

- Not touching the seven "outdated" modules absent from `go.mod`. They are test-only dependencies of other projects' test suites and are not reachable by `go get` in this module.
- Not changing `enforce_admins` on `main`. Turning the gate on for admins is the durable fix for recurrence, but it is a repository-administration change that also blocks emergency direct pushes, so it is raised as a follow-up rather than done silently here.
- Not retroactively fixing the v0.20.0 tag. A published release's build result is immutable; the repair lands in v0.21.0 and the CHANGELOG names the cause.
- Not sweeping all 20 version references inside this PR. The remaining 18 files are updated post-release, following the precedent of `e8b91e7`, `55a4c35`, and `746120e`.
- Not regenerating protos. No proto changes, so `pkg/pb/` is untouched.

## Decisions

**Fix formatting with `gofmt -w`, then inspect the diff.** Rather than hand-editing the alignment, run `gofmt -w` on the three named files and review the resulting diff. Rationale: `gofmt` is the same tool CI runs, so the result is correct by construction. A hand edit risks leaving a second misalignment that only surfaces on the next Go release. Alternative considered — manually collapsing the double spaces — rejected as the same edit with more ways to be wrong.

**Use `go get -u ./...` rather than pinning each module.** Rationale: it resolves the four indirect bumps that a real dependency refresh would produce, without inventing version choices. Verified in a scratch copy of the tree: exactly four modules move and `go.sum` grows 47→55 lines, purely additive. Alternative considered — `go get -u=patch`, which would have left the three minor-version Prometheus bumps in place — rejected because the user asked for a full update and the minor bumps carry no proto or trading path.

**Order the commits format-first, deps-second.** The format fix is independently reviewable — three files, whitespace only — and it is the change that makes CI green. Putting it first means a reviewer can validate the substantive claim (the gate passes) before evaluating the dependency bump. It also means the dependency commit never carries the burden of explaining why the build was red.

**Deliver through a PR, not a direct push.** This is the decision that changes the project's failure mode. Direct push keeps the status checks advisory, which is precisely how a red `main` survived ten days. A PR makes them blocking. The cost is real — a branch and a merge click — and it is a workflow change for a repository whose 14 existing PRs are all dependabot. Accepted because the alternative is a gate that cannot act.

**Cut the tag only after the merge, and only from a locally verified commit.** `release.yml` extracts its notes from the `CHANGELOG.md` section matching the tag, so the `## [0.21.0]` heading must exist before tagging or the workflow silently falls back to generated commit notes. The unprefixed form `## [0.21.0]` is used to match the 39 unprefixed headings in the file against 18 prefixed ones; the workflow's `awk` matches `^## \[v?` and accepts both.

**Defer the doc sweep to a post-release commit.** Rationale: matches three prior releases, and keeps this PR's diff limited to the fix. The cost is a brief window where the tree is at v0.21.0 while 18 files still read v0.20.0 — the same state v0.20.0 was tagged in, and corrected within days.

## Risks / Trade-offs

**`enforce_admins: false` remains, so the gate is still bypassable.** This change makes the *first* delivery respect the checks, but nothing prevents the next direct push from skipping them again. The recurrence risk is therefore reduced, not eliminated. Mitigation is a follow-up: set `enforce_admins: true` once the emergency-push case is understood.

**The v0.20.0 tag stays red permanently.** Consumers looking at that release see a failed build. Accepted: the alternative — a `v0.20.1` patch that only reformats — publicly documents the regression as its own release, which is more signal than this project needs for a whitespace fix. The v0.21.0 CHANGELOG names `d435f9b` so the cause is traceable.

**Prometheus minor bumps could alter metric output.** `common` v0.70.1→v0.72.0 and `procfs` v0.21.1→v0.22.0 are minor versions. `go test -race -count=1 ./...` passes 29/29 packages and `govulncheck` reports no vulnerabilities, but the test suite does not assert on emitted metric text. The `/metrics` endpoint should be eyeballed once after deploy.

**`golang.org/x/sys` is shared across the module graph.** v0.47.0→v0.48.0 is a patch bump and is the one indirect module not confined to Prometheus. It builds clean on linux/amd64, the CI platform; other platforms are exercised by contributors, not by this gate.

**Two commits, one release, means the tag covers both.** A bisect landing between the commits sees a green tree either way, so the split is safe. The trade-off is that reverting the dependency bump later means reverting part of a released tag, which is no different from any other release.
