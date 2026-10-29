# ADR-0034: Go library for PO files

- Status: Accepted
- Date: 2026-10-29
- Drafted by: agent (Claude Code, claude-opus-5-5); decision by owner

## Context

### Problem

ADR-0019 chose GNU gettext PO files for the messages of the web UI, the command line, action names and problem details, with English as the source and the fallback. It left one check for later: a Go library that reads PO files, with plural rules. The library is compiled into the server.

### Factors

- Plural rules differ by language: French has 2 forms, Polish 3, Arabic 6. The `Plural-Forms` header of each PO file carries the rule.
- A dependency compiled into the server: licence compatible with AGPL-3.0 (ADR-0025), no C (ADR-0020), few transitive modules.
- The server picks a language from the request (`Accept-Language`) and falls back to English.
- Messages are extracted from the Go code into a template (POT) for translators.
- `golang.org/x/text/message` does not read PO files.

### Test

French and Polish PO files, plural rules from their headers, `n` = 0, 1, 2, 5, 22, 25, 112; Go 1.27.3, `CGO_ENABLED=0`, in Docker, 2026-10-29. All three candidates return the right form for every `n` in both languages.

## Options

| Option | Licence | Version, release | Dependencies | Issues | Impacts |
|---|---|---|---|---|---|
| A. `github.com/leonelquinteros/gotext` | MIT | v1.8.0, 2026-09-22 | none | No language negotiation: done with `golang.org/x/text/language` | Extraction tool `xgotext`; PO and MO; domains |
| B. `github.com/vorlif/spreak` | MIT | v1.0.0, 2025-09-23 | `golang.org/x/text` | One maintainer; one release in the past year | Language matching built in; CLDR plural rules as a fallback when a file has no `Plural-Forms`; extraction tool `xspreak` |
| C. `github.com/chai2010/gettext-go` | BSD-3-Clause | v1.0.3, 2024-04-10 | none | No release for two and a half years | Mirrors the C gettext API |

## Recommendation

A, `gotext`: no dependency, the most recent release, an extraction tool. Language negotiation with `golang.org/x/text/language`, behind the same i18n port, so that B stays a drop-in alternative.

## Decision

OK for gotext (A).

## Reason

The recommendation, accepted by the owner: no dependency, the most recent release, an extraction tool.

## Next steps

- The i18n port: translate a message, with a plural, for a language picked from the request.
- `xgotext` in the development image (ROADMAP) and in CI, to keep the template in step with the code.

## References

- [ADR-0019: Internationalisation](ADR-0019-internationalisation.md)
- [ADR-0020: Target platforms](ADR-0020-target-platforms.md)
- [ADR-0025: Licences](ADR-0025-licences.md)
- gotext: <https://github.com/leonelquinteros/gotext>
- spreak: <https://github.com/vorlif/spreak>
- gettext-go: <https://github.com/chai2010/gettext-go>
- GNU gettext, plural forms: <https://www.gnu.org/software/gettext/manual/html_node/Plural-forms.html>
