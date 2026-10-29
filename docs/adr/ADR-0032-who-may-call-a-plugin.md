# ADR-0032: Who may call a plugin

- Status: Accepted
- Date: 2026-10-29
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

Plugins are services reached by URL (ADR-0028). The specification says what a plugin receives and returns (Section 7), not how it travels: this is the plugin HTTP binding, decided in four parts (ADR-0030 to ADR-0033). Already decided: a plugin declares its kinds and options and is called only for those (specification 4.10, 7.1); errors are problem details [RFC 9457] (7.1); TLS once a plugin is outside the host or private network (Security Considerations); a plugin keeps nothing it receives once it has answered (4.10).

### Problem

A plugin answers anyone who reaches its URL. Who may call it, and how does a plugin tell?

### Factors

- The demo compose puts the core and plugins on one private network (ADR-0028).
- A provider may run a plugin over the internet, sometimes at a cost per call.
- Simple adoption: no setup for the first steps (APR-0002).
- TLS is already required outside the host or private network (Security Considerations).

## Options

| Option | Issues | Impacts |
|---|---|---|
| F. Nothing on a private network; a bearer token, set on both sides, otherwise | Two settings to keep in step; a token to rotate | Zero setup for the demo compose; a provider can refuse unknown servers |
| G. Mutual TLS | Certificates to issue and renew on both sides | Strongest; heavy for self-hosting |
| H. Nothing | Anyone who reaches a plugin uses it, e.g. a paid provider's model | Only for localhost |

## Recommendation

F, always with TLS outside the host or private network.

## Decision

As recommended. F, always with TLS outside the host or private network.

## Reason

These options are simple and in line with our principles and previous ADRs.

## References

- [ADR-0028: Core and plugins as services](ADR-0028-core-and-plugins-as-services.md)
- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- RFC 9457, Problem Details for HTTP APIs: <https://www.rfc-editor.org/rfc/rfc9457>
