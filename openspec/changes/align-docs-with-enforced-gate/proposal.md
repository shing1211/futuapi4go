# Proposal

## Why

Branch protection on `main` now sets `enforce_admins: true`, so the four required status checks apply to administrators too. That change was made *after* the release process was documented and after the `build-compliance` spec was written, which left two artifacts describing a workflow GitHub now rejects.

`AGENTS.md:155` instructs the release step as `git push origin main && git push gitee main`. That command now fails with `GH006: Protected branch update failed`, so the project's primary agent-facing guide tells the next operator to do something forbidden. The `build-compliance` spec has the mirror-image problem: it records `enforce_admins` as `false` and describes a direct push as a permitted bypass that should merely be "treated as a gap to close", when the push is in fact rejected.

The spec is the durable record, so a stale scenario is worse than a stale command: it would let a future change conclude that enforcing the gate for admins is still outstanding, and relax it.

## What Changes

- **`AGENTS.md` Release Process** — rewritten from commit-on-`main` to branch → pull request → required checks → merge → tag → mirror. The documented order is also corrected: the previous text pushed `origin main` before tagging, whereas `main` can now only advance through a merged PR, so the Gitee mirror must be synced after the merge rather than alongside the tag push.
- **`AGENTS.md` Session Checklist** — one line added, so the checklist reflects that delivery is via PR with all four checks green.
- **`AGENTS.md` footer date** — `Last updated: 2026-09-17` corrected to the date of this edit.
- **`build-compliance` spec** — the *Required status checks gate merges into main* requirement gains a clause making `enforce_admins` a MUST, so the enforcement cannot be silently turned off. Its *A direct push is attempted* scenario is corrected from "does not gate it, and `enforce_admins` being false permits the bypass" to the push being rejected with `GH006`.

No breaking changes. No code, no protos, no public API.

## Capabilities

### Modified Capabilities

- `build-compliance`: The *Required status checks gate merges into main* requirement now states that the checks MUST be enforced for administrators as well, and its direct-push scenario asserts rejection with `GH006` rather than a permitted bypass.

### New Capabilities

None.

## Impact

**Affected files**

- `AGENTS.md` — Release Process block (lines 147-162), one checklist line (229), footer date (234)
- `openspec/specs/build-compliance/spec.md` — updated via a `MODIFIED` delta through this change, then synced on archive; never hand-edited

**Public API compatibility: none.** Documentation only; no identifier, type, or exported name changes in `client/` or `pkg/`.

**Protos affected: none.** No change touches `api/proto/` or `pkg/pb/`.

**Verification** — `gofmt -l .` empty, `make docs-check` (the `docs guards` CI job, which runs only `scripts/check_i18n.py`; `AGENTS.md` is not a translated README so it cannot affect it), and `openspec validate --specs --strict` after the delta syncs.

**Deliberately unchanged**

- `docs/PHASE4_API_COVERAGE_PLAN.md:902` and `docs/PHASE5_BUGFIX_HARDENING_PLAN.md:725` also contain `git push origin main`. Both are COMPLETE point-in-time records, the same category as the historical version markers left untouched during the v0.21.0 sweep. Rewriting them would alter the record of how those phases were actually delivered.
- The `build-compliance` requirement *Main is gofmt-clean* is untouched. Its first scenario asserts that an unformatted commit "does not reach `main`, because the required status checks gate the merge" — that statement was untrue when written and became true when enforcement was enabled, so it is now accurate and needs no edit.
