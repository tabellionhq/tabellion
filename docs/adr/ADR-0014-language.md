# ADR-0014: Language

- Status: Accepted
- Date: 2026-10-18
- Drafted by: agent (Claude Code, claude-opus-5-5) and owner

## Context

With ADR-0012, the client starts on every command and the server is auto-spawned: both must start very fast. The software ships as one binary, and running the server in containers is an enviable option.

## Options

| Option | Cold start | Distribution | HTTP over Unix socket | Notes |
|---|---|---|---|---|
| A. Go | milliseconds | static binary, simple cross-compilation; minimal container images | standard library | simple to read and review; SQLite needs cgo or a pure-Go port; local embeddings through bindings |
| B. Rust | milliseconds | static binary | needs crates (tokio, hyper or axum) | strong for local embeddings; steeper learning curve |
| C. Python or Node | tens to hundreds of milliseconds | needs a runtime | standard library | fast to prototype; heavy for a command called often |
| D. JVM | slow without ahead-of-time compilation | needs a runtime or GraalVM | standard library | ahead-of-time compilation adds complexity |

Heavy adapters (e.g. local embeddings) may run as separate processes in any language.

## Decision

A: Go.

## Reason

Go seems like the best option: running the server in Docker or containers is an enviable option, and Go is perfect for containers.

Adapters can be made with the best language for their job (e.g. Rust for local embeddings), as separate processes.

## References

- [ADR-0012: HTTP API](ADR-0012-http-api.md)
