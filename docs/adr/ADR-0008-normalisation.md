# ADR-0008: Normalisation of the data model

- Status: Accepted
- Date: 2026-10-18
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

Searching facts combines an ESCO category (hierarchy), a view (dimensions) and a meaning (vectors on Wikidata concepts). Copying ESCO ancestors and concept vectors onto each fact simplifies queries but duplicates reference data: fact → skill → ancestors is a transitive dependency, which third normal form (3NF) forbids.

## Options

| Option | Description | Pros | Cons |
|---|---|---|---|
| A. Normalised model, denormalised adapters | The logical model is in 3NF: facts reference identifiers; ESCO hierarchy, labels and vectors live once as reference data. An adapter may denormalise for performance, documented. | One source of truth; model independent of storage and search engines; adapters free to optimise | Two shapes to keep in sync (logical, physical); each adapter's denormalisation must be documented and tested |
| B. Normalised everywhere | 3NF in the model and in every adapter; queries join reference data | Simplest to reason about; no duplication anywhere | Some engines (vector stores, search engines) expect data on the row and filter poorly across joins; may block useful adapters |
| C. Star schema as the model | Facts with denormalised dimensions are the canonical model | Fast, simple analytical queries (progress, aggregations); well-known pattern | Model tied to an analytical use; reference data duplicated in the core; changes to ESCO or vectors ripple through all facts |
| D. Document per person | One nested document per person (record, views) | Matches "the person owns and exports everything" (APR-0001); easy to export | Cross-person queries (UC9, UC11) are awkward; reference data duplicated or joined by hand; weak constraints |
| E. Event log with projections | Facts are an append-only log of events (practice recorded, level set, view exposed); searchable shapes (3NF, star, vector index) are projections rebuilt from it, never the log itself | A person's progress flows naturally; history for free; any projection can be rebuilt | More moving parts; projections lag (eventual consistency) and are rebuilt when reference data changes; deletion clashes with an immutable log (crypto-shredding: one key per person); events keep their old shapes forever (versioning); corrections are new events |

## Recommendation

A, with facts append-only: a dated Practice is already an event, so a person's progress is the list of their facts.

## Decision

A: append-only facts are the source of truth, in a normalised model; denormalisation only inside adapters.

The skill curriculum may provide events, as an event stream for a larger system at company level.

## Reason

Best data model for vast and manageable volumes.

A has the history benefits of E without the added complexity. E is better for large distributed systems, which we are trying to avoid.

Keeping a log matters: I can be a noob in Rust in January, intermediate in June, and average in September after a lot of summer practice.

## References

- [Data model](../design/data-model.md)
- [Data first](../design/data-first.md)
