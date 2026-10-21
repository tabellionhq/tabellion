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
