# ADR-0006: The model is agnostic of the human organisation

- Status: Accepted
- Date: 2026-10-15
- Drafted by: owner

## Decision

Think about it like git: we provide the model, and it is agnostic of the human organisation. A developer can work locally, and a megacorp can have a huge monorepo with 3-person reviews through a git forge.

Who proposes facts (the person, a training provider, an employer) and who accepts them is out of scope.

## Notes

The git analogy is about organisation only: storage is probably a database, and the model has no commits.

## References

- [ADR-0004: Trust and progress are out of scope](ADR-0004-trust-and-progress-out-of-scope.md)
- [ADR-0005: Scope is skill curriculum recording](ADR-0005-scope-assessment-method.md)
