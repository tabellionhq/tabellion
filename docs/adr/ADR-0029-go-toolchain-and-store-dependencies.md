# ADR-0029: Go toolchain and store dependencies

- Status: Proposed
- Date: 2026-10-28
- Drafted by: agent (Claude Code, claude-opus-5-5)

## Context

ADR-0022 chose `ncruces/go-sqlite3` and listed checks to run before relying on it. The checks pass ([spikes/sqlite](../../spikes/sqlite/README.md)): encryption with a service key and with a passphrase, full-text search, vector search, a start of 3 to 6 ms, builds for the six targets of ADR-0020.

Two things the checks pinned are escalation items:

- The toolchain. The current `ncruces/go-sqlite3`, v0.35.6, requires Go 1.26 or later. ADR-0023 tested code generation with Go 1.25. Go 1.27 is the current release; Go supports each release until two newer ones are out, so Go 1.25 is no longer supported.
- The dependencies compiled into the server for the store, and their licences.

| Module | Version | Licence | Used for |
|---|---|---|---|
| `github.com/ncruces/go-sqlite3` | v0.35.6 | MIT | SQLite driver, `database/sql` |
| `github.com/ncruces/go-sqlite3-wasm/v6` | v6.3.35304 | MIT No Attribution | SQLite and its extensions, compiled to Go |
| `github.com/ncruces/julianday` | v1.0.0 | MIT | dates |
| `lukechampine.com/adiantum` | v1.1.1 | MIT | encryption at rest |
| `golang.org/x/crypto` | v0.57.0 | BSD-3-Clause | Argon2id, key from a passphrase |
| `golang.org/x/sys` | v0.48.0 | BSD-3-Clause | system calls |

All are compatible with AGPL-3.0 (ADR-0025). Extensions used: `vfs/adiantum` (encryption), `ext/fts5` (full-text), `ext/vec1` (vectors). FTS5 and vec1 are registered on each connection.

vec1 is SQLite's own vector extension, version 0.8; its roadmap lists further work (graph indexes, more element types). ADR-0022 named sqlite-vec or vec1.

Checked on Go 1.27.3, 2026-10-28, in Docker:

| Check | Go 1.26 | Go 1.27 |
|---|---|---|
| Store checks (encryption, FTS5, vec1, start) | pass | pass |
| Server generated from `api/openapi.yaml` by `oapi-codegen` v2.8.1 (ADR-0023 options), `go vet`, `CGO_ENABLED=0` builds for linux, darwin, windows on amd64 and arm64 | not run | pass |

## Options

Toolchain:

| Option | Issues | Impacts |
|---|---|---|
| A. Go 1.27, the current release | Supported until Go 1.29, about August 2027; contributors need 1.27 | Current `go-sqlite3` and `oapi-codegen`; one upgrade a year to stay supported |
| B. Go 1.26, the oldest release the current `go-sqlite3` accepts | Supported until Go 1.28, about February 2027 | Wider range of contributor toolchains; an upgrade within months |
| C. Go 1.25, with an older `go-sqlite3` | Go 1.25 is unsupported: no security fixes; the last `go-sqlite3` for 1.25 is not identified, nor whether it has vec1 | Matches ADR-0023's test; older store code |

Upgrades:

| Option | Issues | Impacts |
|---|---|---|
| D. Follow Go releases: move to each new release within a set time, e.g. 3 months | Regular upgrade work | Always supported; dependencies may require new releases anyway |
| E. Upgrade when a dependency or a security fix requires it | Upgrades come unplanned, sometimes urgent | Less routine work |

Dependencies: pinned in `go.sum`; upgraded by a pull request that runs the checks and the licence check in CI (ADR-0025).

## Recommendation

A and D: Go 1.27, `go 1.27` in `go.mod`, the Docker build image pinned to the same release; move to each new Go release within 3 months. Keep the store dependencies above, with vec1 for vectors behind the Search port, so that sqlite-vec stays a fallback.

## Decision

Owner's words.

## Reason

Owner's words.

## Next steps

- `go.mod` of the server: `go 1.27`; CI and Docker images on `golang:1.27`.
- Licence check of Go dependencies in CI, with the server's first code (ADR-0025).
- ADR-0023: code generation checked on Go 1.27 with `oapi-codegen` v2.8.1.

## References

- [ADR-0020: Target platforms](ADR-0020-target-platforms.md)
- [ADR-0022: Store engine](ADR-0022-store-engine.md)
- [ADR-0023: HTTP stack](ADR-0023-http-stack.md)
- [ADR-0025: Licences](ADR-0025-licences.md)
- Go release policy: <https://go.dev/doc/devel/release#policy>
- vec1: <https://sqlite.org/vec1>
- ncruces/go-sqlite3: <https://github.com/ncruces/go-sqlite3>
