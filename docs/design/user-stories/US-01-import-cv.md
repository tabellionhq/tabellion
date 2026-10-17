# US-01: Import from a CV

- Status: Draft
- Catalogue: UC7 bootstrap from a CV
- Actor: Anne ([personas](../personas.md))

## Story

Anne wants to bootstrap her curriculum by uploading her CV.

## What the CV holds

| CV content | Example | Becomes |
|---|---|---|
| Occupations with periods | "Backend developer, ACME, 2021–2024" | occupation held by Anne, from–to, work |
| Skills and technologies per occupation | "Built a Kafka pipeline in Rust" | Practices: skill, concept, context, at that occupation, over that period |
| Personal projects | "Maintainer of a laser tracer" | occupation, pleasure |
| Education, certifications | "Kubernetes certification, 2023" | ? |
| Languages | "English: fluent" | ? (ESCO has language skills) |
| Identity and contact | name, email, phone, address | not stored: Person is opaque |

## Data in, data out

- In: a CV document.
- Out: proposed facts in Anne's record; terms that could not be resolved.

## Functions required

1. Extract statements from a document: occupations, periods, skill and technology terms, each with the excerpt it comes from.
2. Resolve a term to identifiers (ESCO skill or occupation, Wikidata concept or context), with candidates and a confidence. This is the search of US-02.
3. Propose facts to the person's record, status proposed, the CV excerpt as evidence.

Accepting proposed facts is outside the model (ADR-0006).

## Data model impacts

- **Occupation held** is missing: person, occupation, kind (work, pleasure), from, to.
- **Practice over a period**: a CV gives periods, not dates.
- **Level**: a CV rarely states one; null by default, or self-declared ("fluent").
- **Evidence**: the CV excerpt behind each fact.
- **Author**: Anne states the facts; an extractor wrote them. Both may need recording.

## Open

- Education and certifications: facts of their own, or evidence for Practices?
- Languages: in scope?
- Importing a second CV: duplicates, or updates?
