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
- [x] Create GitHub `tabellionhq`
- [ ] Manual trademark check

## Specification and contract

- [ ] Term matching also returns the identity (name, contacts) of the CV
- [ ] Plugin contracts: HTTP binding
- [ ] Preferences (3.10, 4.11): per-person entries, opaque to the specification, owner only, never exposed by views; e.g. `cv/1: {identity, view, theme, format}` for a saved CV; a plugin gets only the entry its call needs
- [ ] CV themes (npm): a few in the `cv` image, more installed on demand into a `cv-themes` volume (a rebuildable cache, 4.10); contract 7.4 lists installed and available themes and gains an install operation; risks: a theme is code run on the CV, installs reach npmjs.org (`--ignore-scripts`, pinned versions), theme licences vary
- [ ] Operation to leave out a term not matched (cv-import result); home shows terms not matched

## Screens

- [ ] Screens drafted in the role-play, not saved: 03 reading and confirm, 06 export (who, what, look, file, build); "My CVs" screen
- [ ] Check screen: drag-and-drop library, and a way to do the same without dragging (WCAG 2.5.7)
- [ ] Review screen wireframe (superseded): redraw minimal, review all at once with pre-filled buckets, or drop

## Checks and setup

- [ ] ncruces/go-sqlite3: encryption (service key, person key), vector search, start time, binary size, memory
- [ ] GDPR household exemption for local use; GDPR duties for the hosted service
- [ ] Licence files (root, api, testdata, docs); DCO and licence checks in CI
- [ ] Demo compose: core, `cv`, `model`
- [ ] Before the first translated message: Go PO library with plural rules; later, a translation service for community translations
