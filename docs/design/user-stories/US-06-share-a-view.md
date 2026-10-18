# US-06: Share a view

- Status: Draft
- Actor: Anne ([personas](../personas.md))
- Catalogue: UC3 prerequisites, UC8 new CV, UC10 yearly HR assessment
- Decisions: ADR-0007 (the person decides what is exposed)

## Story

As Anne, I want to decide what of my curriculum I share, and with whom, so that each reader sees only what I chose.

## Concept

A view is a database view used for access control: a stored selection of Anne's facts (rows) and projection of their fields (columns), granted to a reader. Known as row-level and column-level security.

| View part | Query | Example |
|---|---|---|
| rules | selection on attributes (ABAC) | my work at ACME since January |
| include_facts | allow list: `OR id IN (…)` | these facts, picked for a job offer |
| exclude_facts | deny list: `AND id NOT IN (…)` | not this charity project |
| show | projection | levels yes, evidence no |
| grantee | grant | Michaela, Ismael, a recruiter, public |

The deny list wins over rules and allow list.

## Examples

| Reader | Built from | Catalogue |
|---|---|---|
| Michaela | rules: work at ACME, this year; evidence hidden | UC10 |
| Ismael | rules: skills under an ESCO category related to the training | UC3 |
| A recruiter | allow list: facts suggested from the job offer (F05), picked by Anne | UC8 |

## Data in, data out

- In: Anne's words ("only my work at ACME since January, no evidence"), or her picks among suggestions; the reader.
- Out: a view, granted to the reader.

## Acceptance

- Anne's words become rules written by an agent; Anne confirms them before the view is granted.
- Suggestions (e.g. from a job offer) are a starting point; only facts Anne picks enter the allow list.
- A view is applied by a deterministic query, never by an agent at read time.
- New facts matching the rules enter the view automatically; new facts never enter an allow list by themselves.
- Anne can change or revoke a view at any time; readers see the change at their next read.

## Functions required

- F05 Find competencies in a record related to a text (suggestions).
- F06 Write a view.
- F07 Read a record through a view.

## Data model impacts

- View.include_facts: allow list.
