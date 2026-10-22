# Data model

- Status: Draft
- Specification: [docs/RFC.md](../RFC.md) prevails over this note
- Decisions: ADR-0001 (Wikidata concepts), ADR-0002 (levels 1–5), ADR-0003 (ESCO), ADR-0004 to ADR-0007 (scope, views); ADR-0008 (append-only facts), ADR-0009 (skills, not assessments), ADR-0010 (person identifier), ADR-0018 (composed concepts)
- Stories: US-01 to US-09

A person's track record is a set of facts. How facts are proposed, reviewed, accepted and stored is outside the model (ADR-0006).

## Objects

### Facts

We work by adding facts: once accepted, a fact SHALL NOT be mutated; a correction is a new fact. Deleting facts is allowed: the owner deletes facts, e.g. by a filter (ADR-0008). A proposed fact can be edited until it is reviewed; its status then changes once, to accepted or rejected, and it is frozen. A person's progress is the list of their facts (ADR-0008).

```yaml
Person:
  id: uuid                # UUID v4, the only identifier; on every fact (ADR-0010)

Identity:                 # how a person presents themself to a reader; several per person (ADR-0010)
  id: string
  person: Person.id
  label: string           # a name, a pseudonym
  contact: string | null  # one email, a phone…

OccupationHeld:           # Anne was a backend developer at ACME from 2021 to 2024
  id: uuid
  person: Person.id
  occupation: uri | null  # ESCO occupation
  label: string           # as stated, e.g. "Backend developer, ACME"
  kind: work | pleasure | study   # study: student, trainee, training
  from: date
  to: date | null
  author: string          # a person, or the task that proposed it (/tasks/{id})
  recorded: datetime
  status: proposed | accepted | rejected
  fidelity: 1 | 2 | 3     # how closely it matches its source: 1 low, 2 average, 3 good; 3 once reviewed

Practice:                 # I practiced <skill> on <concept> in <context> at <occupation> over <period>
  id: uuid
  person: Person.id
  skill: uri              # ESCO skill, e.g. esco:skill/…
  concept:                # composed of Wikidata items (ADR-0018)
    is: uri               # genus, e.g. wd:Q1071360 publish–subscribe pattern
    with: [uri]           # what sets it apart; empty when the genus is the concept
    label: string | null  # e.g. "block value", for display; may be fed to the embedding engine
  context: uri | null     # Wikidata item, e.g. wd:Q575650 Rust
  occupation: OccupationHeld | null
  from: date              # may be missing while proposed; required to accept
  to: date | null         # null: a single date, or ongoing
  level: 1..5 | null      # ADR-0002; preset, set by the person, set externally, or empty
  evidence: object | null # any JSON, e.g. a CV excerpt {excerpt, file, line}, a signed token with claims, issuer, signature
  author: string          # who states the fact: the person, a training provider, an agent…
  recorded: datetime
  status: proposed | accepted | rejected

```

### Views

What a person exposes of their record, to whom (ADR-0007). A view is a database view used for access control: a stored selection (rows) and projection (fields) of the person's facts, granted to a reader. A view is not a fact: its owner changes it freely.

```yaml
View:
  id: string
  person: Person.id
  include:                # rules on dimensions; null: no restriction
    occupation_kind: [work | pleasure | study] | null
    occupation: [OccupationHeld] | null
    period: {from: date | null, to: date | null} | null   # a fact is in if its period overlaps this one
    skill_under: [uri] | null     # ESCO categories
    status: [proposed | accepted]  # default: both; a view with a grant: accepted only
  include_facts: [fact id]        # allow list
  exclude_facts: [fact id]        # deny list
  show:
    level: bool
    evidence: bool
```

```yaml
Grant:                    # who reads a view (ADR-0013)
  id: uuid
  view: View.id
  reader: string          # a person, an organisation, a system
  identity: Identity.id | null    # how the person presents themself to this reader
  auth: socket | bearer | oidc
  status: active | revoked
  expires: datetime | null

AccessLog:                # who read which view, when
  grant: Grant.id
  at: datetime
```

Semantics: (rules OR include_facts) AND NOT exclude_facts, then the projection given by `show`. The deny list wins.

Rules may be written by an agent from the owner's words; the owner confirms them.

### Tasks

Slow work run by the server, e.g. a CV import (US-01). A task is not a fact. Facts it proposes have the task as `author`: a failed task's facts can be rejected together (US-09).

```yaml
Task:
  id: uuid
  person: Person.id
  kind: cv-import
  state: queued | running | done | failed
  started: datetime
  ended: datetime | null
  result: object | null   # per kind; cv-import: {unresolved: [{term, excerpt}]}
  problem: object | null  # RFC 9457 problem details, when failed
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

- **Competency**: a person's latest accepted level per skill, concept and context, computed from their Practices.
- Gaps, readiness, progress, matches: reads over facts.

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

## Rules

- A reader reads a record only through a view, under a grant. The owner needs no grant: she reads her whole record, or any of her views.
- The view filters facts first; search and derived data use the filtered facts only. Never search first and filter afterwards: scores, counts or a truncated result list would reveal hidden facts.
- Derived data seen through a view (e.g. competency) is computed from the visible facts only.
- Fields hidden by `show` are removed from results.
- Views are applied by a deterministic filter, never by an agent at query time.

- No action list: the action is the ESCO skill.
- Search: ESCO skills by index (exact IDs, ancestors); concepts by semantic search over the vectors of their Wikidata IDs and of their label. Label vectors belong to the search adapter (ADR-0008).
- `level` is always an integer from 1 to 5, or null.
- Human languages are Practices: the ESCO language skill, the Wikidata language as concept, their own period (native: since birth).
- A skill named without a period (e.g. a CV summary) is proposed without a start date; the person completes it before accepting.
- `evidence` is opaque to the model: it is stored and returned, never interpreted (ADR-0004).
- Who set a level is the fact's `author`.
- Status moves once: `proposed` → `accepted` | `rejected`. An accepted or rejected fact SHALL NOT be mutated.
- Status, kind, state and fidelity are named values in the model and the API. How a store encodes them (e.g. small integers) is the store adapter's choice; the UI translates their names for display (ADR-0019).
- Responses embed the label of each referenced skill, occupation and concept, in the requested language; the store keeps identifiers only.
- Reference data: ESCO is shipped with the release; Wikidata items are looked up online and cached.

## Open

- Exchange format: the data model serialised (US-07); view freshness (ADR-0011), access control (ADR-0013).
