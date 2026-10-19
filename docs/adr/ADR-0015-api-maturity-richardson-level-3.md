# ADR-0015: API maturity, Richardson level 3

- Status: Accepted
- Date: 2026-10-18
- Drafted by: agent (Claude Code, claude-opus-5-5); target level set by owner

## Context

The HTTP API is the contract for every client: command line, GUI, other systems and agents (ADR-0012). The MCP adapter may not last, so agents must be able to use the HTTP API directly. The target is Richardson maturity level 3: responses carry the links and actions available next, and clients follow them instead of building URLs.

For agents, links are not enough: each action needs its method, target, expected input and a title in plain words. Our data already references ESCO and Wikidata, which publish linked data (RDF).

## Options

| Option | Status | Pros | Cons |
|---|---|---|---|
| A. HAL | Expired IETF draft | Smallest; widely known | Links only: no method, no input; agents must guess |
| B. HAL + HAL-FORMS | Community specifications | Adds templates (method, fields) to HAL | Two specifications, small community |
| C. Siren | One author's specification; IANA media type `application/vnd.siren+json`, last updated 2012 | Actions built in: method, fields, title | No standards body; limited adoption; we would carry it alone |
| D. JSON-LD 1.1 + schema.org actions | JSON-LD is a W3C Recommendation; schema.org is maintained | Robust standard; ESCO and Wikidata URIs become linked data directly; `potentialAction` with `EntryPoint` gives method, URL template, content type; compacted JSON-LD reads as plain JSON | Heavier to learn; schema.org actions are mostly used for search engines, not APIs |
| E. JSON-LD 1.1 + Hydra | Hydra: W3C community group draft, not finished since 2020 | Designed for hypermedia APIs: operations, expected input, collections | Draft status; small community |

## Recommendation

D: JSON-LD 1.1 as the format, schema.org actions for what can be done next. Hydra (E) stays available for what schema.org lacks, such as collections and paging.

## Decision

D: Richardson level 3, oriented for agentic use. JSON-LD 1.1 as the format, schema.org actions for what can be done next; Hydra for what schema.org lacks.

## Reason

Richardson level 3 needs to be oriented for agentic use. JSON-LD is more robust than Siren. MCP is not guaranteed to stay: we want a robust API and not depend 100% on MCP.

## Out of scope

- What an agent can or cannot do with the API: a moving target, and a good prompt can teach it.

## References

- [APR-0003: Skills curriculum standard, existing standards for the rest](../apr/APR-0003-curriculum-standard-defer-the-rest.md)
- [ADR-0012: HTTP API](ADR-0012-http-api.md)
- [Resources](../design/resources.md)
- JSON-LD 1.1: <https://www.w3.org/TR/json-ld11/>
- schema.org actions: <https://schema.org/docs/actions.html>, <https://schema.org/EntryPoint>
- Hydra: <https://www.hydra-cg.com/spec/latest/core/>
- Siren media type: <https://www.iana.org/assignments/media-types/application/vnd.siren+json>
