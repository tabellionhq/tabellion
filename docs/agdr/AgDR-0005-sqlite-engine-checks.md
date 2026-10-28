---
id: AgDR-0005
timestamp: 2026-10-28T20:28:25Z
agent: Claude Code
model: claude-opus-5-5
trigger: ADR-0022 consequences, "to check before relying on them"; owner request
status: executed
---

# AgDR-0005: SQLite engine checks

> In the context of ADR-0022 choosing `ncruces/go-sqlite3` with checks to run first, facing unmeasured encryption, search and start time, I decided to check them with a throwaway program in `spikes/sqlite/`, to confirm the engine before building on it, accepting made-up data and one machine.

## Options considered

- A: A throwaway program, one check per requirement, in Docker.
- B: Check while building the Store adapter: no throwaway code, but a failure would come after the adapter is written.

## Decision

A. Versions pinned: `github.com/ncruces/go-sqlite3` v0.35.6, with `go-sqlite3-wasm/v6` v6.3.35304, `lukechampine.com/adiantum` v1.1.1, `golang.org/x/crypto` v0.57.0, `golang.org/x/sys` v0.48.0, `ncruces/julianday` v1.0.0. Extensions: `vfs/adiantum`, `ext/fts5`, `ext/vec1`.

## Consequences

- Every check passes: results in [spikes/sqlite/README.md](../../spikes/sqlite/README.md). Start: 4 to 6 ms, 7.4 MiB; builds: 9 to 10 MiB on all six targets.
- v0.35.6 requires Go 1.26; ADR-0023 tested code generation with Go 1.25.
- FTS5 and vec1 are extensions registered on each connection, not built in.
- One write per transaction costs about 3.8 ms: imports write in one transaction.
- A passphrase costs about 60 ms and 64 MiB at each open (Argon2id).
- Not checked: vec1 with a trained index (ANN), real embeddings, real ESCO data, macOS and Windows at run time.

## Review

- Reviewer sub-agent: OK | Warning | Refused: (one line)
- Owner: kept | changed | pushed up to ADR-000x (date)
