# User stories

- Status: Draft
- Sources: [personas](../personas.md), [data model](../data-model.md)

Detailed stories: [US-01 import from a CV](US-01-import-cv.md), [US-02 competencies related to recent work](US-02-related-competencies.md), [US-03 shared progress for the yearly review](US-03-shared-progress.md), [US-04 prerequisites](US-04-prerequisites.md), [US-05 certification](US-05-certification.md), [US-06 share a view](US-06-share-a-view.md), [US-07 recruitment](US-07-recruitment.md), [US-08 new CV filtered for an offer](US-08-new-cv.md), [US-09 review imported facts](US-09-review-imported-facts.md), [US-10 install an app, run no server](US-10-hosted-mobile.md), [US-11 preferences](US-11-preferences.md).


## Catalogue

### UC1: Recent commits understood

Claude is fired as a sub-agent, with a task to help Anne make sure all the recent commits in a project are well understood and mastered.

The commit list is a source of data. The skills assessment database is a real help: without it there is no way to provide an insightful assessment.

### UC2: Learn something new

Anne wants to learn about Kafka, for example, because why not.

Outside the curriculum: the learning system uses it (ADR-0009).

### UC3: Prerequisites

Ismael wants to know if Anne matches the prerequisites of a training.

### UC4: During training

During the session, Ismael wants to assess Anne's understanding and create practical work (travaux pratiques) closely related to her capacities.

### UC5: Training completed

Post training, Ismael wants to assess that Anne has completed her progress.

### UC6: Signed skills

Post training, Ismael can assess with a signature the new skills gained by Anne, which she can integrate into her curriculum.

### UC7: Bootstrap from a CV

Anne wants to bootstrap her curriculum by uploading her CV.

### UC8: New CV

Anne wants to create a new CV based on her curriculum, up to date and filtered according to a specific job offer.

Detailed in US-08.

### UC12: Use as a service

Thomas installs an app and uses the hosted service, so that he runs no server.

Detailed in US-10.

### UC9: Strategic analysis

Anne's company integrates her with her coworkers into a strategic analysis. They want to assess:

- a. if their workforce is ready to work on embedded systems;
- b. who would already be fit to work on embedded systems;
- c. who would be fit to receive professional education in embedded systems, based on their skills track;
- d. to orient Anne to learn about embedded systems for strategic purposes.

Covered: a–c by US-06 and US-07, inside the organisation; aggregation, fit and training decisions are the company's side (ADR-0004), and results only include what employees shared (ADR-0007). d is outside the curriculum: goals belong to learning systems (ADR-0009).

### UC10: Yearly HR assessment

Anne's company wants to track who has progressed this year, and how, for a yearly HR assessment.

### UC11: Recruitment

Michaela looks on the market for new potential recruits who can do Rust, Kafka and embedded systems, and wants to use the skills assessment system to filter job applications.

## Data in, data out

"Record" is the person's facts. "View" is what a person exposes of their record (ADR-0007).

| UC | Actor | In | Reads | Out |
|---|---|---|---|---|
| UC1 recent commits | Claude for Anne | commit texts (possibly curated); session preferences | Anne's record | assessment; then answers → Practices |
| UC2 learn something new | Anne | a topic (e.g. Kafka); session preferences | Anne's record | assessment; then answers → Practices |
| UC3 prerequisites | Ismael | training prerequisites | Anne's view | gaps |
| UC4 during training | Ismael | training content; session preferences | Anne's view | assessment, practical work; then answers → Practices |
| UC5 training completed | Ismael | training objectives | Anne's view | gaps |
| UC6 signed skills | Ismael | Practices with signed evidence | | Practices proposed to Anne's record |
| UC7 bootstrap from a CV | Anne | CV text | | Practices proposed to Anne's record |
| UC8 new CV | Anne | | Anne's view | view data; rendering by a plugin, outside the specification |
| UC9a–c strategic analysis | Anne's company | embedded-systems requirements | coworkers' views | readiness; who is fit; who is trainable |
| UC9d orientation | Anne's company | embedded-systems requirements | Anne's view | Target proposed to Anne |
| UC10 yearly HR | Michaela | a period | coworkers' views | progress per person |
| UC11 recruitment | Michaela | Rust, Kafka, embedded-systems requirements | candidates' views | matching candidates |
| UC12 use as a service | Thomas | his actions in the app | his record | his record, on his hosted instance |

## Data missing from the data model

First analysis, before the detailed stories. Since then: requirements and assessments stay with the requester (US-04, ADR-0009); targets belong to learning systems (ADR-0009); proposed facts are covered by status (US-01).

- **Assessment**: questions or practical work, each tied to an ESCO skill, a Wikidata concept and a level aimed, with its sources (e.g. commit references). UC1, UC2, UC4.
- **Requirements**: a set of skill, concept, level not tied to a person (training prerequisites and objectives, company needs, job requirements). Target is tied to a person. UC3, UC5, UC9, UC11.
- **Session preferences**: time available, difficulty, preferences. UC1, UC2, UC4.
- **Proposed facts**: facts written by someone else wait for the person's acceptance. UC6, UC7, UC9d. Acceptance itself is out of scope (ADR-0006); the model may need a state.

Derived, not stored: gaps, readiness, progress, matches. They are reads over records or views compared with requirements.

## Open

- None: exchange between systems is covered by US-07.
