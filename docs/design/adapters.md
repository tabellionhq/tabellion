# Adapters

- Status: Draft
- Decisions: ADR-0012 (HTTP API), ADR-0013 (access control), ADR-0014 (language), ADR-0016 (core web UI), ADR-0017 (core, adapter or external system), ADR-0018 (composed concepts), ADR-0028 (core and plugins as services)
- Specification: plugins, [RFC.md](../RFC.md) Sections 4.10 and 7

Adapters connect the core to people, systems and standards. Some are compiled into `tabellion`; others are plugins, optional services the core calls by URL (ADR-0028). An adapter keeps no state of its own, or only a cache it can rebuild; anything that needs its own store is an external system, a client of the API under a grant.

- Driving: something calls the core.
- Driven: the core calls something.

| Side | Port | Adapters | State | Where | Functions, ADR |
|---|---|---|---|---|---|
| Driving | Client | command line, web UI, MCP | none | core | all; ADR-0012, ADR-0016 |
| Driving | Auth | Unix socket, bearer token, OpenID Connect, UI one-time token | none; grants live in the core | core | ADR-0013, ADR-0016 |
| Driven | Store | one database engine | the central store | core | all; ADR-0008 |
| Driven | Reference data | ESCO (skills, ancestors, labels), Wikidata (labels) | rebuildable copy | core | F02 |
| Driven | Search | index: ESCO IDs and ancestors | rebuildable index | core | F02; ADR-0018 |
| Driven | Term matching | CV text to proposed facts; embeddings of Wikidata concepts and concept labels | rebuildable vectors, no personal data | `model` plugin | F01, F02, F05; ADR-0018 |
| Driven | Extraction | PDF, DOCX, ODT to text | none | `cv` plugin | F01 |
| Driven | Import of structured CVs | JSON Resume, Europass, HR Open | none | core | F01 |
| Driven | Export | CV formats out (JSON Resume, Europass, HR Open), Wikidata draft (QuickStatements) | none | core | ADR-0018; [cv-standards.md](cv-standards.md) |
| Driven | Rendering | JSON Resume to HTML or PDF with a theme | installed themes, a rebuildable cache | `cv` plugin | [cv-standards.md](cv-standards.md) |
| Driven | Event stream | e.g. an Atom feed of new facts | none | core | ADR-0008 |

## Rules

- Search: the view filters facts first; then ESCO skills by index, concepts by embeddings (data model rules).
- A plugin may call an external service, such as a language model provider; another provider is another plugin of the same kind. Data leaving the local system needs the person's agreement (roadmap).
- Level 0 export is not an adapter: it is `GET /views/{id}/facts` as a file ([resources.md](resources.md)).

## Open

- MVP: the adapters US-01 and US-08 require.
