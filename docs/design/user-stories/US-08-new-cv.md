# US-08: New CV filtered for an offer

- Status: Draft
- Catalogue: UC8 new CV
- Actor: Anne ([personas](../personas.md))
- Follows: [US-01 import from a CV](US-01-import-cv.md)

## Story

As Anne, I want to export a CV from my curriculum, filtered for a job offer, so that I send a targeted, up-to-date CV without retyping it.

## Flow, in the web UI

1. Import: Anne uploads her old CV, as a file (pdf, docx, odt, txt), a link or text (US-01).
2. Check: she completes, accepts or rejects the proposed facts (F04, US-09).
3. New CV: she picks the identity that signs it, then a view for the offer: occupations, periods, skill categories, facts to include or exclude (F06).
4. Export: as JSON Resume, or as PDF or HTML rendered with a JSON Resume theme (F07). The CV is saved, and rebuilt later with new facts.

## Data in, data out

- In: a CV; Anne's choice of identity and her picks for the view; a theme and a format.
- Out: a JSON Resume document, or a PDF or HTML file, holding only the facts the view shows.

## Functions required

F01 extract, F02 find the skill meant by words, F03 propose, F04 review, F06 write a view, F07 read through a view.

## Adapters required

| Port | Adapter |
|---|---|
| Client | web UI (ADR-0016) |
| Auth | UI one-time token |
| Store | one database engine |
| Extraction | plain text CV in the core; PDF, DOCX, ODT in the `cv` plugin |
| Term matching | `model` plugin: a small local language model to test; other providers as other plugins |
| Reference data | ESCO, Wikidata |
| Search | ESCO index in the core; embeddings on concepts in the `model` plugin |
| Export, rendering | JSON Resume by the core; PDF and HTML in the `cv` plugin |

## Rules

- The owner reads her own record through any of her views without a grant: grants are for other readers (ADR-0013).

## Open

- Which language model matches the terms of the CV: a small one running locally, to test; an ADR when it is chosen.
