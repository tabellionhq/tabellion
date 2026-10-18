# ADR-0016: GUI in the main binary

- Status: Accepted
- Date: 2026-10-18
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

The server exposes a limited set of core features. Adapters make them useful: mainly inputs (CV import) and outputs (CV export). Without a GUI to unite them, adapters are pretty useless to a person. Anne edits her views, reviews proposed facts and accepts changes; the GUI may later be augmented by an agent. APR-0002 asks for one install, nothing to configure, and a fast bootstrap.

## Options

| Option | Pros | Cons |
|---|---|---|
| A. Web GUI served by `skills-track`, opened in the browser | One binary; no install beyond it; the GUI is a client of the same HTTP API | The GUI needs a local TCP port or a browser-reachable address, since browsers do not speak Unix sockets |
| B. Desktop application, separate | Native feel | A second package to install and update; contradicts APR-0002 |
| C. No GUI; third parties build it | Smallest scope | Anne must use the command line or an agent first; slow adoption |

## Recommendation

A: `skills-track ui` starts the server if needed and opens the browser. The GUI uses the public HTTP API only, with no private endpoint. It listens on localhost only, with a one-time token in the URL it opens.

## Decision

A: we provide a core web UI in the main binary. It exposes the adapters (inputs such as CV import, outputs such as CV export) and the core features, such as reviewing facts.

## Reason

The server exposes a limited set of core features, made useful by adapters. Without a GUI to unite them, the adapters are pretty useless.

## Consequences

- The web UI is a client of the public HTTP API, driven by its actions (ADR-0015): an adapter the server offers appears in the UI as an action, with no UI change.
- How the UI authenticates is still open.

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [ADR-0012: HTTP API](ADR-0012-http-api.md)
- [ADR-0013: Access control](ADR-0013-access-control.md)
- [ADR-0015: API maturity, Richardson level 3](ADR-0015-api-maturity-richardson-level-3.md)
