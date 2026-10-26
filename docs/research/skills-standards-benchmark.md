# Skills standards benchmark

- Date: 2026-10-14
- Scope: open standards to describe developer skills, levels and evidence, at the granularity of code concepts (e.g. "Rust: block value", "`?` operator").
- Constraint: [APR-0001](../apr/APR-0001-free-and-open.md), developer owns and exports everything, schema open.

## Bottom line

No single standard fits. Candidate model, by layer:

| Layer | Holds | Candidates |
|---|---|---|
| 1. Concepts | What a skill is, with a stable ID | Wikidata (fine), ESCO (anchor), own open concepts for gaps |
| 2. Framework | Concepts + levels + relations | CTDL-ASN or ELM |
| 3. Evidence | Quizzes, answers, verdicts | A separate knowledge-checks repository (quizzes, sessions) |
| 4. Claim | Signed, portable "I reached level Y on X" | W3C Verifiable Credentials 2.0, via Open Badges 3.0 or an ELM credential |

ELM and Open Badges 3.0 are both built on W3C Verifiable Credentials, so W3C VC is not a separate choice: it comes with either.

Two coherent stacks:

- **EU stack:** Wikidata + ESCO → ELM (framework and credential). One vocabulary end to end, free EU issuing infrastructure.
- **Global stack:** Wikidata + ESCO → CTDL-ASN (framework) → Open Badges 3.0 (credential). Wider adoption outside the EU; O\*NET already publishes in CTDL.

Not verified: how well ELM models a fine-grained competency framework compared with CTDL-ASN.

Avoid: SFIA, e-CF, Lightcast. Their licences forbid distribution or commercial use without a fee, which contradicts APR-0001.

## A. Skill taxonomies

| Candidate | Free / open | Granularity for code | Machine-readable | Fit |
|---|---|---|---|---|
| Wikidata | ✅ CC0 | ✅ concept level: Rust `Q575650`, exception handling `Q471748`, statement `Q613299`, exit status `Q2553232`; no levels | ✅ SPARQL, API | ✅ concept IDs |
| ESCO | ✅ free reuse with attribution | ❌ language level at best; no Rust concept (API search, 2026-10-14) | ✅ RDF, API | ✅ anchor only |
| O\*NET | ✅ CC BY 4.0 | ❌ occupation + tools | ✅ Excel, CTDL JSON-LD | ⚠️ US-centric, coarse |
| ACM/IEEE CS2023 | ❓ licence not found | ⚠️ knowledge units + learning outcomes + Bloom levels; language-agnostic | ⚠️ report, PDF | ⚠️ level model |
| DigComp 3.0 | ❓ licence not found | ❌ 21 competences, citizens | ✅ Excel, JSON-LD | ❌ |
| Lightcast Open Skills | ❌ free for nonprofits on approval, commercial licence otherwise | ⚠️ ~33,000 skills | ✅ API | ❌ |
| SFIA 9 | ❌ no distribution or credentials without fee | ❌ role level | ⚠️ | ❌ |
| e-CF (EN 16234-1) | ❌ paid standard | ❌ 41 role-level competences | ❌ | ❌ |

## B. Schemas

| Candidate | Free / open | Purpose | Fit |
|---|---|---|---|
| CTDL-ASN | ✅ CC BY 4.0 | competency frameworks in JSON-LD; used by O\*NET | ✅ framework |
| European Learning Model 3.x | ✅ CC BY 4.0 (one source) | EU learning and credential vocabulary, on W3C VC, links to ESCO | ✅ framework + credential, EU |
| 1EdTech CASE 1.1 | ⚠️ limited-permission specification | frameworks, rubrics; schools and higher education | ⚠️ |
| W3C VC 2.0 + Open Badges 3.0 | ✅ open specifications | signed claims held by the developer | ✅ claim |

## CASE 1.1, checked 2026-10-18

1EdTech Competencies and Academic Standards Exchange (CASE) 1.1 (<https://www.imsglobal.org/spec/case/v1p1>):

- Publishes competency frameworks (CFDocument, CFItem, CFAssociation, rubrics) through a REST/JSON API, providers to consumers.
- No learner records: frameworks only.
- Sectors: K-12, higher education, workforce.
- Use under a 1EdTech licence; certification for paid members only.

Fit: a framework format, like ESCO's role, not a curriculum format. Useful at most as an import adapter on the requester side (e.g. a training provider publishing its objectives in CASE). Not a candidate for the curriculum model or exchange (APR-0003).

## Open decisions

1. Stack: EU (ELM) or global (CTDL-ASN + Open Badges 3.0).
2. Level scale: revised Bloom (as in CS2023) or DigComp's 4 levels.
3. Licence of own concepts: CC0 or CC BY 4.0.
4. Whether quiz evidence is attached to the credential, given code snippets may be employer IP.

## Unconfirmed

- Licences of DigComp 3.0 and CS2023.
- ELM's CC BY 4.0 for the latest version.

## Sources

- ESCO licence FAQ: <https://esco.ec.europa.eu/fr/node/173>
- ESCO API: <https://esco.ec.europa.eu/ar/use-esco/use-esco-services-api>
- DigComp 3.0: <https://joint-research-centre.ec.europa.eu/scientific-activities/key-competences-lifelong-learning/digital-competence-framework-digcomp/digcomp-30_en>
- JRC144705: <https://publications.jrc.ec.europa.eu/repository/handle/JRC144705>
- O\*NET database: <https://onetcenter.org/database.html>
- O\*NET competency frameworks: <https://www.onetcenter.org/competencyFrameworks.html>
- Lightcast access options: <https://docs.lightcast.io/lightcast-api/docs/api-access>
- SFIA licensing: <https://sfia-online.org/en/licensing-sfia>
- DIN EN 16234-1: <https://dinmedia.de/en/standard/din-en-16234-1/317078638>
- CS2023 knowledge model: <https://csed.acm.org/wp-content/uploads/2024/04/1.3-Introduction-to-Knowledge-Model.pdf>
- CTDL-ASN terms: <https://purl.org/ctdlasn/terms/>
- 1EdTech CASE 1.1: <https://standards.1edtech.org/case/specifications/standards/v1p1>
- ELM launch: <https://europass.europa.eu/en/news/launch-european-learning-model-new-step-interoperability-learning>
- ELM browser: <https://europa.eu/europass/elm-browser/index.html>
