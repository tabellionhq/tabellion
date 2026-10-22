# Data model

- Status: Draft
- Specification: [docs/RFC.md](../RFC.md), Sections 3 and 4: the objects and their rules. It prevails over this note.
- Decisions: ADR-0001 (Wikidata concepts), ADR-0002 (levels 1–5), ADR-0003 (ESCO), ADR-0004 to ADR-0007 (scope, views); ADR-0008 (append-only facts), ADR-0009 (skills, not assessments), ADR-0010 (person identifier), ADR-0018 (composed concepts)
- Stories: US-01 to US-10

This note holds what the specification leaves to an implementation: reference data, derived data, search, storage choices, and why the model is shaped this way.

## In the specification

| Object or rule | Specification |
|---|---|
| Person, Identity | Sections 3.1, 3.2 |
| Facts: occupation held, practice, concept | Sections 3.3 to 3.6 |
| View, Grant, access log | Sections 3.7, 3.8 |
| Task | Section 3.9 |
| Status, frozen facts, deleting, completing a proposed practice | Section 4.1 |
| Identifiers, named values | Section 4.2 |
| Levels 1 to 5 | Section 4.3 |
| Composed concepts, human languages, skills without a period | Section 4.4 |
| View semantics: deny list wins, period overlap, status defaults, filter before search | Section 4.5 |
| Access through views and grants; the owner needs no grant | Section 4.6 |
| Facts proposed by a task | Section 4.7 |

## Implementation

### Grant

A grant also records how the reader authenticates (ADR-0013):

```yaml
Grant:
  auth: socket | bearer | oidc
```

### Evidence of a CV import

```yaml
evidence: {excerpt: string, file: string, line: integer}
```

### Task result

```yaml
cv-import: {unresolved: [{term: string, excerpt: string}]}
```

### Reference data

Shared by all persons, computed from ESCO and Wikidata, never personal.

```yaml
SkillAncestor:            # ESCO hierarchy, closure: every group above a skill
  skill: uri
  ancestor: uri

Label:                    # labels and descriptions, all languages
  id: uri
  language: string
  label: string
  description: string | null

ConceptVector:            # for semantic search on Wikidata concepts; tied to the model that computed it
  id: uri
  model: string
  vector: [float]
```

### Derived, not stored

- **Competency**: a person's latest accepted level per skill, concept and context, computed from their practices.
- Gaps, readiness, progress, matches: reads over facts.

### Rules of this implementation

- Search: ESCO skills by index (exact IDs, ancestors); concepts by semantic search over the vectors of their Wikidata IDs and of their label. Label vectors belong to the search adapter (ADR-0008).
- How the store encodes status, kind, state and fidelity (e.g. small integers) is the store adapter's choice; the UI translates their names for display (ADR-0019).
- Responses embed the label of each referenced skill, occupation and concept, in the requested language; the store keeps identifiers only.
- Reference data: ESCO is shipped with the release; Wikidata items are looked up online and cached.
- View rules may be written by an agent from the owner's words; the owner confirms them.
- No action list: the action is the ESCO skill.

## Uses

- CV: ESCO skill + level is simple enough. ESCO is very coarse, level is approximate.
- Day-to-day teaching: Wikidata concepts are very useful, though overly complex for some contexts.
- Analytics, out of scope as functions: the model only keeps the dimensions they need. Examples: where did I learn the most (occupation), in which month (period), do I keep training old skills (repeated practices of a skill over time, which requires append-only facts).

## Pros and cons

| Pros | Cons |
|---|---|
| Open, shared identifiers (ESCO, Wikidata) | Searching is messy: Wikidata's hierarchy is crowd-edited, with shortcuts and cycles |
| Multilingual for free (labels in ESCO and Wikidata) | ESCO is very coarse; level is approximate |
| Every fact carries its author and date | Free text (commits, training descriptions) must be mapped to identifiers |
| Agnostic of organisation and storage | Identifiers are external: Wikidata items can be merged or redirected |
| Progress graphs are plain aggregations over facts | Trust in a fact is left to the reader (ADR-0004) |

## Outside the model

- Requirements (prerequisites, job offers, company needs) are held by requesters (US-04).
- Assessments and goals are held by requesters (ADR-0009).

## Open

- Exchange format: the data model serialised (US-07); view freshness (ADR-0011), access control (ADR-0013).
