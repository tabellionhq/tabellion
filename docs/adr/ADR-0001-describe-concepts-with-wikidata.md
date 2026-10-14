# ADR-0001: Describe concepts with Wikidata

- Status: Accepted
- Date: 2026-10-14

## Decision

Concepts are identified by their Wikidata ID (e.g. `wd:Q1071360` publish–subscribe pattern).

A skill combines an action, a concept and a context:
(designing) a (pubsub) in (rust) → design · `wd:Q1071360` · `wd:Q575650`

We don't need to map everything in every language: basic language features (e.g. Rust's block value) are not concepts of their own.

## Alternatives

1. Wikidata.
2. Build our own taxonomy.
3. Use a more generic one (e.g. ESCO).

Wikidata is best: open data (CC0), reputable, broad, maintained.

## Open

- The list of actions.
- Concepts missing from Wikidata.

## References

- Wikidata: <https://www.wikidata.org>
- Benchmark: [skills standards](../research/skills-standards-benchmark.md)
