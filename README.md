# Tabellion

Code: [github.com/tabellionhq](https://github.com/tabellionhq) · site: [tabellion.app](https://tabellion.app)

Checks that a developer can explain code they merged from agent-written commits.

Exercises are generated from real commits of a project and target the developer's level.

## Specification

[docs/RFC.md](docs/RFC.md): data model and rules of the skills curriculum, beta (`draft-tabellion-curriculum-00`), CC BY 4.0. It prevails over the API contract and the code.

## Licences

| Part | Licence |
|---|---|
| Specification and documentation | CC BY 4.0 |
| API contract, JSON-LD context, test data | Apache-2.0 |
| Server, web UI, command line | AGPL-3.0 |

Contributions are signed off under the Developer Certificate of Origin.

## Principles

- Free and open: see [APR-0001](docs/apr/APR-0001-free-and-open.md).
- Simple adoption curve: see [APR-0002](docs/apr/APR-0002-simple-adoption-curve.md).
- Skills curriculum standard, existing standards for the rest: see [APR-0003](docs/apr/APR-0003-curriculum-standard-defer-the-rest.md).

## Exercise types

| Type | The developer… | Status |
|---|---|---|
| Quiz | answers questions about code shown inline | in use |
| Interactive | writes or runs code | planned |
| Definition | defines a concept in their own words | planned |
| Input → output | predicts what comes out for a given input | planned |

## Quiz format

- 1 or 2 code blocks taken from one commit.
- 3 questions, easy to hard.
- Everything needed to answer is in the code blocks: no running code, no reading other files or functions.

## Layout

```text
docs/apr/       Architecture Principle Records
docs/adr/       Architecture Decision Records
docs/agdr/      Agent Decision Records
docs/research/  benchmarks and studies
docs/design/    personas, user stories, data model, functions, design notes
quizzes/, sessions/  sample exercises; storage is not decided
```
