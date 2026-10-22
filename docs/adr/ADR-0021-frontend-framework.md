# ADR-0021: Frontend framework of the web UI

- Status: Accepted
- Date: 2026-10-21
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

The web UI is part of the main binary and a client of the public HTTP API, driven by its actions (ADR-0015, ADR-0016). An action the server offers appears in the UI without a UI change.

The MVP screens: home, CV import and its progress, review of proposed facts, identity, view form with a preview of the selected facts, export. The review screen is the most interactive: accept or reject, edit, split, expand, accept a group, resolve a term (US-09).

Constraints:

- One binary on Windows, macOS and Linux (ADR-0020); UI files embedded in it.
- Easy maintenance; code that people and agents already know well.
- Few dependencies (APR-0002); the server is in Go (ADR-0014).
- Messages translated (APR-0004, ADR-0019).
- Online only. Apps that work offline, e.g. on a phone, are separate clients of the API (US-10).

## Options

| Option | How it works | Pros | Cons |
|---|---|---|---|
| A. Server-rendered HTML with HTMX | The server answers `Accept: text/html` with an HTML rendering of each resource (Go `html/template`); actions become links, forms and buttons; HTMX swaps page parts without reloading | One language (Go); no JavaScript build in CI; HTML is hypermedia, so actions render with one generic template; messages translated on the server, same catalogue as the API; one small script, no package manager | Rich interactions (split, expand, drag) need small hand-written scripts; fewer ready-made components; HTML responses are a second rendering of every resource to test |
| B. Single-page app, React and TypeScript | A JavaScript app reads the JSON-LD responses and renders actions generically; built files embedded in the binary | The most common frontend stack: many components, many developers, well known to coding agents; rich interactions are easy | Node.js toolchain and npm dependencies in CI and in every contributor's setup; two languages; a build step before each Go build; translations in a second catalogue unless shared |
| C. Single-page app, Svelte or Vue | As B, with a lighter framework | Smaller bundles than React; simpler code for small apps | As B; fewer developers and components than React |
| D. Web components, no build step | Plain JavaScript modules (with or without Lit) served as files, reading the JSON-LD responses | No bundler; standard browser APIs; works for years without upgrades | Fewer ready-made components; state and routing written by hand |
| E. Angular | Full framework, TypeScript | Complete: forms, routing, i18n built in | Heaviest option; most code for a small UI |

## Recommendation

A, HTMX with server-rendered HTML. The UI functions ([ui-functions.md](../design/ui-functions.md)) are mostly server actions; its own functions are access, presentation, sequences of actions and language.

- It follows the API: an action is a link or a form, rendered by one template, so a new action needs no UI change (ADR-0016).
- The binary and CI stay Go only (ADR-0020).
- One message catalogue for the API and the UI (ADR-0019).

Risk: the review screen. Check it on the wireframes first; if split and expand need more than small scripts, D for that screen only, or B.

Not verified: HTMX licence and current version; Go libraries for content negotiation.

## Decision

A: HTMX, with HTML rendered by the server under `/ui`, next to the JSON-LD API under `/api`.

## Reason

Returning web pages seems ugly, but in a hexagonal approach it is clean: `/api` is a JSON format driven by the OpenAPI contract, `/ui` an HTML format. Both are adapters of the same core.

## Consequences

- Two driving adapters on the same core: `/api` (JSON-LD, the OpenAPI contract) and `/ui` (HTML). Neither holds rules; both show the actions the core makes available.
- The `/ui` adapter calls the core directly, not the HTTP API. No feature exists only in the UI.
- The review screen's split and expand are written as small scripts; to check on the wireframes.

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [APR-0004: International](../apr/APR-0004-international.md)
- [ADR-0014: Language](ADR-0014-language.md)
- [ADR-0015: API maturity, Richardson level 3](ADR-0015-api-maturity-richardson-level-3.md)
- [ADR-0016: GUI in the main binary](ADR-0016-gui-in-the-binary.md)
- [ADR-0019: Internationalisation](ADR-0019-internationalisation.md)
- [ADR-0020: Target platforms](ADR-0020-target-platforms.md)
- [US-09: Review imported facts](../design/user-stories/US-09-review-imported-facts.md)
- [US-10: Install an app, run no server](../design/user-stories/US-10-hosted-mobile.md)
- HTMX: <https://htmx.org>
