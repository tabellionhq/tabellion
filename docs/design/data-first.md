# Data first

Good software is done via smart data and dumb code.

## Position

Functional design decomposes a system into functions, each with one responsibility. Its weak point: data is subservient to functionality, treatments are seen as the main aspect.

This project designs data first. Facts reference open identifiers (ESCO, Wikidata), carry their author and evidence, and are exposed through views. Most behaviour is then a query over facts. Real treatments (tagging text with identifiers, generating and grading assessments) sit at the edges, as adapters.

## Lineage

| Who | When | Statement |
|---|---|---|
| David Parnas | 1972 | "On the Criteria to Be Used in Decomposing Systems into Modules": decompose by design decisions to hide, mostly data representations, not by processing steps. |
| Fred Brooks | 1975 | "Show me your flowcharts and conceal your tables, and I shall continue to be mystified. Show me your tables, and I won't usually need your flowcharts; they'll be obvious." (*The Mythical Man-Month*) |
| Michael Jackson | 1975 | Jackson Structured Programming: program structure follows data structure. (*Principles of Program Design*) |
| Alan Perlis | 1982 | "It is better to have 100 functions operate on one data structure than 10 functions on 10 data structures." (*Epigrams on Programming*) |
| Rob Pike | 1989 | "Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident." (*Notes on Programming in C*, rule 5) |
| Eric S. Raymond | 2003 | Rule of Representation: "Fold knowledge into data so program logic can be stupid and robust." (*The Art of Unix Programming*) |
| Linus Torvalds | 2006 | "Bad programmers worry about the code. Good programmers worry about data structures and their relationships." (git mailing list) |

Later descendants: event sourcing, data-oriented design, Rich Hickey's "data > functions > macros".

## Counterpoint

Domain-driven design (Eric Evans, Martin Fowler) calls data without behaviour an anemic domain model, an anti-pattern: rules belong with the data. Use cases still decide which data matters; data first does not mean use cases last.

## Method

For each use case, state which data goes in and which comes out ([use cases in/out](../spec/use-cases-io.md)). Missing data found this way goes to the [data model](../spec/data-model.md).
