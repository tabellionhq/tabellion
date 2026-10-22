# ADR-0019: Internationalisation of UI and command line messages

- Status: Accepted
- Date: 2026-10-22
- Drafted by: agent (Claude Code, claude-opus-5-5); need, decision and reason by owner

## Context

The software is written in English; its users are anywhere in the world (APR-0004). Reference data is already multilingual: ESCO and Wikidata labels exist in many languages. Messages written by `tabellion` are not: web UI texts, command line output, error details, action names in API responses.

APR-0003: an existing standard for messages and translations, rather than our own format.

## Options

| Option | Format | Pros | Cons |
|---|---|---|---|
| A. GNU gettext | PO files | The most widespread translation format; supported by translation platforms (Weblate, Transifex) and tools | Plural rules per language in each file; message keys are the English text |
| B. Unicode MessageFormat 2.0 | MF2 messages, CLDR plural and gender rules | Unicode standard; handles plurals, genders, selections | Recent; Go support not verified |
| C. go-i18n | TOML, JSON or YAML message files, CLDR plural rules | Go library in common use; message IDs separate from text | A library format, not a standard |
| D. XLIFF 2 | XML exchange format | OASIS standard for exchange with translators | Verbose; used as an exchange format, rarely at runtime |

## Recommendation

A: gettext PO files, the format translators and platforms already use. English is the source language and the fallback.

Language selection, by order:

- Web UI and API: the `Accept-Language` request header; action names in `potentialAction` and problem details are translated.
- Command line: `LC_ALL`, `LC_MESSAGES`, `LANG`.

Not verified: the state of Go libraries reading PO files.

## Decision

A: GNU gettext PO files.

## Reason

A mature solution, which we can share between applications.

## Consequences

- Messages of the web UI, the command line, action names and problem details are in PO files; English is the source and the fallback.
- The documentation is not translated through gettext; its translation is a separate question.
- Community translations may need a translation service working with PO files (e.g. Weblate); to choose when translators come.
- To check before the first translated message: a Go library reading PO files, with plural rules.

## References

- [APR-0003: Skills curriculum standard, existing standards for the rest](../apr/APR-0003-curriculum-standard-defer-the-rest.md)
- [APR-0004: International](../apr/APR-0004-international.md)
- GNU gettext PO format: <https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html>
- Unicode MessageFormat: <https://unicode.org/reports/tr35/tr35-messageFormat.html>
- RFC 9110, `Accept-Language`: <https://www.rfc-editor.org/rfc/rfc9110#name-accept-language>
