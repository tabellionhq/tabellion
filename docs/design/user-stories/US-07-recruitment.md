# US-07: Recruitment

- Status: Draft
- Actor: Michaela ([personas](../personas.md))
- Catalogue: UC11 recruitment

## Story

As Michaela, I want to find among applicants those whose curriculum relates to Rust, Kafka and embedded systems, so that I spend time on the right applications.

## Scope

- In scope: applicants who shared a view with their application (US-06), e.g. a link in the application.
- Out of scope: discovering people who did not apply. Crawling and indexing public curricula is a requester-side system (e.g. a job board); the curriculum only makes public views readable.

## Data in, data out

- In: Michaela's search text ("Rust, Kafka, embedded systems"); the views applicants granted to her company.
- Out: per applicant, their visible competencies related to the text, with levels.

## Acceptance

- Only views granted by applicants are read: nobody is searched without having shared.
- A view may live in another system (an applicant's local curriculum, another company's): it is read through an open format and an address.
- The exchange format is the data model, serialised: flat facts with ESCO and Wikidata identifiers, and the view's header (who shared what with whom). Consumers store and search at scale on their side.
- Ranking or scoring applicants is Michaela's side (ADR-0004): the system returns related competencies per applicant, not a verdict.

## Functions required

- F05 Find competencies in a record related to a text.
- F07 Read a record through a view.
- F10 Serve a view to another system, in the exchange format.

## Data model impacts

- A view has an address a reader can use from elsewhere.
- A person needs an identifier that is unique across systems: a UUID; name and contact come from the identity chosen for the view (ADR-0010).
- A reader needs to know whether its copy of a view is fresh (ADR-0011); access and revoking a reader are access control (ADR-0013).
