# ADR-0026: Git workflow

- Status: Proposed
- Date: 2026-10-22
- Drafted by: agent (Claude Code, claude-opus-5-5); decision stated by owner

## Context

A push to the main branch sets a new semantic version and publishes a release; the full build for every target runs then (ADR-0020). Every build must be release-ready. Contributions are signed off under the DCO (ADR-0025).

## Options

| Option | How | Issues |
|---|---|---|
| A. Git flow | Long-lived `develop` and `main`, feature, release and hotfix branches | Many branches and merges; releases batched; long-lived branches drift |
| B. Short-lived feature branches | One branch per change, from `main`, merged back within days through a pull request with CI passing | Unfinished work must hide behind feature flags or branch by abstraction |
| C. Trunk only | Everyone commits to `main` | No review before a release, since each push to `main` releases |

## Recommendation

B.

## Decision

B: short-lived feature branches. Semantic versioning and conventional commits are enforced.

## Reason

Owner's words.

## Consequences

- The main branch is protected: changes arrive by pull request only, with CI passing and every commit signed off.
- Commit messages follow Conventional Commits, with a gitmoji after the type: `type(scope): :emoji: subject`. CI refuses other messages.
- The version is computed from the commit types merged since the last release: `fix` raises the patch, `feat` the minor, a breaking change (`!` or `BREAKING CHANGE:`) the major. The tool is chosen with the CI.
- The specification is versioned apart, as an Internet-Draft (ADR-0024).
- The current default branch is named `master`; ADR-0020 names it `main`. To rename before the first release.

## References

- [ADR-0020: Target platforms](ADR-0020-target-platforms.md)
- [ADR-0024: The specification is the source of truth](ADR-0024-specification-source-of-truth.md)
- [ADR-0025: Licences](ADR-0025-licences.md)
- Conventional Commits 1.0.0: <https://www.conventionalcommits.org/en/v1.0.0/>
- Semantic Versioning 2.0.0: <https://semver.org/spec/v2.0.0.html>
- gitmoji: <https://gitmoji.dev/>
