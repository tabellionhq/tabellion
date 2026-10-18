# ADR-0016: GUI in the main binary

- Status: Proposed
- Date: 2026-10-18
- Drafted by: agent (Claude Code, claude-opus-5-5); scope set by owner

## Context

Anne edits her views, reviews proposed facts and accepts changes. She needs a graphical interface, which may later be augmented by an agent. APR-0002 asks for one install, nothing to configure, and a fast bootstrap.

## Options

| Option | Pros | Cons |
|---|---|---|
| A. Web GUI served by `skills-track`, opened in the browser | One binary; no install beyond it; the GUI is a client of the same HTTP API | The GUI needs a local TCP port or a browser-reachable address, since browsers do not speak Unix sockets |
| B. Desktop application, separate | Native feel | A second package to install and update; contradicts APR-0002 |
| C. No GUI; third parties build it | Smallest scope | Anne must use the command line or an agent first; slow adoption |

## Recommendation

A: `skills-track ui` starts the server if needed and opens the browser. The GUI uses the public HTTP API only, with no private endpoint. It listens on localhost only, with a one-time token in the URL it opens.

## Decision

Owner's words.

## Reason

Owner's words.

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [ADR-0012: HTTP API](ADR-0012-http-api.md)
- [ADR-0013: Access control](ADR-0013-access-control.md)
