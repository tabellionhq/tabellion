# US-01: Import from a CV

- Status: Draft
- Catalogue: UC7 bootstrap from a CV
- Actor: Anne ([personas](../personas.md))

## Story

As Anne, I want to upload my CV and review what is extracted, so that my curriculum is bootstrapped without typing it all.

## What the CV holds

| CV content | Example | Becomes |
|---|---|---|
| Occupations with periods | "Backend developer, ACME, 2021–2024" | occupation held by Anne, from–to, work |
| Skills and technologies per occupation | "Built a Kafka pipeline in Rust" | Practices: skill, concept, context, at that occupation, over that period |
| Personal projects | "Maintainer of a laser tracer" | occupation, pleasure |
| Education, certifications, trainings | "MSc Computer Science, 2016–2018" | occupation (student, trainee) over that period; skills gained there are Practices at that occupation |
| Languages, human and computer | "English: fluent", "Rust" | Practices: human languages as ESCO language skills, computer languages as Wikidata context |
| Identity and contact | name, email, phone, address | not stored: Person is opaque |

## Data in, data out

- In: a CV document.
- Out: proposed facts in Anne's record; terms that could not be resolved.
- The import cannot be done without Anne's review: she augments, filters and validates the proposed facts. The user interface matters for Anne from this first story.

## Functions required

1. Extract statements from a document: occupations, periods, skill and technology terms, each with the excerpt it comes from.
2. Resolve a term to identifiers (ESCO skill or occupation, Wikidata concept or context), with candidates and a confidence. This is the search of US-02.
3. Propose facts to the person's record, status proposed, the CV excerpt as evidence.
4. Review proposed facts: Anne augments, filters, validates.

ADR-0006 leaves out who accepts facts in an organisation; Anne reviewing her own import is a function of the system.

## Data model impacts

- **Occupation held** is missing: person, occupation, kind (work, pleasure), from, to.
- **Practice over a period**: a CV gives periods, not dates.
- **Level**: a CV rarely states one; null by default, or self-declared ("fluent").
- **Evidence**: the CV excerpt behind each fact.
- **Author**: Anne states the facts; an extractor wrote them. Both may need recording.

## Scope

Import is a one-time bootstrap. Afterwards the record is the source: CVs are generated from it, up to date and filtered for a specific job offer (UC8).

## Open

- None yet.
