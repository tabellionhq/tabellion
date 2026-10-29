# ADR-0030: Plugin self-description

- Status: Proposed
- Date: 2026-10-29
- Drafted by: agent (Claude Code, claude-opus-5-5)

## Context

Plugins are services reached by URL (ADR-0028). The specification says what a plugin receives and returns (Section 7), not how it travels: this is the plugin HTTP binding, decided in four parts (ADR-0030 to ADR-0033). Already decided: a plugin declares its kinds and options and is called only for those (specification 4.10, 7.1); errors are problem details [RFC 9457] (7.1); TLS once a plugin is outside the host or private network (Security Considerations); a plugin keeps nothing it receives once it has answered (4.10).

### Problem

How does a server learn what a plugin does: its kinds, their options, and where to call each one?

### Factors

- One URL to configure per plugin, in the server's settings or the compose file.
- Reverse proxies often serve a service under a path prefix: a plugin's base URL may be `https://example.org/plugins/model/`.
- The core's API is hypermedia: an entry point lists the actions (ADR-0015).
- A plugin's options, e.g. the languages it matches, change what the server offers the owner (4.10: no action whose plugin cannot do it).

## Options

| Option | Issues | Impacts |
|---|---|---|
| A. `GET` on the base URL returns the description, with an action per kind: method, target, media types | none known | Same pattern as the core's entry point; works under a path prefix |
| B. `GET /.well-known/tabellion-plugin` | Well-known URIs sit at the host root (RFC 8615): breaks under a path prefix; a public name is registered with IANA | Conventional for discovery across hosts |

## Recommendation

A: the description at the base URL, with an action per kind.

## Decision

Owner's words.

## Reason

Owner's words.

## References

- [ADR-0028: Core and plugins as services](ADR-0028-core-and-plugins-as-services.md)
- [ADR-0015: API maturity, Richardson level 3](ADR-0015-api-maturity-richardson-level-3.md)
- RFC 8615, Well-Known URIs: <https://www.rfc-editor.org/rfc/rfc8615>
- RFC 9457, Problem Details for HTTP APIs: <https://www.rfc-editor.org/rfc/rfc9457>
