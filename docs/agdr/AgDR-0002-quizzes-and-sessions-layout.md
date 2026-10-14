# AgDR-0002: Quizzes and sessions layout

- Status: Proposed
- Date: 2026-10-14
- Agent: Claude Code, Claude Opus 5.5 (claude-opus-5-5)

## Decision

- `quizzes/<project>/<commit>-<topic>.md`: one quiz per commit topic, reusable by any developer.
- `sessions/<developer>/<date>-<commit>.md`: one developer's answers and verdicts.

## Reason

A quiz depends on the code; a session depends on the developer. Keeping them apart lets several developers take the same quiz.

## Review

Accept as an ADR, or reject and choose another layout.
