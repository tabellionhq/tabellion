# ADR-0002: Skill level scale

- Status: Accepted
- Date: 2026-10-15

## Decision

Skill levels use our own scale, 1 to 5:

| Level | Name | Meaning |
|---|---|---|
| 1 | Basic | requires guidance |
| 2 | Intermediate | autonomous: applies solutions |
| 3 | Average | experience: solves problems |
| 4 | Advanced | defines problems |
| 5 | Expert | finds problems |

## Alternatives

1. Revised Bloom taxonomy (Anderson & Krathwohl, 2001): 6 levels of thinking, remember, understand, apply, analyse, evaluate, create.
2. DigComp: 4 proficiency levels in DigComp 3.0 (basic, intermediate, advanced, highly advanced); 8 sub-levels in DigComp 2.1/2.2.
3. Our own scale.

## Reason

Progression from executing to discovering: first you are given orders and tools to fix a simple problem; then you are given harder problems and may need to make your own teams and tools; then you are asked to find new problems.

Levels 3 and 4 follow that order: solving a given problem comes before defining one.

## References

- DigComp 3.0: <https://joint-research-centre.ec.europa.eu/scientific-activities/key-competences-lifelong-learning/digital-competence-framework-digcomp/digcomp-30_en>
- [ADR-0001: Describe concepts with Wikidata](ADR-0001-describe-concepts-with-wikidata.md)
