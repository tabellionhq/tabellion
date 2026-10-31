# Changelog

All notable changes to this project are documented here. Versions follow [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added

- Quiz format and first quiz: pp_rust_tracer `022318b`, `main` exit code.
- First session record.
- APR-0001 free and open.
- Agent Decision Records AgDR-0001 to AgDR-0003, pending review.
- Skills standards benchmark.
- ADR-0001 describe concepts with Wikidata.
- ADR-0002 skill level scale, 1 to 5.
- ADR-0003 professional skills referential: ESCO.
- ADR-0004 trust and progress are out of scope.
- ADR-0005 scope is skill curriculum recording: interfaces and data model.
- ADR-0006 model agnostic of the human organisation.
- Data model draft.
- Personas and use cases.
- ADR-0007 the person decides what is exposed.
- Design note: data first.
- Use cases in/out analysis.
- Design directory; US-01 import from a CV; functions F01–F03.
- User stories US-02 to US-07; functions F04 to F10; data model: occupations held, periods, status, identities, views, grants, reference data.
- APR-0002 simple adoption curve; APR-0003 skills curriculum standard, existing standards for the rest.
- ADR-0008 append-only normalised facts; ADR-0009 track skills, not assessments; ADR-0010 person identifier; ADR-0011 view freshness; ADR-0012 HTTP API; ADR-0013 access control; ADR-0014 Go.
- ADR-0015 API maturity: Richardson level 3, JSON-LD with schema.org actions; ADR-0016 core web UI in the main binary, one-time token swapped for a session cookie; ADR-0017 where a feature belongs: core, adapter or external system, no plugins; ADR-0018 concepts composed of Wikidata concepts.
- Design: resources exposed by the HTTP API; fact identifiers; accepted facts are never mutated.
- API contract draft for the MVP: `api/openapi.yaml` (OpenAPI 3.1, actions per resource), JSON-LD context `api/contexts/curriculum.jsonld`.
- APR-0004 international. US-08 new CV, US-09 review imported facts; MVP session map. Data model: tasks, fidelity, proposed facts editable until reviewed, views filter by status. ADR-0019 internationalisation (proposed).
- ADR-0008 clarified: deleting facts is allowed. ADR-0020 target platforms: Windows, macOS, Linux.
- Design note: CV standards. Roadmap: glossary, CV round trip.
- APR-0005 hosted service possible. MVP requirements: decided, to decide, open.
- Project name: Tabellion, replacing the working name `skills-track`. ADR-0021 frontend framework (proposed).
- ADR-0022 store engine: ncruces/go-sqlite3, PostgreSQL later. ADR-0023 HTTP stack (proposed). Test data: Anne's facts and CV.
- Data model: human languages and undated skills are practices completed in review; a proposed practice may miss its start date; evidence line; view periods match by overlap.
- Persona Thomas; US-10 install an app, run no server: apps are separate clients, the server is the authority. UI functions.
- ADR-0024: a specification is the source of truth for the model and the rules; the OpenAPI contract interprets it.
- Specification draft-tabellion-curriculum-00: `docs/RFC.md`, CC BY 4.0.
- Data model and resources notes link to the specification instead of restating it; the contract links to it.
- ADR-0025 licences: CC BY 4.0 for the specification and documentation, Apache-2.0 for the contract and test data, AGPL-3.0 for the server; contributions under the DCO.
- CONTRIBUTING: licences, DCO sign-off, short-lived branches, conventional commits with gitmoji, semantic versions.
- ADR-0021 web UI: HTMX, HTML adapter under /ui next to the JSON-LD API under /api. ADR-0023 HTTP stack: standard library, oapi-codegen strict server. ADR-0026 git workflow: GitHub flow.
- ADR-0019 internationalisation: GNU gettext PO files.
- APR-0006 accessibility. ADR-0027 accessibility standard: WCAG 2.2 AA.
- ADR-0028 core and plugins as services: Go core, optional plugin services (`cv` in TypeScript, `model` in Rust); demo compose runs the core and both.
- Specification: plugin defined next to client; rules for plugins (4.10); contracts for CV extraction, term matching and CV rendering (7); sections 7 to 12 renumbered 8 to 13.
- Design review: README describes Tabellion, knowledge checks become a potential use case; adapters, data model, MVP session and US-08 follow ADR-0028 (plugins); UI to API map with a build order.
- Vocabulary IRI: https://tabellion.app/vocab#.
- Wireframe: review screen, with its accessibility rules.
- The owner corrects any fact by editing it, whatever its status; progress over time is a new practice (specification 4.1, ADR-0008 clarification). Owner rights in the specification are written as server duties.
- Tasks: canonical states ongoing, finished, abandoned, interrupted; stages and transitions defined by each kind of task; the CV import gains a confirm stage for the extracted text.
- Fidelity removed. An import fills only what its source states; completion, computed from the fields, measures each fact and the record (specification 4.8).
- Tasks are kept with their result until the owner deletes them; terms an import did not match stay there until resolved or left out.
- Level 3 renamed from Average to Experienced (ADR-0002).
- Dates of periods take the precision their source gives: year, year-month or full date.
- Specification: a server sends no usage data (no telemetry).
- Review screen wireframe removed; screen 04 covers the review.
- Leave out a term an import did not match: `DELETE /tasks/{taskId}/terms/{termId}` (specification 4.7).
- Licence texts in `licences/`: AGPL-3.0, Apache-2.0, CC BY 4.0 (ADR-0025).
- Specification: CV extraction and CV rendering move out to the CV adapter contracts (`docs/design/cv-adapter.md`); term matching, on text, becomes 7.2.
- Delete a fact: `DELETE /practices/{id}`, `DELETE /occupations/{id}`; an occupation with practices is deleted only if its practices are moved or deleted in the same call (specification 4.1).
- `GET /server`: demo mode and where records are stored, for the owner.
- Screens 03 confirm and read the CV, 06 export a new CV.
- CI: DCO check of every new commit (AgDR-0004).
- ncruces/go-sqlite3 checked: encryption, FTS5, vec1, start time, builds (AgDR-0005).
- ADR-0029 Go toolchain and store dependencies: the latest Go release (1.27), each new one within 3 months; ncruces/go-sqlite3 with Adiantum, FTS5, vec1.
- CV roadmap (`ROADMAP-CV.md`): everything the CV adapter needs, built after the core.
- Plugin HTTP binding, accepted: ADR-0030 self-description, ADR-0031 slow calls, ADR-0032 who may call a plugin, ADR-0033 contracts and versions.
- ADR-0034 Go library for PO files: gotext.
- ADR-0035 preferences store: preferences and records in two SQLite files, one key; a table per kind of plugin.
- US-11 preferences: languages, plugin preferences, devices, export, consent.
- Specification: preferences (3.10, 4.11) and consents (3.11, 4.12); plugin calls carry the preferences of their kind (4.10, 7.1).
- API: preferences of the server and of each kind of plugin; consents, given and withdrawn.
- Specification: a grant states how its reader authenticates (`auth`, 3.8); occupations held take an optional description, Markdown recommended (3.4); a view shows it or not (`show.description`, 3.7).
- API: grants per view, with auth methods socket and bearer (OpenID Connect to define), revocation and access log; occupation description.
