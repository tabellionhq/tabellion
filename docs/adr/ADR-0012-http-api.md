# ADR-0012: HTTP API

- Status: Accepted
- Date: 2026-10-18
- Drafted by: owner and agent (Claude Code, claude-opus-5-5)

## Context

APR-0002 asks for a fast bootstrap, easy plug-in with agents, and complexity discoverable later. Clients include the command line, agents (MCP) and other systems reading views (US-07).

## Options

| Option | Pros | Cons |
|---|---|---|
| A. Library and command line, no API | Simplest | Every client re-implements access; no remote |
| B. HTTP API, long-lived service | One interface; remote access | A service to install and run |
| C. HTTP API, auto-spawned server | One interface; nothing to manage; one writer | Start-up latency on the first call |

## Decision

C. The software API is all HTTP, like Docker and its daemon. One binary, `tabellion`.

- **Client**: turns commands into HTTP requests and manages auth. Subcommands for people (`tabellion practice add …`); a raw escape hatch for power users and agents (`tabellion api POST /practices -d @fact.yaml`).
- **Server**: always an HTTP endpoint.
  - Default: auto-spawned on the first request, on a Unix socket owned by the user; exits after 5 minutes idle.
  - `tabellion serve [--detach]`: TCP, to share with others; encryption (TLS) enforced; third-party auth adapters (ADR-0013).
- **One writer**: the running server owns the store; intensive writes (CV import) and request bursts (MCP searches) go through it.
- **No configuration required**: smart defaults (APR-0002); configuration only overrides them.
- Agents connect through an MCP adapter, itself an HTTP client.

## Reason

In the direct line of APR-0002: like Docker, one HTTP API for every client. Auto-spawn detaching for a time is the best of both worlds: it fits our needs, as importing a CV is intensive writing and searching from MCP can require multiple requests.

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [ADR-0011: View freshness](ADR-0011-view-freshness.md)
