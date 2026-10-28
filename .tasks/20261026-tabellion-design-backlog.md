# Task: Tabellion design backlog

- Date: 2026-10-26
- Model: Claude Opus 5.5 (claude-opus-5-5)
- Tool: Claude Code
- Status: Open
- From: [20261014-poc-knowledge-validation](20261014-poc-knowledge-validation.md), open items deferred

## Owner decisions

- [>] Language model for CV extraction: CV roadmap
- [ ] Concept search in the MVP: Wikidata online search or local embeddings
- [>] CV prose: CV roadmap
- [x] Level names: level 3 renamed Experienced (ADR-0002, specification 4.3)
- [x] Date precision: year, year-month or full date, as the source gives (specification 3, 4.5; contract `PartialDate`)
- [x] Rule: no telemetry (specification 11)
- [x] Consent before data leaves the local system: moved to ROADMAP ("Where the data goes")
- [x] Create GitHub `tabellionhq`
- [-] Manual trademark check: dropped by the owner

## Specification and contract

- [>] Identity picked up from a CV: CV roadmap
- [ ] Plugin contracts: HTTP binding
- [ ] Preferences (3.10, 4.11): per-person entries, opaque to the specification, owner only, never exposed by views; e.g. `cv/1: {identity, view, theme, format}` for a saved CV; a plugin gets only the entry its call needs
- [>] CV themes: CV roadmap
- [x] Operation to leave out a term not matched: `leaveOutTerm` (specification 4.7)
- [ ] Home shows terms not matched

## Screens

- [x] Screens 03 reading and confirm, 06 export: saved
- [>] "My CVs" screen: CV roadmap
- [ ] Check screen: drag-and-drop library, and a way to do the same without dragging (WCAG 2.5.7)
- [x] Review screen wireframe: dropped, screen 04 covers the review

## Checks and setup

- [x] ncruces/go-sqlite3 checks: all pass (AgDR-0005, `spikes/sqlite/`)
- [-] GDPR: dropped by the owner
- [x] Licence files: `licences/`, linked from the README
- [x] DCO check in CI (AgDR-0004)
- [ ] Licence check of Go dependencies in CI, once there is Go code
- [ ] Demo compose: core and `model`; `cv` with the CV adapter (CV roadmap)
- [ ] Before the first translated message: Go PO library with plural rules; later, a translation service for community translations
