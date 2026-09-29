# Design

## Context

Two changes landed out of order. `build-compliance` was written and archived during the v0.21.0 release work, and `enforce_admins: true` was enabled afterwards as a follow-up. Both artifacts describing the enforcement were therefore authored against `enforce_admins: false`.

The consequence is not symmetric. The stale command in `AGENTS.md` is an operational trap: the next operator who follows the documented release process gets `GH006` and has to diagnose the protection rule mid-release. The stale spec scenario is a correctness problem: it records a bypass as a tolerated condition, which invites a later change to treat enabling enforcement for admins as still-pending work and relax it.

The spec is the artifact with a memory, so it takes priority.

## Goals / Non-Goals

**Goals:**

- Make the documented release process executable under enforced branch protection.
- Correct the spec so it describes the enforcement that now exists, and state it strongly enough that it cannot be quietly reversed.
- Keep the change documentation-only, with no code, proto, or public API impact.

**Non-Goals:**

- Not adding `required_pull_request_reviews`. With one maintainer, a self-approval requirement is ceremony; the status-check gate is the control that matters.
- Not touching the Gitee mirror's protection. Gitee offers branch protection in its UI, but automating it is a separate investigation and the mirror is a secondary target, not the merge path.
- Not rewriting `docs/PHASE4_API_COVERAGE_PLAN.md` or `docs/PHASE5_BUGFIX_HARDENING_PLAN.md`, which also contain `git push origin main`. Both are COMPLETE point-in-time records.
- Not adding an `enforce_admins` CI check. The setting lives in GitHub's API, not the repository, so a workflow could only detect drift by reading the API and would need credentials; the spec records the requirement and drift would be caught by the next release attempt.

## Decisions

**Route the spec change through a delta, never a direct edit.** `openspec/specs/build-compliance/spec.md` is a main spec. Editing it in place would bypass the review path and leave the change folder inconsistent with the main spec, so the correction travels as a `MODIFIED` delta and lands via `openspec archive`. The requirement heading was matched byte-for-byte against the main spec, because a near-miss title would silently create a second requirement instead of modifying the first.

**Add the `enforce_admins` clause to the requirement text, not only to a scenario.** The scenario records the current behavior; the requirement is what future changes are checked against. Stating it as a MUST in the requirement means a later proposal that disables enforcement is visibly contradicting a documented requirement, rather than appearing to be a fresh consideration. The new *Enforcement is weakened* scenario makes the failure mode legible: it names the setting, the consequence, and the fact that it is a regression.

**Rewrite the release process rather than annotate it.** The existing block is ordered wrong under protection: it pushes `origin main` before tagging, and the gitee mirror is synced in the same breath as the tag. With protection, `main` advances only through a merged PR, so the mirror sync has to move after the merge. Patching individual lines would leave the ordering subtly wrong; a rewrite makes the dependency explicit — you cannot sync the mirror until the merge exists.

**Document the Gitee mirror step explicitly as unprotected.** The mirror accepts direct pushes, so the two remotes can drift: a hotfix pushed to Gitee alone leaves GitHub behind with no error anywhere. The rewritten process syncs `gitee main` after the merge and names the asymmetry, so the mirror is a deliberate follow-up step rather than an incidental one.

**Add exactly one line to the Session Checklist, and leave the Documentation checklist alone.** The Session Checklist is the end-of-task gate an operator actually runs, so it is where "delivered via PR" belongs. The Documentation checklist at line 125 duplicates "committed with descriptive messages"; adding the same idea twice in one file invites them to drift apart, and the delta is not a documentation artifact.

**Correct the footer date as part of the same edit.** `Last updated: 2026-09-17` becomes the date of this change. Leaving it would make the file's own metadata wrong the moment it is edited, which is the same class of staleness this change exists to fix.

## Risks / Trade-offs

**A spec that encodes current state can become stale again.** This change exists because a spec written two days ago described a setting that was subsequently flipped. The mitigation is the third scenario: naming `enforce_admins` explicitly means the next person to change it sees the requirement they are about to violate. It is not a complete defense, but it converts an invisible drift into a documented conflict.

**The required status checks are named in the spec as literals.** `build & test`, `docs guards`, `scan`, and `analyze` are workflow job names. If a workflow is renamed, the spec becomes inaccurate. Accepted: the names are the contract, and a rename that broke the gate would surface immediately as a blocked PR.

**No review requirement means the PR gate is the only gate.** With `required_pull_request_reviews: false`, a single maintainer can merge their own PR as soon as checks pass. The protection stops broken code, not wrong judgment. Accepted for a single-maintainer repository, where a second reviewer is not available.

**The docs guards job does not validate AGENTS.md.** `docs guards` runs only `scripts/check_i18n.py`, so nothing in CI will catch a future edit that reintroduces a direct-push instruction. The release process documentation therefore depends on review rather than on enforcement.
