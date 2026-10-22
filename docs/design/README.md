# Design

Functional design, data first.

| Document | Holds |
|---|---|
| [personas.md](personas.md) | who uses the system |
| [user-stories/](user-stories/README.md) | catalogue, data in/out, detailed stories |
| [data-model.md](data-model.md) | objects, facts and dimensions |
| [functions.md](functions.md) | functions, traced to stories |
| [resources.md](resources.md) | resources exposed by the HTTP API |
| [adapters.md](adapters.md) | adapters by port |
| [mvp-session.md](mvp-session.md) | the MVP session, request by request |
| [ui-functions.md](ui-functions.md) | what the web UI does itself, and the server actions it shows |
| [wireframes/review.md](wireframes/review.md) | review screen: layout, actions, focus, WCAG 2.2 AA |
| [data-first.md](data-first.md) | design position and its lineage |
| [cv-standards.md](cv-standards.md) | CV import and export through existing standards |

## Method

1. Take one user story.
2. State what it requires abstractly: data in, data out, functions, data model impacts, open questions.
3. Update the data model and the functions with what the story requires.
