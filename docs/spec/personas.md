# Personas

Interfaces are designed from these personas and their scenarios.

## P1: Anne

28, queer, developer. Likes to build things and to learn a lot. Has a short attention span and quickly forgets. Learns fast but is often in a rush. Has open source, charitable and professional development projects.

Uses the system to check that what is committed is understood, and wants to use it to learn more.

## P2: Claude

5, agent. Knows approximately a lot, has no personality, has no memory of other people. Can synthesise and generate a lot very fast, and sometimes very wrong.

Its role in the system is to help create valuable skills assessments and record them.

## Scenarios (draft)

### Anne

- **A1. Did I get it?** Anne merged agent commits on a project. Between two tasks, they ask to check. They get a few questions on the code they merged, answer in a few minutes, get corrections. A Practice is recorded.
- **A2. I forgot.** Weeks later, the same concept comes back in a short check, because Anne forgets quickly.
- **A3. Learn more.** From a concept Anne struggled with, they get a next step to learn it, not only a verdict.
- **A4. Many projects.** Anne's open source, charitable and professional projects feed one track record; professional code may not leave its project.
- **A5. My progress.** Anne sees a graph of their progress per skill and concept.

### Claude

- **C1. Read the work.** Claude reads commits and decision records of a project and proposes ESCO skills and Wikidata concepts they involve. Being sometimes very wrong, each proposal points to the code and to an existing identifier.
- **C2. Assess.** Claude generates questions at Anne's level, grades the answers, and proposes a Practice with level and evidence (questions, answers, verdicts). Anne accepts or rejects it.
- **C3. Start from the record.** Having no memory of Anne, Claude reads Anne's track record before assessing: levels, concepts forgotten, targets.

## Risks

- Anne, in a rush, accepts a wrong assessment from Claude.
- Claude proposes an identifier that does not exist or does not match the code.
