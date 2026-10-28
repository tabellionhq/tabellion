# Tabellion

Code: [github.com/tabellionhq/tabellion](https://github.com/tabellionhq/tabellion) · site: [tabellion.app](https://tabellion.app)

Records the skills a person has practised over time: occupations held, and practices of a skill on a concept, over a period, at a level. The record belongs to the person, who exposes part of it through views, such as a CV.

Skills are ESCO skills; concepts are Wikidata items.

## Specification

[docs/RFC.md](docs/RFC.md): data model and rules of the skills curriculum, beta (`draft-tabellion-curriculum-00`), CC BY 4.0. It prevails over the API contract and the code.

## Licences

| Part | Licence |
|---|---|
| Specification and documentation | [CC BY 4.0](licences/CC-BY-4.0.txt) |
| API contract, JSON-LD context, test data | [Apache-2.0](licences/Apache-2.0.txt) |
| Server, web UI, command line | [AGPL-3.0](licences/AGPL-3.0.txt) |

Contributions are signed off under the Developer Certificate of Origin: see [CONTRIBUTING.md](CONTRIBUTING.md).

## Principles

- Free and open: see [APR-0001](docs/apr/APR-0001-free-and-open.md).
- Simple adoption curve: see [APR-0002](docs/apr/APR-0002-simple-adoption-curve.md).
- Skills curriculum standard, existing standards for the rest: see [APR-0003](docs/apr/APR-0003-curriculum-standard-defer-the-rest.md).
- International: see [APR-0004](docs/apr/APR-0004-international.md).
- Hosted service possible: see [APR-0005](docs/apr/APR-0005-hosted-service-possible.md).
- Accessibility: see [APR-0006](docs/apr/APR-0006-accessibility.md).

## Potential use case: knowledge checks

Checks that a developer can explain code they merged from agent-written commits. Exercises are generated from real commits of a project and target the developer's level; a passed check can become evidence of a practice.

### Exercise types

| Type | The developer… | Status |
|---|---|---|
| Quiz | answers questions about code shown inline | in use |
| Interactive | writes or runs code | planned |
| Definition | defines a concept in their own words | planned |
| Input → output | predicts what comes out for a given input | planned |

### Quiz format

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
api/           API contract (OpenAPI) and JSON-LD context
testdata/      a sample record and the CV it was written from
licences/      licence texts
```
