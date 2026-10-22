# ADR-0022: Store engine

- Status: Proposed
- Date: 2026-10-22
- Drafted by: agent (Claude Code, claude-opus-5-5)

## Context

One server holds one person's record (ADR-0012): thousands of rows, mostly added, rarely deleted (ADR-0008). Reference data (ESCO skills and occupations, labels in many languages) and search indexes are rebuildable copies next to it.

Constraints:

- Built with `CGO_ENABLED=0` for Windows, macOS and Linux, on amd64 and arm64 (ADR-0020).
- No separate database server to install (APR-0002).
- Hosted instances wake on request and scale to zero; their storage may be encrypted, with the key held by the service or by the person (APR-0005). Start time matters.
- Search: full-text on ESCO labels; vectors for Wikidata concepts later (ADR-0018).

An embedded SQLite database meets the first two constraints: one file per person, no server. Two SQLite libraries for Go build without CGO.

## Options

| Option | Licence | Encryption at rest | Full-text, JSON, vectors | Maintenance | Issues |
|---|---|---|---|---|---|
| A. `ncruces/go-sqlite3` | MIT | Built in: Adiantum or AES-XTS file layer, key from the service or the person | FTS5, JSON, vector search | One main maintainer; before v1; releases about monthly | Higher memory use; slower on many rows and concurrent queries; encryption deterministic and without tamper protection, so backups need a signature |
| B. `modernc.org/sqlite` | BSD-3-Clause | None built in; the whole file encrypted by us before it leaves the instance | FTS5, JSON, vector search (sqlite-vec) | Two maintainers; releases about every two weeks; newest SQLite | 1.3 to 2 times the CPU time of SQLite in C; live file unencrypted |
| C. A database server (e.g. PostgreSQL) | PostgreSQL licence | Server or disk level | Full-text built in; vectors through an extension | Large community | A second program to install and run; against APR-0002; one server per person is heavy |

Not measured: start time and binary size for A and B. No published figures.

## Recommendation

A, `ncruces/go-sqlite3`, if a measurement confirms it: build a small binary with A and with B, compare start time, binary size and memory on the MVP data.

- Encryption at rest is available when the hosted service needs it (APR-0005), without code of ours.
- Both go through `database/sql` behind the Store port: switching to B later changes one adapter.

## Decision

Owner's words.

## Reason

Owner's words.

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [APR-0005: Hosted service possible](../apr/APR-0005-hosted-service-possible.md)
- [ADR-0008: Normalisation](ADR-0008-normalisation.md)
- [ADR-0012: HTTP API](ADR-0012-http-api.md)
- [ADR-0018: Concepts missing from Wikidata](ADR-0018-concepts-missing-from-wikidata.md)
- [ADR-0020: Target platforms](ADR-0020-target-platforms.md)
- ncruces/go-sqlite3: <https://github.com/ncruces/go-sqlite3>; encryption: <https://pkg.go.dev/github.com/ncruces/go-sqlite3/vfs/adiantum>
- modernc.org/sqlite: <https://pkg.go.dev/modernc.org/sqlite>
- Benchmark, March 2026: <https://github.com/cvilsmeier/go-sqlite-bench>
