# Use cases: data in, data out

- Status: Draft
- Sources: [personas and use cases](personas.md), [data model](data-model.md)

"Record" is the person's facts (Practice, Target). "View" is what a person exposes of their record (ADR-0007).

| UC | Actor | In | Reads | Out |
|---|---|---|---|---|
| UC1 recent commits | Claude for Anne | commit texts (possibly curated); session preferences | Anne's record | assessment; then answers → Practices |
| UC2 learn something new | Anne | a topic (e.g. Kafka); session preferences | Anne's record | assessment; then answers → Practices |
| UC3 prerequisites | Ismael | training prerequisites | Anne's view | gaps |
| UC4 during training | Ismael | training content; session preferences | Anne's view | assessment, practical work; then answers → Practices |
| UC5 training completed | Ismael | training objectives | Anne's view | gaps |
| UC6 signed skills | Ismael | Practices with signed evidence | | Practices proposed to Anne's record |
| UC7 bootstrap from a CV | Anne | CV text | | Practices proposed to Anne's record |
| UC8 new CV | Anne | | Anne's view | view data (rendering out of scope) |
| UC9a–c strategic analysis | Anne's company | embedded-systems requirements | coworkers' views | readiness; who is fit; who is trainable |
| UC9d orientation | Anne's company | embedded-systems requirements | Anne's view | Target proposed to Anne |
| UC10 yearly HR | Michaela | a period | coworkers' views | progress per person |
| UC11 recruitment | Michaela | Rust, Kafka, embedded-systems requirements | candidates' views | matching candidates |

## Data missing from the data model

- **Assessment**: questions or practical work, each tied to an ESCO skill, a Wikidata concept and a level aimed, with its sources (e.g. commit references). UC1, UC2, UC4.
- **Requirements**: a set of skill, concept, level not tied to a person (training prerequisites and objectives, company needs, job requirements). Target is tied to a person. UC3, UC5, UC9, UC11.
- **Session preferences**: time available, difficulty, preferences. UC1, UC2, UC4.
- **Proposed facts**: facts written by someone else wait for the person's acceptance. UC6, UC7, UC9d. Acceptance itself is out of scope (ADR-0006); the model may need a state.

Derived, not stored: gaps, readiness, progress, matches. They are reads over records or views compared with requirements.

## Open

- UC11: candidates' views come from outside the company; exchange between systems is not covered yet.
