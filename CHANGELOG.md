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
- Wireframe: review screen, with its accessibility rules.
- The owner corrects any fact by editing it, whatever its status; progress over time is a new practice (specification 4.1, ADR-0008 clarification). Owner rights in the specification are written as server duties.
- Tasks: canonical states ongoing, finished, abandoned, interrupted; stages and transitions defined by each kind of task; the CV import gains a confirm stage for the extracted text.
- Fidelity removed. An import fills only what its source states; completion, computed from the fields, measures each fact and the record (specification 4.8).
- Tasks are kept with their result until the owner deletes them; terms an import did not match stay there until resolved or left out.
