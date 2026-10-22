# US-09: Review imported facts

- Status: Draft
- Catalogue: UC7 bootstrap from a CV
- Actor: Anne ([personas](../personas.md))
- Follows: [US-01 import from a CV](US-01-import-cv.md); used by [US-08 new CV](US-08-new-cv.md)

## Story

As Anne, I want to check what the import understood from my CV, at my own pace, so that my record holds what I stand behind.

## Rules

- Review is not a gate. Imported facts that are not yet reviewed are usable.
- Views include proposed facts by default; shared views (with a grant) exclude them by default. Both are configurable.
- Every review operation is available, from a plain "ok" to splitting or expanding a fact.
- Levels are set by Anne, or proposed by the import when the CV allows it (e.g. "5 years of Rust").
- Each imported fact shows its fidelity, how closely it matches what the CV says: 1 low, 2 average, 3 good. A fact reviewed by Anne is always 3.
- A term the import did not understand: Anne resolves it into a fact, or abandons it; abandoning records nothing.
- What Anne does with a fact that is true but unwanted is her choice: reject it, or keep it and leave it out of a view.

## Review operations

| Operation | Example | Result |
|---|---|---|
| Accept | "develop software · publish–subscribe · Rust" is right | the fact is accepted |
| Reject | "manage databases" is wrong, or unwanted | the fact is rejected |
| Edit, then accept | wrong skill, period or level | the proposed fact is changed, then accepted |
| Split | "Kafka pipelines in Rust and Go" → one practice in Rust, one in Go | the fact is replaced by several, same excerpt |
| Expand | "Kafka pipelines" also means "designed a pub/sub system" | new facts added, same excerpt |
| Accept a group | everything in one occupation is right | all its facts are accepted |
| Resolve a term | "Zephyr RTOS" → Anne picks a concept | a new fact |
| Complete, then accept | "Rust, Go, C" names three skills with no period: Anne sets the start date and the level of each | the facts are accepted |
| Abandon a term | "Zephyr RTOS" is not worth recording | nothing recorded |
| Clean up a failed import | the import failed halfway | all facts proposed by that import are rejected |

## Data kept from the CV

The raw CV text is not kept once the import task ends. Each fact keeps its excerpt and where it comes from in the file (file name, line).

## Data in, data out

- In: the facts proposed by an import, with their excerpts and fidelity; terms not understood.
- Out: facts accepted, rejected, edited, split or added, at Anne's pace.

## Data model impacts

- Anne can correct any fact, reviewed or not, by editing it. Progress over time is a new practice.
- Views select by status; default: accepted and proposed; shared views: accepted only.
- Facts carry a fidelity, 1 to 3; 3 once reviewed.
- Split and expand create facts that share the original excerpt as evidence.
- Evidence of an imported fact: the excerpt, the file name, the line in the file.

## Open

- How the import computes fidelity.
