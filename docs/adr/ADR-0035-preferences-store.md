# ADR-0035: Preferences store

- Status: Accepted
- Date: 2026-10-30
- Drafted by: agent (Claude Code, claude-opus-5-5); decision by owner

## Context

A person sets preferences: languages, time zone, accessibility needs (APR-0004, APR-0006). Plugins need some too, such as the CV filters and aliases a person used last time.

Constraints:

- The owner exports everything (APR-0001).
- A hosted service may store them, encrypted with a key held by the service or the person (APR-0005).
- A plugin keeps nothing it receives once it has answered (specification 4.10, ADR-0028).
- SQLite with `ncruces/go-sqlite3`, encryption checked (ADR-0022, ADR-0029).

Preferences are not facts: they change in place, carry no proof, and are not shared through views.

## Options

| Option | Issues | Impacts |
|---|---|---|
| A. Key-value tables in the record database | One file to back up and encrypt | Preferences and facts share one lifecycle; deleting the record deletes the preferences |
| B. Key-value tables in a separate preferences database | Two files, two keys or one key for both | Preferences survive a reset of the record, and the other way round; export covers both files |
| C. A configuration file (YAML, TOML) | No encryption without code of ours; concurrent writes | Easy to edit by hand; against APR-0005 for hosted instances |

In A and B:

- A main table holds the core's preferences; each plugin kind has its own table, named after the kind without its version (ADR-0033): `term-matching`, not `term-matching/1`.
- Tables are created by the server for the kinds it declares, never by a plugin.
- A key-value row: key (text), value (JSON), updated (datetime).
- The first development version is unencrypted; encryption comes with the hosted service, through a migration that rewrites the file.

How a plugin stores preferences without keeping data (4.10):

| Way | Issues |
|---|---|
| 1. The server sends the kind's preferences with each call; the plugin returns changes with its result; the server checks and stores them | Calls grow with the preferences; needs a size limit per value and per table |
| 2. The plugin calls the server's API to write them | The plugin needs a grant; against "never access to a record" (4.10) |

## Recommendation

B with way 1.

- Separate files, as the owner proposes: `Preferences.sqlite` and `Records.sqlite`.
- A size limit per value and per table keeps files out of preferences (no PDF). Limits to set.

Open point: paths. `~/.config` holds configuration on Linux (XDG); data usually goes in `~/.local/share` (`XDG_DATA_HOME`). Windows and macOS have their own folders. Records in `~/.config` would break that convention.

## Decision

Recommendation validated: B with way 1.

- One key for both files: easier.
- Export: a zip contains both files.
- Consent is in Records, in a table of its own.

## Reason

Owner's words.

## Next steps

- Paths: to decide.
- Specification: preferences section, plugin call carries the kind's preferences.
- OpenAPI: preferences resources.
- Store adapter: the two files, the tables, the limits.

## References

- [APR-0001: Free and open](../apr/APR-0001-free-and-open.md)
- [APR-0004: International](../apr/APR-0004-international.md)
- [APR-0005: Hosted service possible](../apr/APR-0005-hosted-service-possible.md)
- [APR-0006: Accessibility](../apr/APR-0006-accessibility.md)
- [ADR-0022: Store engine](ADR-0022-store-engine.md)
- [ADR-0028: Core and plugins as services](ADR-0028-core-and-plugins-as-services.md)
- [ADR-0029: Go toolchain and store dependencies](ADR-0029-go-toolchain-and-store-dependencies.md)
- [ADR-0033: Plugin contracts and versions](ADR-0033-plugin-contracts-and-versions.md)
- XDG Base Directory Specification: <https://specifications.freedesktop.org/basedir-spec/latest/>
