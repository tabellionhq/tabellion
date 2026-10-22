# ADR-0017: Where a feature belongs: core, adapter or external system

- Status: Accepted
- Date: 2026-10-19
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

The core exposes a limited set of features. Adapters make them useful (CV import, CV export), united by the core web UI (ADR-0016). Some features need their own data: sharing the central store with them would couple their lifecycle to the curriculum's and break its rules (append-only facts, views, grants).

No user story needs third parties to add code to `tabellion`.

## Options

| Option | How | Pros | Cons |
|---|---|---|---|
| A. A rule by state, adapters compiled in | Core: the curriculum. Adapter: a stateless translator, compiled in. External system: anything with its own data, a client of the API under a grant. | No plugin system; no shared store; external systems already work through the API (ADR-0012, ADR-0013) | Adding a translator means a new release of `tabellion` |
| B. Plugin system for imported adapters | Versioned executables in a directory (like Terraform providers), started on demand, metadata at a `.well-known` route, web views in sandboxed frames at their own `*.localhost` origin | Third parties extend the binary | Specified without a use case; zero trust needs a container or an OS sandbox; plugins needing a store either share the central one or bring their own |
| C. Plugins with access to the central store | Plugins read and write the store directly | Easy for plugin authors | Bypasses views and grants; couples schemas; rejected |

## Recommendation

A.

| Kind | What it is | State | Examples |
|---|---|---|---|
| Core | The curriculum: facts, views, grants, review | The central store | Practice, View, Grant |
| Adapter | A translator between the model and an existing standard or tool (APR-0003) | None, or a cache it can rebuild (ADR-0008) | CV formats in and out; auth (bearer, OIDC); search (keyword, embeddings); ESCO and Wikidata loaders |
| External system | A product with its own data and lifecycle; a requester (ADR-0009) | Its own store | Learning system and goals; assessments (Ismael, Claude); HR dashboards (Michaela); recruitment |

Rule: a feature that needs its own store is an external system, never an adapter.

Reopen B when a translator that `tabellion` does not ship is needed and cannot be built as an external client.

## Decision

A. We won't do plugins: a feature that needs its own store is an external system.

## Reason

Sensible: we are interoperable, so an external system uses the API.

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [APR-0003: Skills curriculum standard, existing standards for the rest](../apr/APR-0003-curriculum-standard-defer-the-rest.md)
- [ADR-0008: Normalisation of the data model](ADR-0008-normalisation.md)
- [ADR-0009: Track skills, not assessments](ADR-0009-track-skills-not-assessments.md)
- [ADR-0012: HTTP API](ADR-0012-http-api.md)
- [ADR-0013: Access control](ADR-0013-access-control.md)
- [ADR-0016: Core web UI in the main binary](ADR-0016-gui-in-the-binary.md)
