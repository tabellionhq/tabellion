# US-04: Prerequisites

- Status: Draft
- Actor: Ismael ([personas](../personas.md))
- Catalogue: UC3 prerequisites

## Story

As Ismael, I want to know if Anne matches the prerequisites of my training, so that I can confirm she's been oriented to the right training session.

## Before

- Anne granted Ismael a view of her record (F06).

## Data in, data out

- In: one training session's prerequisites, as text; Anne's view granted to Ismael.
- Out: Anne's visible competencies related to the prerequisites, as a graph (skills, concepts, levels, and how they relate).

The prerequisites stay on Ismael's side: they are not in Anne's curriculum. Comparing the graph with them is done by Ismael or his tools, outside the system.

## Acceptance

- Related, not only identical: "Docker basics" finds Anne's practice of containerisation.
- Only the view is used: a competency held by a hidden fact does not appear, and nothing hints at it.
- A practice without a level reads "practiced, level unknown", not level 1.

## Functions required

- F05 Find competencies in a record related to a text.
- F07 Read a record through a view.
