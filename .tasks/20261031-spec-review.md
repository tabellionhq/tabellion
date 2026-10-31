# Task: Specification review

- Date: 2026-10-31
- Model: Claude Opus 5.5 (claude-opus-5-5)
- Tool: Claude Code; review by a general-purpose sub-agent, read-only
- Status: Open
- Roadmap: Specification review
- Scope: `docs/RFC.md`, `api/openapi.yaml`, `api/contexts/curriculum.jsonld`, ADRs, APRs, design documents, roadmaps, README. Not checked in depth: ADR-0020, 0023, 0026, 0027, 0029, 0034; screens 02, 03, 05; US-02, US-04.

Precedence: specification, then contract, then everything else. F-numbers are the reviewer's.

## Mechanical

- [ ] F3 Identity per view, in ADR-0010:24, cv-standards.md:36, US-07:37, US-08:16, screens/06-export.md:87: reword to "per grant (reader) or per export (owner)"; update note on ADR-0010.
- [ ] F5 `resources.md`: keep its rules, link the contract, drop the route table (routes missing or not in the contract, stale "deleting not in the contract yet", PATCH media type). `data-model.md`: replace the table with a link to Sections 3 and 4; drop YAML copies of the contract (evidence, task result).
- [ ] F6 JSON-LD context: `from` and `to` typed `xsd:date` but may be YYYY or YYYY-MM; add `CancelAction`; remove the unused `proposed` term.
- [ ] F12 Contract aligned on the specification: stage progress `to do` vs `to-do`; `show.*` required in 3.7, optional with defaults in the contract; Identity and View ids `string` vs `uuid`; Task `person` missing; `updated` missing on server preferences; rejections action limited to interrupted or abandoned tasks, 4.7 allows any state.
- [ ] F12 ADR-0013: note that `socket` is also a reader method (grant auth), besides owner access.
- [ ] F13 "Accepted facts SHALL NOT be mutated", "frozen", "append-only" in ADR-0018:9,17-18, ADR-0024:20,38, data-model.md:5,96, mvp-session.md:90: update notes on the ADRs (ADR-0008 clarification, specification 4.1); reword the design documents.
- [ ] F14 Stale against later decisions: ui-functions.md:7 (ADR-0021: `/ui` calls the core), cv-standards.md:12,47,54,59 (ADR-0028: rendering in the `cv` plugin), mvp-session.md:126 ("no summary", now `description`), adapters.md:16,29 and ROADMAP "Where the data goes" (now specification 4.12, 11; keep the UI display part), "one database engine" (ADR-0035: two files), ROADMAP-CV.md:8-9 (summary covered by `description`).
- [ ] F15 user-stories/README.md:96-101 "Data missing": keep only what is open; US-01:20 identity from the CV: roadmap (ROADMAP-CV); "grantee" → "reader (grant)" in functions.md:14-15, US-03:18.
- [ ] F16 US-11 "Resources this suggests": link the contract tags instead; US-09:15,49: link specification 4.5 instead of restating it.
- [ ] F17 References: ADR-0011:11 ADR-0012 → ADR-0013; US-11:55 Section 4.10 → 4.12; US-11:58 Section 4.7 → 4.12; RFC 11 "(Section 4.11)" on plugin access → 4.10; ADR-0012:23 `POST /practices` example.
- [ ] F18 Progress notes out of durable documents: ui-api-map.md:14, 129-130; mvp-session.md:154, 173-176, 187; resources.md:17; ui-functions.md:44-45. Contract: drop "(ADR-0013)" and the "parameters to define" placeholder on `oidc`.
- [ ] F19 RFC form: 4.4.1 and 4.4.2 one heading level down; RFC 9457 to normative references; cite [ESCO] and [WIKIDATA] in 4.2.

## Owner decisions

- [ ] F1 The owner creates accepted facts (4.1), the contract creates every fact proposed: an optional `status` (owner only), or the owner's facts created accepted.
- [ ] F2 `?identity=` on `/views/{id}/resume` lets a reader pick any identity (against 3.8, 11): under a grant, the grant's identity; `?identity=` for the owner only.
- [ ] F4 Merge patch schemas miss fields (occupation `description`, practice `occupation`) and status back to proposed: allow it, or state in 4.1 that a status does not return to proposed.
- [ ] F6 JSON-LD `occupation`: an ESCO occupation on an occupation held, the occupation held on a practice. Keep one term, or split.
- [ ] F7 Grant consent: withdrawing it revokes the grant, or 4.12 states that it is withdrawn by revoking only.
- [ ] F8 Owner authentication (10: authenticate every request, also on localhost): declared in the contract (session, socket), or outside the binding.
- [ ] F9 Wikidata lookups send CV terms online (against 11): a consent under 4.12, or 11 names reference-data lookups.
- [ ] F10 Saved CVs planned as preferences, against "preferences do not hold documents" (4.11): references only (identity, view, theme), or stored elsewhere.
- [ ] F11 Form of `author`: task URI and person UUID suggested; align examples (resources.md:63, mvp-session.md:66,90).
- [ ] F12 Language of UI messages: the stored `languages` (4.11) or `Accept-Language` (ADR-0019:28, ui-functions.md:24).

## AI (APR-0007)

Gaps in the specification, to work through with APR-0007: AI is a helper; everything can be done by a human or through index-based search.

- [ ] Search contract: F02 find the skill meant by words, F05 related competencies (US-02); index-based by default.
- [ ] Embeddings of personal data: who computes, who stores, part of the record or rebuildable.
- [ ] Model and version behind a vector; rebuild when it changes.
- [ ] Term matching (7.2): the server checks each excerpt is in its line; candidates and confidence; which plugin proposed a fact.
- [ ] US-06 rules into the specification: the owner confirms rules written by an agent; a view is applied by a deterministic query, never by an agent at read time.
- [ ] Agent clients: how they identify themselves in `author` and the access log.
- [ ] Regulation, uncertain: EU AI Act, recruitment use (UC11).
