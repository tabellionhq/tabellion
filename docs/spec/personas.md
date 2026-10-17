# Personas

Interfaces are designed from these personas and their scenarios.

## P1: Anne

28, queer, developer. Likes to build things and to learn a lot. Has a short attention span and quickly forgets. Learns fast but is often in a rush. Has open source, charitable and professional development projects.

Uses the system to check that what is committed is understood, and wants to use it to learn more.

## P2: Claude

5, agent. Knows approximately a lot, has no personality, has no memory of other people. Can synthesise and generate a lot very fast, and sometimes very wrong.

Its role in the system is to help create valuable skills assessments and record them. For now, Claude is bound to the code-production post-review workflow.

## P3: Ismael

43, professional teacher (formateur), trains by passion. Teaches DevOps, containers, cloud. Works for a training organisation (organisme de formation) that delivers in-person training for Anne's company.

## P4: Michaela

51, HR at Anne's company. Not so great with computers, but loves people and seeing them grow professionally. Has read-only access to the views each person exposes (ADR-0007).

## Use cases

### UC1: Recent commits understood

Claude is fired as a sub-agent, with a task to help Anne make sure all the recent commits in a project are well understood and mastered.

The commit list is a source of data. The skills assessment database is a real help: without it there is no way to provide an insightful assessment.

### UC2: Learn something new

Anne wants to learn about Kafka, for example, because why not.

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

Anne wants to create a new CV based on her curriculum.

### UC9: Strategic analysis

Anne's company integrates her with her coworkers into a strategic analysis. They want to assess:

- a. if their workforce is ready to work on embedded systems;
- b. who would already be fit to work on embedded systems;
- c. who would be fit to receive professional education in embedded systems, based on their skills track;
- d. to orient Anne to learn about embedded systems for strategic purposes.

### UC10: Yearly HR assessment

Anne's company wants to track who has progressed this year, and how, for a yearly HR assessment.

### UC11: Recruitment

Michaela looks on the market for new potential recruits who can do Rust, Kafka and embedded systems, and wants to use the skills assessment system to filter job applications.
