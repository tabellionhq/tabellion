# ADR-0003: Professional skills referential

- Status: Accepted
- Date: 2026-10-15

## Decision

Professional skills are taken from ESCO. Technologies missing from ESCO (e.g. Rust, Kubernetes) are taken from Wikidata, as in ADR-0001.

## Criteria

- Good for engineering, computing, data.
- Soft and hard skills.
- International.
- Open (APR-0001).

## Alternatives

| Referential | Engineering, computing, data | Soft and hard | International | Open |
|---|---|---|---|---|
| ESCO v1.2 | ⚠️ ~13,900 skills; general (machine learning, data warehouse, Python); no Rust, no Kubernetes | ✅ transversal skills: thinking, social interaction, application of knowledge, attitudes and values | ✅ 28 languages; official crosswalk to O\*NET occupations (2022) | ✅ free reuse with attribution |
| O\*NET | ⚠️ occupations + "hot technologies" (Python, Go, Kafka, Spark…); no Rust | ✅ skills, work styles | ⚠️ US, English | ✅ CC BY 4.0 |
| SFIA 9 | ✅ built for digital professions | ✅ | ✅ | ❌ no distribution or credentials without fee |
| e-CF (EN 16234-1) | ⚠️ ICT roles only, 41 competences | ⚠️ | ⚠️ Europe | ❌ paid standard |
| Lightcast Open Skills | ✅ ~33,000 skills | ✅ | ✅ | ❌ commercial licence |
| Our own | | | | |

Our own is unthinkable: it is way too vast.

## Reason

None offers completeness or sufficient granularity, but the semantics framework (ADR-0001) covers this and more. ESCO is fully open, internationalised, interoperable (crosswalk), and that's all we need for now.

## Open

- The ESCO–O\*NET crosswalk maps occupations; not verified whether skills are mapped too.

## References

- ESCO: <https://esco.ec.europa.eu/en/classification/skill>
- ESCO–O\*NET crosswalk: <https://esco.ec.europa.eu/en/about-esco/publications/publication/crosswalk-between-esco-and-onet-technical-report>
- O\*NET hot technologies: <https://www.onetonline.org/search/hot_tech/>
- SFIA licensing: <https://sfia-online.org/en/licensing-sfia>
- [Skills standards benchmark](../research/skills-standards-benchmark.md)
- [ADR-0001: Describe concepts with Wikidata](ADR-0001-describe-concepts-with-wikidata.md)
