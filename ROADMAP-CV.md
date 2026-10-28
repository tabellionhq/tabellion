# CV roadmap

The CV adapter is built after the core: the `cv` plugin and what CVs need from the rest of the system. Its contracts: [cv-adapter.md](docs/design/cv-adapter.md).

- CV round trip: import an old CV, export a fresh one dedicated to a job offer, through CV standards ([design note](docs/design/cv-standards.md)).
- Language model for CV extraction (ADR); a small local model to test first.
- Identity picked up from a CV: the import reads the name and contacts at the top of a CV and proposes an identity. A CV-adapter specification may come later.
- Align on JSON Resume, an existing standard (APR-0003); more parts may need fixing along the way ([cv-standards.md](docs/design/cv-standards.md)):
  - Occupation held gains its details: position, organisation, location, url, summary, highlights. The import copies them as written in the CV; a weak CV leaves them empty and completion shows it.
  - Identity gains the `basics` fields: headline, summary, location, links, several contacts.
- CV prose (summary, highlights): written by the person, generated, or both; where it is kept ([cv-standards.md](docs/design/cv-standards.md), Questions). A language model as an optional helper, e.g. drafting a summary from the practices; generated text is a draft the person accepts or edits, and never changes a fact.
- CV themes (npm): a few in the `cv` image, more installed on demand into a `cv-themes` volume, a rebuildable cache (specification 4.10). The CV rendering contract lists installed and available themes and gains an install operation. Risks: a theme is code run on the CV; installs reach npmjs.org (`--ignore-scripts`, pinned versions); theme licences vary.
- Saved CVs and the "My CVs" screen: identity, view, look and file saved together, built again with new facts ([screen 06](docs/design/screens/06-export.md)).
- Focus a CV on a job offer.
