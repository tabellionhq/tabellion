# Adapters

- Status: Draft
- Decisions: ADR-0012 (HTTP API), ADR-0013 (access control), ADR-0014 (language), ADR-0016 (core web UI), ADR-0017 (core, adapter or external system), ADR-0018 (composed concepts)

Adapters connect the core to people, systems and standards. All are compiled into `tabellion` (ADR-0017). An adapter keeps no state of its own, or only a cache it can rebuild; anything that needs its own store is an external system, a client of the API under a grant.

- Driving: something calls the core.
- Driven: the core calls something.

| Side | Port | Adapters | State | Functions, ADR |
|---|---|---|---|---|
| Driving | Client | command line, web UI, MCP | none | all; ADR-0012, ADR-0016 |
| Driving | Auth | Unix socket, bearer token, OpenID Connect, UI one-time token | none; grants live in the core | ADR-0013, ADR-0016 |
| Driven | Store | one database engine | the central store | all; ADR-0008 |
| Driven | Reference data | ESCO (skills, ancestors, labels), Wikidata (labels) | rebuildable copy | F02 |
| Driven | Search | index: ESCO IDs and ancestors; embeddings, computed locally: Wikidata concepts and concept labels | rebuildable index and vectors | F02, F05; ADR-0018 |
| Driven | Extraction | CV formats in (JSON Resume, Europass, HR Open), free text and PDF | none | F01 |
| Driven | Rendering | CV formats out (Europass, JSON Resume, HR Open), Wikidata draft (QuickStatements) | none | ADR-0018; [cv-standards.md](cv-standards.md) |
| Driven | Event stream | e.g. an Atom feed of new facts | none | ADR-0008 |

## Rules

- Search: the view filters facts first; then ESCO skills by index, concepts by embeddings (data model rules).
- An adapter may call an external service, such as an LLM for extraction; that is its own design and configuration.
- Level 0 export is not an adapter: it is `GET /views/{id}/facts` as a file ([resources.md](resources.md)).

## Open

- MVP: the adapters US-01 and US-08 require.
