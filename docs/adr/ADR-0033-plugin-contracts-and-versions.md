# ADR-0033: Plugin contracts and versions

- Status: Proposed
- Date: 2026-10-29
- Drafted by: agent (Claude Code, claude-opus-5-5)

## Context

Plugins are services reached by URL (ADR-0028). The specification says what a plugin receives and returns (Section 7), not how it travels: this is the plugin HTTP binding, decided in four parts (ADR-0030 to ADR-0033). Already decided: a plugin declares its kinds and options and is called only for those (specification 4.10, 7.1); errors are problem details [RFC 9457] (7.1); TLS once a plugin is outside the host or private network (Security Considerations); a plugin keeps nothing it receives once it has answered (4.10).

### Problem

Where is each plugin contract written as an HTTP interface, and how does it change without breaking servers and plugins already deployed?

### Factors

- Contract first, code generated (ADR-0023); plugins are written in other languages (TypeScript, Rust), which also generate from OpenAPI.
- Servers and plugins are upgraded separately, by different people.
- The description already gives each kind's address (ADR-0030).
- The specification prevails over the contract (ADR-0024).

## Options

| Option | Issues | Impacts |
|---|---|---|
| I. One OpenAPI 3.1 file per kind, in `api/plugins/`, published with the core's contract; a kind name carries its major version (`term-matching/1`); the description names the specification revision | A file per kind to keep in step with Section 7 | Code generated on both sides; a breaking change is a new kind name, so old and new plugins coexist |
| J. Version in the path (`/v1/term-matching`) | Couples the address to the version | Familiar; the description already gives addresses, so it adds nothing |

## Recommendation

I.

## Decision

Owner's words.

## Reason

Owner's words.

## References

- [ADR-0028: Core and plugins as services](ADR-0028-core-and-plugins-as-services.md)
- [ADR-0023: HTTP stack](ADR-0023-http-stack.md)
- [ADR-0024: The specification is the source of truth](ADR-0024-specification-source-of-truth.md)
- RFC 9457, Problem Details for HTTP APIs: <https://www.rfc-editor.org/rfc/rfc9457>
