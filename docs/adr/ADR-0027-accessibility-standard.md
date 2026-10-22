# ADR-0027: Accessibility standard

- Status: Accepted
- Date: 2026-10-22
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

Accessibility is mandatory (APR-0006). The web UI is HTML rendered by the server, with HTMX swapping parts of pages (ADR-0021); the command line prints text for people and structured output for programs. Users are anywhere in the world (APR-0004); the hosted service may be offered in the European Union (APR-0005).

Reference texts:

- WCAG 2.2, W3C Recommendation: the international guidelines for web content, levels A, AA, AAA.
- EN 301 549 V3.2.1: the European standard for public procurement of ICT; for the web, it requires WCAG 2.1 level AA.
- RGAA 4.1: the French method to check WCAG conformance, for public bodies.
- European Accessibility Act (Directive (EU) 2019/882): accessibility of some products and services sold in the EU, applicable since 28 June 2025. Whether it applies to a hosted Tabellion service: not verified.

## Options

| Option | Target | Issues |
|---|---|---|
| A. WCAG 2.2 level AA | Web UI; the command line follows the applicable criteria | The usual target of laws and procurement; covers EN 301 549 for the web, which requires WCAG 2.1 AA |
| B. WCAG 2.2 level AAA | All criteria | The W3C does not recommend AAA as a general policy: some criteria cannot be met for all content |
| C. EN 301 549 or RGAA | A legal framework | Tied to one region; both build on WCAG 2.1 AA |

## Recommendation

A, WCAG 2.2 level AA, for the web UI, the documentation site and the command line where the criteria apply. AAA criteria are met where they cost little, e.g. no time limits.

## Decision

A: WCAG 2.2 level AA.

## Reason

It is the legal standard.

## Consequences

- Wireframes show, for each screen: the heading structure, the keyboard order, the focus after each action, the label of every control, and what a screen reader announces.
- HTMX swaps: after a swap, focus moves to the changed part, or a live region announces the change; a status is never shown by colour alone.
- Every page sets its language (`lang`) from the language served (ADR-0019).
- Contrast at least 4.5:1 for text; targets at least 24 by 24 CSS pixels; usable at 400 % zoom and 320 CSS pixels wide.
- Command line: no information by colour alone; honours `NO_COLOR`; structured output available for assistive tools.
- Checks: automated checks of every page in CI (in Docker); a manual pass with a keyboard and a screen reader before each minor release.
- A conformance statement published with each release.

## References

- [APR-0004: International](../apr/APR-0004-international.md)
- [APR-0005: Hosted service possible](../apr/APR-0005-hosted-service-possible.md)
- [APR-0006: Accessibility](../apr/APR-0006-accessibility.md)
- [ADR-0019: Internationalisation](ADR-0019-internationalisation.md)
- [ADR-0021: Frontend framework](ADR-0021-frontend-framework.md)
- WCAG 2.2: <https://www.w3.org/TR/WCAG22/>
- EN 301 549 V3.2.1: <https://www.etsi.org/deliver/etsi_en/301500_301599/301549/03.02.01_60/en_301549v030201p.pdf>
- RGAA: <https://accessibilite.numerique.gouv.fr/>
- European Accessibility Act, Directive (EU) 2019/882: <https://eur-lex.europa.eu/eli/dir/2019/882/oj>
- NO_COLOR: <https://no-color.org/>
