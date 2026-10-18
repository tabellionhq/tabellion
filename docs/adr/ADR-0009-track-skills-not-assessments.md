# ADR-0009: Track skills, not assessments

- Status: Accepted
- Date: 2026-10-18
- Drafted by: owner; options drafted by agent (Claude Code, claude-opus-5-5)

## Context

Assessments (questions, practical work, grading) are produced by requesters: Ismael at the end of a training, Claude after recent commits. An assessment can result in failures.

## Options

| Option | Description | Pros | Cons |
|---|---|---|---|
| A. Skills only | The record holds skills (Practices with levels); assessments and their failures stay with the requester | Record reads as progress, not surveillance (APR-0001); simple model; consistent with ADR-0004, ADR-0005 | "Never tried" and "tried and failed" look the same in the record |
| B. Skills and assessments | The record also holds assessments and their outcomes, failures included | Full history for the next assessment | Record of failures reads as monitoring; the model grows with assessment formats it does not own |
| C. Skills, plus private notes | The record holds skills; the person may keep private notes ("struggling with X"), never exposed | Keeps a hint for the next assessment | A second kind of data to protect; notes are not skills |

## Decision

Skills are positive; assessments can result in failures. Our system tracks skills, not assessments.

We shift left to the requester the tracking of why they asked and what they do with our data.

## Reason

KISS: do one thing and do it well.

## Consequences

- We design something simple, meshable, pluggable. The system that pushes new teachings at Anne's rhythm is out of the curriculum tracking system, although it has all the good reasons to use it. Goals (targets) belong to it, not to the curriculum.

- A failure records nothing in the curriculum. Level 1 (basic) is a skill, not a failure.
- A requester that needs assessment history (e.g. Claude after commits, UC1) keeps it on its side.

## References

- [APR-0001: Free and open](../apr/APR-0001-free-and-open.md)
- [ADR-0004: Trust and progress are out of scope](ADR-0004-trust-and-progress-out-of-scope.md)
- [ADR-0005: Scope is skill curriculum recording](ADR-0005-scope-assessment-method.md)
