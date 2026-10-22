# Roadmap

- Glossary: define the bounded context's terms. No term has several meanings in the bounded context; meanings outside it are cited (e.g. "assessment").
- CV round trip: import an old CV, export a fresh one dedicated to a job offer, through CV standards ([design note](docs/design/cv-standards.md)).
- Deployment modes, updating ADR-0016: a demo `compose.yaml` running `tabellion serve --insecure` (no login, discreet "Demo mode: not for production" banner, published on `127.0.0.1` only, requests refused unless `Host` is `localhost`, `127.0.0.1` or `[::1]`); production setups as infrastructure as code, with login and HTTPS.
