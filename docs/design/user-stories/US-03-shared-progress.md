# US-03: Shared progress for the yearly review

- Status: Draft
- Actor: Michaela ([personas](../personas.md))
- Catalogue: UC10 yearly HR assessment
- Decisions: ADR-0007 (the person decides what is exposed)

## Story

As Michaela, I want to see what each coworker chose to share of their progress this year, so that the yearly review rests on facts they agreed to show.

## Before

Each coworker has a view granted to Michaela. Anne wrote hers ([US-06](US-06-share-a-view.md)) in her own words ("only my work at ACME since January, no evidence"); an agent turned it into rules; Anne confirmed them.

## Data in, data out

- In: Michaela as grantee.
- Reads: Anne's facts filtered by the view Anne granted to Michaela.
- Out: the visible facts, hidden fields removed.

What Michaela views is out of scope (ADR-0005): its shape or form can range from a single PDF document to a complex interface showing all employees' progress.

## Acceptance

- Michaela gets nothing outside the view: not a fact, not a count, not a hint.
- Without a view granted to her, Michaela gets "nothing shared", not "no facts".
- Derived data (e.g. levels) comes from visible facts only.
- Evidence is absent when the view hides it.

## Functions required

- F06 Write a view from the owner's words.
- F07 Read a record through a view.
