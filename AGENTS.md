# Principles

Unless asked for, ALWAYS apply these rules

## Global

1. Answer, don't explain
2. Ask the user: avoid inventing needs or requirements
3. Say "I don't know" when in doubt: flag what is uncertain
4. For docs and discussions: Bottom line first, plain words, everyone's busy
5. Avoid hallucinations by grounding your assertion in truth and facts
6. Prove it: evidence over plausibility
7. Challenge and correct the user when a critical approach is due
8. Track agentic work in task files and commits: which model, which GenAI tool, which task
9. README, code comments, ADR, changelog: write as if no task existed: every sentence must stand without the task files

## Planning

1. Require a workflow as soon as possible: define steps and triggers
2. Require early User experience scenarios to ground usability
3. Scope the search: explore only with clear guidelines
4. Plan before code, 3 bullets: FILES to add/edit, STRATEGY chosen, RISKS e.g. side-effects, breaking changes, edge cases

## Safety

1. Protect data: prefer copies and backups by default
2. Ask before running anything that would delete, move or install
3. Evaluate blast radius before running commands: alert on danger

## Coding

1. Use Spec Driven Dev
2. Use TDD: Red Green
3. Use hexagonal
4. Use SOLID
5. Add no dependency the existing stack can cover
6. Handle errors explicitly on IO, API, user input
7. Output complete code: no placeholder, no "unchanged" marker

## DevOps

1. Use Changelog and semantic versioning
2. Use gitmoji, conventional commit
3. Use Continuous integration, early as possible
4. Every build is release-ready: short lived feature flags, branch by abstraction, expand-contract migrations
5. Require Continuous Deployment early: cloud-ready, IAC

## Command line tools

1. Return help, usage and examples for easy tool discovery
2. Test with E2E in docker
3. STDOUT and STDERR are mandatory: strive for valuable info on both channels
4. Prefer structured output: YAML where possible, for human and LLM use
5. System operations require dry-run options

# Project workflow

- Tasks are tracked in `.tasks/yyyymmdd-(subject).md`.
- Principles are tracked in APR `docs/apr/APR-000x-(subject).md`: the owner's words only.
- Decisions are tracked in ADR `docs/adr/ADR-000x-(subject).md`, from `docs/adr/TEMPLATE.md`. The owner writes the decision and the reason; an agent may draft context, options, issues, impacts and recommendation.
- Decisions agents make while implementing are tracked in AgDR `docs/agdr/AgDR-000x-(subject).md`, from `docs/agdr/TEMPLATE.md` (format: <https://github.com/me2resh/agent-decision-record>).
- README, comments, all user facing docs must not be poisoned with context and agent thoughts.

## Agent decisions

Direct by default: the agent decides, writes the AgDR, commits it with the code.

- Direct (AgDR): new tests, new adapters, new interfaces, editing a method, adding a step to CI, and similar changes that are easy to reverse.
- Escalate (ADR) when the decision has a strong impact on the future of the project:
  - adding or upgrading a dependency (including its licence),
  - defining a storage method or data format,
  - changing the way software is built or delivered (release, deploy, publish steps),
  - breaking changes (public API, config, data already written),
  - security (permissions, secrets, disabled checks), personal data, paid services.
- Tie-breaker: a change that touches any escalation item escalates, whatever its size or form (an adapter that brings a new vendor, a CI step that publishes).
- Doubt: escalate.
- Escalated decisions follow the ADR process: the agent drafts the ADR (status Proposed) with options, issues, impacts, recommendation, next steps; the owner decides. Code that depends on it waits; the agent works on something else meanwhile.

## Post-commit review

When chatting with the owner, every AgDR committed since the last review is presented to the owner:

1. A sub-agent dedicated to reviewing spawns with a short, low-cost prompt. It reads the AgDRs and their diffs and answers per AgDR: OK, Warning or Refused, with one line of reason, and flags any AgDR that should have been escalated.
2. The agent presents every AgDR to the owner: one line each, with the reviewer verdict; details only for warnings and refusals.
3. The owner keeps it, asks for a change, or pushes it up to an ADR. The AgDR records the owner's answer in its Review section.
