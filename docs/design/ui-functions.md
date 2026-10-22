# UI functions

- Status: Draft
- Stories: US-01, US-08, US-09
- Contract: [api/openapi.yaml](../../api/openapi.yaml); session: [mvp-session.md](mvp-session.md)

The web UI is a client of the HTTP API (ADR-0016). What a person can do is decided by the server: the UI shows the actions each response carries. The UI has functions of its own only where the server cannot act: access, presentation, sequences of actions, language.

The web UI is online only. Apps that work offline are separate clients (US-10).

## Functions of the UI

| ID | The person… | How | Stories | MVP |
|---|---|---|---|---|
| U01 | opens the UI from the command line | one-time token swapped for a session cookie (ADR-0016) | US-08 | yes |
| U02 | signs in to an exposed or hosted instance | auth adapter (ADR-0013) | US-10 | |
| U03 | sees what to do next | the actions of `/`, with the number of facts waiting for review | US-08 | yes |
| U04 | goes through the review queue | proposed facts grouped by occupation, each with its excerpt and completion | US-09 | yes |
| U05 | splits a fact | edit the proposed fact, then add facts with the same evidence | US-09 | yes |
| U06 | expands a fact | add facts with the same evidence | US-09 | yes |
| U07 | accepts all the facts of an occupation | one accept per fact | US-09 | |
| U08 | abandons a term not resolved | hidden from the import summary; nothing sent to the server | US-09 | yes |
| U09 | writes a view while seeing what it shows | the view form next to its facts | US-08 | yes |
| U10 | picks the language | `Accept-Language` on every request (ADR-0019) | | |

## Server actions shown in the UI

From the contract (`x-actions`). The UI renders them; their rules are the server's.

| Action | On | Stories | MVP |
|---|---|---|---|
| Import a CV; follow the task; review what it proposed | entry point, task | US-01 | yes |
| Reject all facts proposed by a failed task | task | US-09 | |
| Edit, accept, reject an occupation or a practice | occupation, practice | US-09 | yes |
| Add an occupation; add a practice | record | US-01, US-09 | yes |
| Correct a fact; record progress as a new practice | accepted practice | US-09 | |
| Find an ESCO skill; find a Wikidata concept | entry point, search results | US-01, US-09 | yes |
| Add an identity | identities | US-08 | yes |
| Write a view; edit it; see its facts | views | US-08 | yes |
| Export a view as a CV (JSON Resume) | view | US-08 | yes |

## Open

- Deleting facts by a filter (ADR-0008): not in the contract yet; which filters.
- Exporting the whole record (APR-0001): not in the contract yet.
