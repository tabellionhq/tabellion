# Task: Tabellion design backlog

- Date: 2026-10-26
- Model: Claude Opus 5.5 (claude-opus-5-5)
- Tool: Claude Code
- Status: Open
- From: [20261014-poc-knowledge-validation](20261014-poc-knowledge-validation.md), open items deferred

## Owner decisions

- [ ] Language model for CV extraction (ADR); a small local model to test first
- [ ] Concept search in the MVP: Wikidata online search or local embeddings
- [ ] CV prose (summary, highlights): written, generated, kept where (cv-standards.md, Questions)
- [ ] Level names: "Intermediate" and "Average" read alike (ADR-0002)
- [ ] Date precision of practices (year only?)
- [ ] Rule: no telemetry (specification or ADR)
- [ ] Consent before data leaves the local system (roadmap)
- [ ] Create GitHub `tabellionhq`; manual trademark check

## Specification and contract

- [ ] Term matching also returns the identity (name, contacts) of the CV
- [ ] Plugin contracts: HTTP binding
- [ ] Operation to leave out a term not matched (cv-import result); home shows terms not matched

## Screens

- [ ] Screens 03 (reading and confirm) and 06 (export, view editor) from the role-play: save from the session transcript if needed
- [ ] Export screen: focus a CV on a job (target ESCO occupation, job offer, saved views)
- [ ] Check screen: drag-and-drop library, and a way to do the same without dragging (WCAG 2.5.7)
- [ ] Review screen wireframe (superseded): redraw minimal, review all at once with pre-filled buckets, or drop

## Checks and setup

- [ ] ncruces/go-sqlite3: encryption (service key, person key), vector search, start time, binary size, memory
- [ ] GDPR household exemption for local use; GDPR duties for the hosted service
- [ ] Licence files (root, api, testdata, docs); DCO and licence checks in CI
- [ ] Demo compose: core, `cv`, `model`
- [ ] Before the first translated message: Go PO library with plural rules; later, a translation service for community translations
