# AgDR-0003: Answer key inside the quiz file

- Status: Warning (agent review): the key in the same file relies on honesty; acceptable only if the quiz is a self-assessment, not a check of the README promise "can explain code"
- Date: 2026-10-14
- Agent: Claude Code, Claude Opus 5.5 (claude-opus-5-5)

## Decision

Each quiz file ends with its answer key in a collapsed `<details>` section.

## Reason

One file per quiz is simpler to generate and review. The key is one click away, so it relies on the developer not opening it before answering.

## Review

Accept as an ADR, or reject and move keys to separate files.
