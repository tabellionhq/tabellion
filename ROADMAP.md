# Roadmap

- Glossary: define the bounded context's terms. No term has several meanings in the bounded context; meanings outside it are cited (e.g. "assessment").
- CV round trip: import an old CV, export a fresh one dedicated to a job offer, through CV standards ([design note](docs/design/cv-standards.md)).
- Deployment modes, updating ADR-0016: a demo `compose.yaml` running `tabellion serve --insecure` (no login, discreet "Demo mode: not for production" banner, published on `127.0.0.1` only, requests refused unless `Host` is `localhost`, `127.0.0.1` or `[::1]`); production setups as infrastructure as code, with login and HTTPS.
- Web UI state between steps: server state (the task at its own URL, with HTMX) or client state (a framework with a store such as Redux, saved in the browser session). The import keeps its steps on the server; revisit if a screen needs state the server does not hold.
- Development image: a Docker image kept locally with the design and build tools installed (Go, oapi-codegen, a YAML and OpenAPI validator), so that checking the contract does not download and install them on every run.
- Align the model on JSON Resume, an existing standard (APR-0003); more parts may need fixing along the way ([cv-standards.md](docs/design/cv-standards.md)):
  - Occupation held gains its details: position, organisation, location, url, summary, highlights. The import copies them as written in the CV; a weak CV leaves them empty and completion shows it.
  - Identity gains the `basics` fields: headline, summary, location, links, several contacts.
  - A language model as an optional helper, e.g. drafting a summary from the practices; generated text is a draft the person accepts or edits, and never changes a fact.
