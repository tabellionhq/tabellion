# ADR-0004: Trust and progress are out of scope

- Status: Accepted
- Date: 2026-10-15
- Drafted by: owner

## Context

How evidence becomes a skill level (1 to 5, ADR-0002).

## Decision

Deciding trustworthiness and progress is out of scope / domain. The system is pluggable, has its own bounded context logic, uses assessment as a base, but does not decide trustworthiness and progress.

We may provide simple solutions:

1. Testable interfaces.
2. A smooth "product adoption" via a 100% local and user-based trust system.

That's probably all.

## Reason

It's a human problem with multiple dimensions: truth vs. cheating, individual vs. collective, local vs. global.

## References

- [APR-0001: Free and open](../apr/APR-0001-free-and-open.md)
- [ADR-0002: Skill level scale](ADR-0002-skill-level-scale.md)
