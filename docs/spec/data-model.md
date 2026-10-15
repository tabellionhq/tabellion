# Data model

- Status: Draft
- Decisions: ADR-0001 (Wikidata concepts), ADR-0002 (levels 1–5), ADR-0003 (ESCO), ADR-0004 to ADR-0006 (scope)

A person's track record is a set of facts. How facts are proposed, reviewed, accepted and stored is outside the model (ADR-0006).

## Objects

```yaml
Person:
  id: string              # opaque; no personal data required

Practice:                 # I practiced <skill> on <concept> in <context> at <occupation> on <date>
  person: Person.id
  skill: uri              # ESCO skill, e.g. esco:skill/…
  concept: uri            # Wikidata item, e.g. wd:Q1071360 publish–subscribe pattern
  context: uri | null     # Wikidata item, e.g. wd:Q575650 Rust
  occupation:
    id: uri | null        # ESCO occupation
    kind: work | pleasure
  date: date
  level: 1..5 | null      # ADR-0002; preset, set by the person, set externally, or empty
  evidence: object | null # any JSON, e.g. a signed token with claims, issuer, signature
  author: string          # who states the fact: the person, a training provider, an agent…
  recorded: datetime

Target:                   # I want <level> in <skill> on <concept> in <context>
  person: Person.id
  skill: uri
  concept: uri
  context: uri | null
  level: 1..5
  by: date | null
  author: string
  recorded: datetime
```

## Rules

- `level` is always an integer from 1 to 5, or null.
- `evidence` is opaque to the model: it is stored and returned, never interpreted (ADR-0004).
- Who set a level is the fact's `author`.

## Open

- Action (ADR-0001) vs ESCO skill: ESCO skills are already phrased as actions ("utilise machine learning"); is a separate action list still needed?
