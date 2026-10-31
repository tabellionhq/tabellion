# ADR-0036: Architecture of AI functionalities

- Status: Accepted
- Date: 2026-10-31
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

AI is a dependency, but a floating one:

- Functionally, several technologies are called "AI": full-text and rule-based search, embeddings, language models.
- Technologically, each is reached in several ways: on the machine or at a provider, open or commercial.
- Strategically, they must not be forced into the core model (APR-0007): everything can be done by a person editing facts, or through index-based search.

Functions that AI could serve ([functions](../design/functions.md)):

| Function | AI technology | Without AI |
|---|---|---|
| F01 read a text into facts | language model | the owner types the facts, helped by search |
| F02 find the skill or concept meant by words | embeddings | full-text index on ESCO and Wikidata labels |
| F05 related competencies in a record | embeddings | index and the ESCO hierarchy |
| F06 write a view from words | language model | the owner picks filters |
| Writing text, such as an occupation description | language model | the owner writes it |

Today only F01 has a contract (specification, Section 7.3). Concept embeddings sit in the `model` plugin without a contract (ADR-0028), apart from the core's search index.

Whether "AI" applies to a plugin is fuzzy: two plugins of the same kind may run rules or a language model, and look the same to the owner.

## Options

| Option | Issues | Impacts |
|---|---|---|
| A. AI in the core | Brings models, providers and their licences into the binary | Against APR-0007 and ADR-0028; the model depends on AI |
| B. Functions behind ports; AI adapters among others | A contract per port the server needs | The core model knows ports only; index, embeddings or a language model are interchangeable |
| C. A separate AI gateway | A second policy point next to the core; the core already calls every plugin (Section 4.10) | One more service to run; provider access belongs inside a plugin |

In B, AI sits in two places, neither in the model:

- Driven ports, called by the server: only the functions the server needs, term matching (F01) and search (F02, F05). The core ships an adapter without AI for each.
- Driving clients: the rest. A client, such as an agent or the web UI, turns words into a view or writes a description, then calls the API like any client; the owner confirms (Section 4.9).

## Recommendation

B.

## Decision

Functions are behind ports: AI functions may or may not participate in the ports. This allows to keep the model free and respect our APR. AI-free core adapters, AI clients for the rest, plugins declare their technologies.

## Reason

Keeps the model free and respects our APR.

## Consequences

- The specification defines ports by function, never by technology. No vector, model name or confidence enters the data model.
- Each driven port has an adapter without AI in the core.
- A plugin declares the technologies it runs in its self-description (ADR-0030).
- Writing views and text from words is done by clients, through the API.
- The core stays the only caller of plugins, and the place where their results are checked. Provider access, local or remote, stays inside a plugin.

## Next steps

- Search contract: a `search` kind for F02 and F05, with the core's index as its adapter without AI.
- Self-description: the field for declared technologies.
- Rules for every helper, AI or not: excerpts checked against their line (Section 7.3); the plugins a task called (Section 3.9); views applied from their fields alone; facts from an agent proposed.
- Embeddings of personal data: who builds the index, where it is kept.

## References

- [APR-0003: Skills curriculum standard, existing standards for the rest](../apr/APR-0003-curriculum-standard-defer-the-rest.md)
- [APR-0007: AI as a helper](../apr/APR-0007-ai-as-a-helper.md)
- [ADR-0017: Where a feature belongs](reversed/ADR-0017-where-a-feature-belongs.md)
- [ADR-0018: Concepts missing from Wikidata](ADR-0018-concepts-missing-from-wikidata.md)
- [ADR-0028: Core and plugins as services](ADR-0028-core-and-plugins-as-services.md)
- [ADR-0030: Plugin self-description](ADR-0030-plugin-self-description.md)
