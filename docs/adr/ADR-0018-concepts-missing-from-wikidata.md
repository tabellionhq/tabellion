# ADR-0018: Concepts missing from Wikidata

- Status: Accepted
- Date: 2026-10-19
- Drafted by: agent (Claude Code, claude-opus-5-5); options, decision and reason by owner

## Context

A Practice names its concept by a Wikidata ID (ADR-0001). Some concepts a person practised have no Wikidata item. Accepted facts SHALL NOT be mutated (ADR-0008), so a fact cannot be corrected later to point at an item created afterwards.

Wikidata keeps an item only if it describes a clearly identifiable concept backed by serious public sources: some fine-grained concepts may be refused or deleted. Recording a practice must not wait on Wikidata (APR-0002).

## Options

| Option | How | Pros | Cons |
|---|---|---|---|
| 1. A string replaces the ID | `Practice.concept` is a Wikidata ID or a free label | Simplest | The fact keeps the string forever, since it cannot be mutated; same words, different meanings across facts; no shared identity for search |
| 2. Local concepts | A local concept (own URI, label, description) in reference data; `same_as` set to the Wikidata ID once the item exists | Facts stay frozen; the concept is shared across facts | One more object to curate; a vocabulary of our own beside Wikidata |
| 3. Force creating the concept on Wikidata | The fact is recorded only once the person has created the Wikidata item | One vocabulary, contributed back to the commons | Blocks recording; needs a Wikidata account; Wikidata may refuse or delete the item |
| 4. Compose base concepts, with roles | A concept is a genus (`is`) and what sets it apart (`with`), both Wikidata IDs, plus an optional label | Only existing Wikidata IDs, nothing to link later; no curation | A composition is approximate; several compositions may describe one concept |

### Against the principles

| Option | APR-0001 Free and open | APR-0002 Simple adoption | APR-0003 Existing standards |
|---|---|---|---|
| 1. String | ✅ | ⚠️ simplest to start, weak search later | ❌ no shared identifier |
| 2. Local concepts | ✅ | ❌ one more object to understand and curate | ⚠️ own vocabulary beside Wikidata |
| 3. Force create | ✅ contributes to the commons | ❌ blocks recording, needs an account | ✅ |
| 4. Compose, with roles | ✅ | ✅ pick existing items, add a label if needed | ✅ Wikidata IDs only |

## Recommendation

4. A concept is always composed of earlier ones: a genus and what sets it apart.

```yaml
Practice:
  concept:
    is: uri               # genus, Wikidata, e.g. wd:Q… expression
    with: [uri]           # what sets it apart, Wikidata, e.g. wd:Q… block; empty when the genus is the concept
    label: string | null  # "block value", for display; may be fed to the embedding engine
  context: uri | null     # e.g. wd:Q575650 Rust
```

- When the concept exists on Wikidata: `is` is its ID, `with` is empty.
- When nothing precise fits: a broad genus and a label; the label carries the precision.
- When a composition is common enough to deserve its own item, the system offers a Wikidata draft in QuickStatements format, Wikidata's batch editing tool, for the person to review and submit. Recording never waits for it.

## Decision

4: concepts are composed of Wikidata concepts, with roles (`is`, `with`) and an optional label.

## Reason

Simple, only Wikidata IDs.

## Consequences

- Search is unchanged: ESCO skills by index (exact IDs, ancestors); concepts by semantic search over their Wikidata IDs' vectors and their label, which may be fed to the embedding engine. How a search adapter combines these vectors is the adapter's choice.

## To check

- Which Wikidata classes and properties fit programming concepts, for the QuickStatements draft.

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [APR-0003: Skills curriculum standard, existing standards for the rest](../apr/APR-0003-curriculum-standard-defer-the-rest.md)
- [ADR-0001: Describe concepts with Wikidata](ADR-0001-describe-concepts-with-wikidata.md)
- [ADR-0008: Normalisation of the data model](ADR-0008-normalisation.md)
- Wikidata notability: <https://www.wikidata.org/wiki/Wikidata:Notability>
- QuickStatements: <https://www.wikidata.org/wiki/Help:QuickStatements>
