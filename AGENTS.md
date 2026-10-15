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
- Decisions are tracked in ADR `docs/adr/ADR-000x-(subject).md`: the owner's decisions only, including previsions and eventualities.
- Implementation decisions are prepared by agents in AgDR `docs/agdr/AgDR-000x-(subject).md`, from `docs/agdr/TEMPLATE.md`.
  - What: a decision to be made because an implementation is required.
  - When: about to code something whose pattern or method has not been determined yet.
  - How:
    1. The agent writes the AgDR: context, options, issues, impacts, recommendation, next steps. Status Proposed.
    2. The agent spawns an honest reviewer sub-agent with a short, low-cost prompt; it answers OK, Warning or Refused with one line of reason, recorded in the AgDR.
    3. The agent helps the owner validate: summary, recommendation, reviewer warnings or refusals only.
    4. The owner accepts, picks another option, or rejects. Status Accepted or Rejected; Withdrawn if the need disappears.
  - Code that depends on an AgDR waits for step 4.
- README, comments, all user facing docs must not be poisoned with context and agent thoughts.
