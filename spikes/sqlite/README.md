# SQLite engine checks

Checks what ADR-0022 requires of `ncruces/go-sqlite3` before relying on it. A throwaway program, not product code.

```shell
docker run --rm -v "$PWD:/w" -w /w golang:1.26 sh -c 'CGO_ENABLED=0 go build -o /tmp/sqlite . && /tmp/sqlite checks /tmp/run'
```

## Results

`ncruces/go-sqlite3` v0.35.6, Go 1.26, `CGO_ENABLED=0`, in Docker on Linux amd64, 2026-10-28.

| Check | Result |
|---|---|
| Encryption, service key (Adiantum, 32-byte key) | file holds no plaintext; right key reads 1000 rows; wrong key refused ("file is not a database") |
| Encryption, person passphrase (Argon2id key derivation) | same; opening takes about 60 ms, the cost of key derivation |
| Full-text search (FTS5), encrypted | 41,701 labels (ESCO size, 3 languages, made up) indexed in 229 ms; a word found in under 1 ms |
| Vector search (vec1), flat index, cosine | 20,000 made-up vectors of 384 dimensions loaded in 121 ms; 5 nearest in 12 ms; the expected vector first |
| Start: open an encrypted database and read it, new process | 4 to 6 ms; 7.4 MiB peak memory |
| Peak memory, all checks in one process | 151 MiB, including Argon2id's 64 MiB |
| Builds, stripped | linux, darwin, windows on amd64 and arm64: all build; 9 to 10 MiB |
| Writes, one transaction per row | about 3.8 ms per row (1000 rows in 3.8 s) |

## Limits

- Vectors are random: the check shows the search works, not the quality of real embeddings.
- vec1 in ANN mode (a trained index, for more than a few thousand vectors) is not checked.
- Labels are made up, at the size of ESCO; real ESCO data is not loaded.
- Measured on one machine, Linux only; memory is measured on Linux only.
