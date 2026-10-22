# ADR-0020: Target platforms

- Status: Accepted
- Date: 2026-10-21
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

The target is to compile `tabellion` for Windows, macOS and Linux (APR-0004). The binary is a client and an auto-spawned server on a socket owned by the user (ADR-0012), with a web UI (ADR-0016). It needs a place for the socket, one for the data, and builds for each system.

Unix domain sockets are available on Windows since Windows 10 version 1803, and Go supports them on all three systems.

## Options

| Option | Platforms | Pros | Cons |
|---|---|---|---|
| A. Linux first | Linux; others later | Smallest effort | Most users run Windows or macOS; paths and builds designed late |
| B. Linux and macOS | Unix systems | Same socket and path model | Excludes Windows users |
| C. Windows, macOS and Linux from the start | all three, on amd64 and arm64 | Matches the target; platform issues found early | Builds and tests on three systems |

## Recommendation

C.

| System | Socket | Data |
|---|---|---|
| Linux | `$XDG_RUNTIME_DIR/tabellion/` | `$XDG_DATA_HOME/tabellion/`, default `~/.local/share/tabellion/` |
| macOS | `$TMPDIR/tabellion/` (per user) | `~/Library/Application Support/tabellion/` |
| Windows | `%LOCALAPPDATA%\tabellion\` | `%LOCALAPPDATA%\tabellion\` |

- Architectures: amd64 and arm64.
- Builds: Go cross-compilation, without C dependencies (`CGO_ENABLED=0`) as long as possible, so that one machine builds every target.
- Native components that cannot avoid C, such as a local language model or an embeddings engine, run as separate processes, so the main binary stays cross-compiled.
- Tests: end-to-end tests in Docker cover Linux; continuous integration builds every target and runs the tests on Windows and macOS runners.

## Decision

C.

## Reason

Our target is to compile for Windows, macOS and Linux.

## Consequences

- The choices of store engine, local language model and embeddings engine must run on the three systems.
- The directory name is the project name, `tabellion`.
- The CI chain must be robust, and may prove costly. Local development compiles for the developer's own system, probably Linux. The full build, for every target, runs only on a full release: a push to main, which sets a new semantic version and publishes a release.

## References

- [APR-0004: International](../apr/APR-0004-international.md)
- [ADR-0012: HTTP API](ADR-0012-http-api.md)
- [ADR-0014: Language](ADR-0014-language.md)
- [ADR-0016: Core web UI in the main binary](ADR-0016-gui-in-the-binary.md)
- AF_UNIX on Windows: <https://devblogs.microsoft.com/commandline/af_unix-comes-to-windows/>
- XDG Base Directory Specification: <https://specifications.freedesktop.org/basedir-spec/latest/>
