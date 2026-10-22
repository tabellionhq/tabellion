# Test data

Test records, each with the document it comes from. The facts are written first; the document is written from them.

| Directory | Person | Stories |
|---|---|---|
| [anne/](anne/) | Anne, backend developer ([personas](../docs/design/personas.md)) | US-01 import from a CV, US-08 new CV, US-09 review |

## Files

- `facts.yaml`: the record after import and review. Each fact quotes the excerpt of the document that states it, and says how it reached that state (accepted, edited, split, expanded, resolved, completed).
- `cv.txt`: the CV as plain text, as saved from a word processor.

## Rules

- Every excerpt in `facts.yaml` appears verbatim in the document.
- Every line of the document is a fact or a term not recorded.
- ESCO and Wikidata identifiers are real, checked against their APIs; the date is in `facts.yaml`.
- People, companies and products are fictional; emails use `example.org`.

## Uses

- Extraction: the import of `cv.txt` is compared with the facts and their excerpts.
- End-to-end: the MVP session ([mvp-session.md](../docs/design/mvp-session.md)) replayed on Anne's data; the view gives the expected facts.
