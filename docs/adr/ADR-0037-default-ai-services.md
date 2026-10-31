# ADR-0037: Default AI services

- Status: Proposed
- Date: 2026-10-31
- Drafted by: agent (Claude Code, claude-opus-5-5)

## Context

AI functions sit behind ports, with an adapter without AI in the core (ADR-0036). Bootstrapping the user well comes first, with simple interfaces and complexity discoverable later (APR-0002): which AI, if any, runs on first use is a default, not a question asked on day one.

Constraints:

- Everything can be done without AI (APR-0007).
- Data leaving the host or private network needs the owner's consent (specification, Section 4.12).
- A plugin keeps nothing it receives (Section 4.10).
- Users don't need to know about tokens and external providers early (APR-0002).

## Options

| Option | First use | Bootstrap, such as a CV import | Issues |
|---|---|---|---|
| 1. No AI | Core only | The owner types the facts, helped by index search | Slow first import |
| 2. Local AI services | Core and the `model` plugin on the same host | Depends on the local model | A model small enough to download and to run on an ordinary laptop, and good enough: not chosen yet |
| 3. Responsible AI services provided by Tabellion | Core and a plugin hosted by the project | Good, without an account at a provider | Data leaves the host: consent. The project runs and pays for a service (an escalation item). "Responsible" to define: e.g. open models, no retention (Section 4.10), no training on the data, hosting region |
| 4. AI service providers | Core and a plugin calling a provider | Best quality | Consent; an account or a key, against APR-0002 early on; the provider's terms |

All four can coexist: each is a plugin of the same kinds, or none. The question is which one an install offers first.

## Recommendation

2, with:

- The adapters without AI always present, as the fallback.
- Installed as a single binary, without plugins: 1 in practice; the UI says AI is available and how to add it.
- 3 and 4 never on by default: offered, with consent.
- The UI shows the technologies each plugin declares (ADR-0036).

## Decision

Option 2, local AI services. Kept open.

## Reason

Owner's words.

## Next steps

- Choose the local model: size, hardware, quality on the conformance test data.
- Define "responsible" for option 3, if it is pursued.
- Demo compose: core and `model` (backlog).

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [APR-0007: AI as a helper](../apr/APR-0007-ai-as-a-helper.md)
- [ADR-0028: Core and plugins as services](ADR-0028-core-and-plugins-as-services.md)
- [ADR-0036: Architecture of AI functionalities](ADR-0036-architecture-of-ai-functionalities.md)
