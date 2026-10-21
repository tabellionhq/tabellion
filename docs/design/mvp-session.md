# MVP session

- Status: Draft
- Stories: US-01 import from a CV, US-08 new CV filtered for an offer
- Contract: [api/openapi.yaml](../../api/openapi.yaml)

Anne's first session, step by step: her action in the web UI, the request it makes, the components called, what is stored.

The web UI calls the same API as `curl`. Below, `$API` stands for the server address:

```shell
# command line, through the user's Unix socket
curl --unix-socket "$XDG_RUNTIME_DIR/skills-track.sock" http://localhost/
# web UI, through localhost, with the session cookie
curl -b "session=…" http://localhost:$PORT/
```

## 0. Start

| | |
|---|---|
| Anne | runs `skills-track ui` |
| Components | client → starts the server if not running (socket) → server opens a localhost port for the UI, creates a one-time token → client opens the browser at `http://localhost:$PORT/ui/?token=…` |
| Request | `GET /ui/?token=…` → `303 See Other`, `Set-Cookie: session=…; HttpOnly; SameSite=Strict` |
| Stored | first run only: the store, one Person (`id: uuid`) |
| UI shows | home: the actions of `GET /` |

```shell
curl $API/
```

```json
{"@id": "/", "@type": "Entry", "person": "/persons/9f1e…",
 "potentialAction": [
   {"@type": "CreateAction", "name": "Import a CV", "target": {"httpMethod": "POST", "urlTemplate": "/imports"}},
   {"@type": "SearchAction", "name": "Review proposed practices", "target": {"httpMethod": "GET", "urlTemplate": "/persons/9f1e…/practices?status=proposed"}},
   {"@type": "ViewAction", "name": "My views", "target": {"httpMethod": "GET", "urlTemplate": "/persons/9f1e…/views"}},
   {"@type": "ViewAction", "name": "My identities", "target": {"httpMethod": "GET", "urlTemplate": "/persons/9f1e…/identities"}}
 ]}
```

## 1. Identity for the CV, when first needed

| | |
|---|---|
| Anne | enters the name and email to put on CVs |
| Request | `POST /persons/9f1e…/identities` |
| Components | core |
| Stored | Identity `{id, person, label: "Anne Martin", contact: "anne@example.org"}` |
| UI shows | identities list |

```shell
curl -X POST $API/persons/9f1e…/identities -H 'Content-Type: application/ld+json' \
  -d '{"label": "Anne Martin", "contact": "anne@example.org"}'
```

## 2. Import the CV

| | |
|---|---|
| Anne | pastes her CV text, or picks a `.txt` file saved from LibreOffice |
| Request | `POST /imports` (text) → `202 Accepted`, `Location: /tasks/41ab…` |
| Components | core creates a task → extraction adapter (local language model): statements (occupation, period, terms), each with its excerpt → search adapter: ESCO index for skills and occupations, embeddings for Wikidata concepts → core: proposed facts |
| Stored | Task `{id, kind: cv-import, state, result: {unresolved}}`; OccupationHeld and Practice rows, `status: proposed`, `author: /tasks/41ab…`, `fidelity`, `evidence: {excerpt}` |
| UI shows | progress, then the import summary |

```shell
curl -X POST $API/imports -H 'Content-Type: text/plain' --data-binary @cv.txt
curl $API/tasks/41ab…          # until "state": "done"
```

```json
{"@id": "/tasks/41ab…", "@type": "Task", "kind": "cv-import", "state": "done",
 "result": {"unresolved": [{"term": "Zephyr RTOS", "excerpt": "Firmware on Zephyr RTOS, 2019"}]},
 "potentialAction": [
   {"@type": "SearchAction", "name": "Review the proposed occupations", "target": {"httpMethod": "GET", "urlTemplate": "/persons/9f1e…/occupations?status=proposed"}},
   {"@type": "SearchAction", "name": "Review the proposed practices", "target": {"httpMethod": "GET", "urlTemplate": "/persons/9f1e…/practices?status=proposed"}}
 ]}
```

## 3. Review

| | |
|---|---|
| Anne | accepts or rejects each proposed occupation, then each proposed practice; adds what the import missed |
| Requests | `GET …/occupations?status=proposed`, `PATCH /occupations/{id}`; `GET …/practices?status=proposed`, `PATCH /practices/{id}`; for a missed or unresolved term: `GET /reference/skills{?q}`, `GET /reference/concepts{?q}`, `POST …/practices` |
| Components | core; search adapter for the lookups |
| Stored | `status` moves to accepted or rejected, once; new Practices, `author: anne` |
| UI shows | one card per proposed fact: occupation, period, skill and concept labels, level, the CV excerpt; Accept, Reject |

```shell
curl -X PATCH $API/practices/7d2c… -H 'Content-Type: application/merge-patch+json' -d '{"status": "accepted"}'
curl "$API/reference/skills?q=embedded%20firmware&language=en"
```

## 4. View for an offer

| | |
|---|---|
| Anne | creates a view for an offer: picks occupations, a period, skill groups; removes some facts |
| Requests | `GET /reference/skills{?q}` to find skill groups; `POST /persons/9f1e…/views`; `GET /views/{id}/facts` as a preview; `PUT /views/{id}` with `If-Match` to adjust |
| Components | core: the deterministic view filter |
| Stored | View `{id, person, label, include, include_facts, exclude_facts, show}` |
| UI shows | the view form, and the preview of the selected facts |

```shell
curl -X POST $API/persons/9f1e…/views -H 'Content-Type: application/ld+json' -d '{
  "label": "Backend developer, ACME offer",
  "include": {"occupation_kind": ["work"], "period": {"from": "2019-01-01"},
              "skill_under": ["http://data.europa.eu/esco/skill/…"]},
  "exclude_facts": ["/practices/88e0…"],
  "show": {"level": true, "evidence": false}}'
curl $API/views/5b7a…/facts
```

## 5. New CV

| | |
|---|---|
| Anne | exports the view as a CV, with her identity |
| Request | `GET /views/5b7a…/resume?identity=e2d4…` |
| Components | core: view filter → rendering adapter: JSON Resume |
| Stored | nothing |
| UI shows | download of `resume.json`; rendering to PDF or HTML is done by a JSON Resume theme, outside `skills-track` |

```shell
curl "$API/views/5b7a…/resume?identity=e2d4…" -o resume.json
```

## Stored at the end of the session

| Object | Rows | Source |
|---|---|---|
| Person | 1 | first run |
| Identity | 1 | Anne |
| Task | 1 | the CV import |
| OccupationHeld | one per occupation in the CV | extraction, reviewed |
| Practice | one per skill on a concept, per occupation | extraction, reviewed; plus Anne's additions |
| View | 1 | Anne |
| Reference data | ESCO skills, occupations, ancestors, labels; Wikidata concepts, labels; vectors | loaded before the first import |

## Gaps found

1. Reference data at first run: ESCO shipped with the release; Wikidata looked up online and cached.
2. Import in the data model: a Task; facts it proposes have it as author.
3. Labels in responses: embedded, in the requested language.
4. Correcting a proposed fact: edited until reviewed (US-09).
5. Rejected occupation: its practices stay. Removing them is Anne's manual action, deleting facts by a filter; deleting facts is allowed (ADR-0008).
6. Raw CV text: not kept once the task ends. Each fact keeps its excerpt, the file name and the position in the file.
7. Socket path: `$XDG_RUNTIME_DIR/<name>/` on Linux; macOS and Windows: ADR on target platforms. `<name>`: the project name, not decided.
8. Identity: not asked at start. Asked when first needed: after a CV import, proposed from the CV; or at the first export or share, typed by Anne. Can be changed or aliased later.
9. Review screen: wireframes, and an ADR for the frontend framework. A task of its own.
