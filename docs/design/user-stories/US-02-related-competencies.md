# US-02: Competencies related to recent work

- Status: Draft
- Actor: Claude ([personas](../personas.md))
- Catalogue: UC1 recent commits understood

## Story

As Claude, I want to search in Anne's curriculum which are her competencies related to a text resulting from a synthesis of multiple commits, so that the assessment builds on what she already masters and targets the rest.

## Data in, data out

- In: a text, Claude's synthesis of multiple commits; Anne's curriculum.
- Out: Anne's competencies related to the text, each with its level, date, and why it is related.

## Acceptance

- Related, not only identical: a text about publish–subscribe finds Anne's practice of the observer pattern.
- No false friends: a text about Rust does not find "rust proofing".
- Any language: Anne's competencies recorded in French are found from an English text.
- Empty is a valid answer: Anne has no related competency.

## Functions required

- F05 Find competencies in a record related to a text.

## Open

- Should the answer also list what the text involves that Anne's curriculum does not hold?
