# Roadmap

- Glossary: define the bounded context's terms. No term has several meanings in the bounded context; meanings outside it are cited (e.g. "assessment").
- CV adapter: [CV roadmap](ROADMAP-CV.md).
- Specification review: trim and align the specification, the contract and the design documents; remove redundancies and contradictions.
- Deployment modes, updating ADR-0016: a demo `compose.yaml` running `tabellion serve --insecure` (no login, discreet "Demo mode: not for production" banner, published on `127.0.0.1` only, requests refused unless `Host` is `localhost`, `127.0.0.1` or `[::1]`); production setups as infrastructure as code, with login and HTTPS.
- Web UI state between steps: server state (the task at its own URL, with HTMX) or client state (a framework with a store such as Redux, saved in the browser session). The import keeps its steps on the server; revisit if a screen needs state the server does not hold.
- Development image: a Docker image kept locally with the design and build tools installed (Go, oapi-codegen, a YAML and OpenAPI validator), so that checking the contract does not download and install them on every run.
- Where the data goes: the UI shows what each plugin does and whether it runs on the local system. Before data leaves the local system, e.g. to an external language model provider, the person is told where it goes and agrees; without agreement, nothing is sent.
