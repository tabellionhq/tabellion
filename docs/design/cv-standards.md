# CV standards

Turning an old CV into a fresh one, dedicated to a job offer, reusing existing CV standards ([APR-0003](../apr/APR-0003-curriculum-standard-defer-the-rest.md)).

## Flow

```text
old CV ──(US-01 import)──→ curriculum ──(US-06 view for the offer)──→ export adapter ──→ JSON Resume ──(theme)────────→ PDF, HTML
                                                                                   └──→ Europass    ──(EU editor)───→ PDF
```

Rendering stays with the standards' own tools ([ADR-0005](../adr/ADR-0005-scope-assessment-method.md)). The system provides the mapping from a view to a CV format, and from a CV format to proposed facts.

## Standards

| Standard | Publisher | Import | Export and rendering | Status |
|---|---|---|---|---|
| JSON Resume | community | yes, if the person has one | themes to HTML, PDF, Markdown, text; an ATS-friendliness check in recent CLI versions | convention without governance; MIT licence reported, not verified |
| Europass CV | European Commission | exact: a Europass PDF embeds its XML | the EU online editor produces the PDF | official in the EU; XML schema v3.0 (2020) and a JSON version; latest version not verified |
| HR Open Standards | HR Open Standards Consortium | Candidate schema, broader than a CV | for applicant tracking systems; a CV schema was announced as future work in 2021 | standards body; free download, licence not verified |

## Order

1. JSON Resume: reaches developers first; richest rendering.
2. Europass: EU recruiters; exact import from Europass PDFs.
3. HR Open Standards: if applicant tracking systems need it.

## Mapping

| Curriculum | JSON Resume |
|---|---|
| OccupationHeld, kind work or pleasure | `work[]`, `projects[]` |
| OccupationHeld, kind study | `education[]` |
| Practice (skill, concept, context, level) | `skills[]`: name, level, keywords |
| Practice of a human language | `languages[]` |
| Identity chosen for the view | `basics` (name, email) |

The mapping loses detail: skill, concept, context and level become a flat list of skills with keywords.

## JSON Resume compared with the model

Compared on 2026-10-22 with the JSON Resume schema and a published example, `registry.jsonresume.org/thomasdavis` (9 jobs, 9 projects, 6 skill groups, prose in every section).

| JSON Resume | Fields | Model | Gap |
|---|---|---|---|
| `basics` | name, label (a headline), image, email, phone, url, summary, location, profiles | Identity: label (a name), one contact | headline, summary, photo, location, links, several contacts; JSON Resume has no date of birth |
| `work` | name (employer), position, location, url, startDate, endDate, summary, highlights | Occupation held: one label, ESCO occupation, kind, from, to | employer and position are one label; no summary, no highlights |
| `education` | institution, area, studyType, score, courses | Occupation held, kind study | institution and degree are one label; no field, grade or courses |
| `projects`, `volunteer` | name or organization, position or roles, description, highlights, keywords, url | Occupation held, kind pleasure | which section a pleasure occupation goes to |
| `skills` | name (a group), level (text), keywords | Practices | fits: name from an ESCO skill group (ancestors), level from the level name, keywords from concepts and contexts |
| `languages` | language, fluency | Practice of an ESCO language skill | "native" is not a level |
| `awards`, `certificates`, `publications`, `references`, `interests` | | not modelled | certificates relate to signed evidence (US-05) |

What a reader reads is prose: `summary` and `highlights`. Facts hold no text of their own; the nearest is the CV excerpt kept as evidence. Most CVs do not come with good prose: the example is good because its prose is.

## Questions

1. Prose: written by the person, generated from the facts by a language model at export, or both (generated, then edited).
2. Where prose is kept: on facts (a description per occupation or practice), per view (text for one CV), or not kept (generated at each export).
3. The promise "update your CV for the last time" holds only if prose does not have to be rewritten by hand for each CV.
4. Occupation held: split the label into position and organisation (degree and institution for studies).
5. Identity: headline, summary, location, links, several contacts.
6. Pleasure occupations: projects or volunteer.
7. Languages: how to say "native".

## Gains

- A CV dedicated to an offer: the view selects facts related to the offer (F05), the adapter maps them, a theme renders them.
- A more reliable import: a Europass PDF or a JSON Resume gives structured data instead of free text to extract (F01).

## Sources

- Europass CV schema documentation v3.0: <https://europa.eu/europass/system/files/2020-08/ECV_Schema_Documentation_v3.0.0_20200602.pdf>
- Europass XML schema v3.0: <https://interoperable-europe.ec.europa.eu/collection/employment-and-working-conditions/solution/europass-xml-schema-v30>
- HR Open recruiting workgroup: <https://assets.noviams.com/novi-file-uploads/hropen/pdfs-and-documents/04_Recruiting_SM-dd7d0929.pdf>
- resume-cli 3.7.2: <https://cdn.jsdelivr.net/npm/resume-cli@3.7.2/README.md>
- JSON Resume schema guide: <https://docs.rxresu.me/guides/json-resume-schema>
- JSON Resume schema: <https://raw.githubusercontent.com/jsonresume/resume-schema/master/schema.json>
- JSON Resume example: <https://registry.jsonresume.org/thomasdavis.json>
