# ia-upskilling

Checks that a developer can explain code they merged from agent-written commits.

Exercises are generated from real commits of a project and target the developer's level.

## Principles

Free and open: see [APR-0001](docs/apr/APR-0001-free-and-open.md).

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
- An answer key sits in a collapsed section under the questions.

## Layout

```text
quizzes/<project>/<commit>-<topic>.md   one quiz per commit topic
sessions/<developer>/<date>-<commit>.md answers given and verdicts
docs/apr/                               Architecture Principle Records
docs/adr/                               Architecture Decision Records
docs/agdr/                              Agent Decision Records, pending review
docs/research/                          benchmarks and studies
```

## Usage

1. Open a quiz, answer the 3 questions without opening the answer key.
2. Record answers and verdicts in `sessions/<developer>/`.
