---
id: AgDR-0004
timestamp: 2026-10-28T20:17:14Z
agent: Claude Code
model: claude-opus-5-5
trigger: .tasks/20261026-tabellion-design-backlog.md, "DCO and licence checks in CI"; owner request
status: executed
---

# AgDR-0004: DCO check in CI

> In the context of ADR-0025 requiring a sign-off on every commit, facing a repository with no CI, I decided to add a GitHub Actions workflow that checks each new commit with the GitHub API and a shell script, to refuse commits without the author's sign-off, accepting that it runs only on GitHub and does not cover commits made before it.

## Options considered

- A: Shell script with `gh api` and `jq`, both on the GitHub runner; no checkout, no third-party action.
- B: A marketplace DCO action: less code, but a third-party dependency running with the repository token.
- C: The DCO GitHub App: no code, but a third-party service with access to the repository.

## Decision

A. The workflow `.github/workflows/dco.yml` runs on pull requests and on pushes to `main`. Each commit needs a `Signed-off-by` line carrying the author's email, compared without case.

## Consequences

- No new dependency; read-only token permissions.
- Commits made before the check (55 of the first 97) are not checked; history is not rewritten.
- Tested in Docker with a stubbed `gh`: a matching sign-off passes; a missing one, or one with another email, fails. Not yet run on GitHub.

## Review

- Reviewer sub-agent: OK: enforces ADR-0025 sign-off; no new dependency, read-only token, no publish step; not run on GitHub yet
- Owner: kept (2026-10-28)
