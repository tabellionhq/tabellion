# Resources

- Status: Draft
- Specification: [docs/RFC.md](../RFC.md) prevails over this note; this note describes its HTTP binding
- Decisions: ADR-0011 (view freshness), ADR-0012 (HTTP API), ADR-0013 (access control); ADR-0015 (Richardson level 3, JSON-LD), ADR-0016 (GUI in the binary)
- Model: [data-model.md](data-model.md); functions: [functions.md](functions.md)

The HTTP API exposes the data model as resources. Every client goes through it: command line, GUI, MCP adapter, other systems.

## Rules

- Clients start at `/` and follow links and actions; they never build URLs.
- The contract: [api/openapi.yaml](../../api/openapi.yaml), with the actions each resource carries (`x-actions`), and the JSON-LD context [api/contexts/curriculum.jsonld](../../api/contexts/curriculum.jsonld). Paths are informative.
- Every response says what can be done next, as actions with method, target, expected input and a title in plain words (ADR-0015).
- The core returns structured data only. Human formats (CV, PDF, page) are adapters outside the core.
- Readers reach a record only through `/views/{id}/facts` (specification, Section 4.6).
- Accepted or rejected facts have no `PUT` or `PATCH` (specification, Section 4.1). Deleting facts by the owner is not in the contract yet.
- Freshness: `ETag`, `Last-Modified`, `Cache-Control` on every read; `HEAD` and `304 Not Modified` supported (ADR-0011).

## Resources

| Resource | Methods | Holds | Function |
|---|---|---|---|
| `/` | GET | entry point: links to what the caller may reach | |
| `/persons/{id}` | GET | Person | |
| `/persons/{id}/identities` | GET, POST | Identity | |
| `/persons/{id}/occupations` | GET, POST | OccupationHeld | F03 |
| `/occupations/{id}` | GET; PATCH `status` while proposed | OccupationHeld | F04 |
| `/persons/{id}/practices` | GET, POST | Practice; `?status=proposed` is the review queue | F03, F04 |
| `/practices/{id}` | GET; PATCH `status` while proposed | Practice | F04 |
| `/persons/{id}/competencies` | GET | Competency, derived; `?related=<text>` | F05 |
| `/persons/{id}/views` | GET, POST | View | F06 |
| `/views/{id}` | GET, PUT | View; owner only | F06 |
| `/views/{id}/facts` | GET, HEAD | the record through the view | F07, F10 |
| `/views/{id}/grants` | GET, POST | Grant | |
| `/grants/{id}` | GET; PATCH `status: revoked` | Grant | |
| `/grants/{id}/access-log` | GET | AccessLog | |
| `/imports` | POST | a document (CV) to extract; returns `202 Accepted` and a task | F01, F03 |
| `/tasks/{id}` | GET | task state and result | F01 |
| `/tasks/{id}/rejections` | POST | reject every fact still proposed by the task | F04 |
| `/reference/skills?q=` | GET | ESCO skills | F02 |
| `/reference/skills/{id}/ancestors` | GET | ESCO categories above a skill | F02 |
| `/reference/concepts?q=` | GET | Wikidata concepts | F02 |

## State changes

The status of facts: specification, Section 4.1. Grants: active to revoked (Section 4.6). Views: changed freely by their owner (Section 4.5).

## Example: a proposed practice

`GET /practices/7d2c…` returns, in compacted JSON-LD (ADR-0015):

```json
{
  "@context": "/contexts/curriculum.jsonld",
  "@id": "/practices/7d2c…",
  "@type": "Practice",
  "skill": "http://data.europa.eu/esco/skill/…",
  "concept": {"is": "http://www.wikidata.org/entity/Q1071360", "with": [], "label": null},
  "context": "http://www.wikidata.org/entity/Q575650",
  "from": "2024-03-01",
  "level": 3,
  "author": "cv-import",
  "status": "proposed",
  "potentialAction": [
    {
      "@type": "AcceptAction",
      "name": "Accept this practice into the record",
      "target": {"@type": "EntryPoint", "httpMethod": "PATCH", "urlTemplate": "/practices/7d2c…", "contentType": "application/ld+json"},
      "object": {"status": "accepted"}
    },
    {
      "@type": "RejectAction",
      "name": "Reject this practice",
      "target": {"@type": "EntryPoint", "httpMethod": "PATCH", "urlTemplate": "/practices/7d2c…", "contentType": "application/ld+json"},
      "object": {"status": "rejected"}
    }
  ]
}
```

Once accepted, the same resource has no action that changes it.

## Open

- Later, not needed yet: others (Ismael, Claude) proposing facts to a person's record, through a grant with write rights or an inbox. The model allows it: facts arrive as `proposed` and the person reviews them.
- Whether `/persons/{id}` stays in the path when a server holds one person only (level 1).
