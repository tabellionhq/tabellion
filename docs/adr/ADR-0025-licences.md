# ADR-0025: Licences

- Status: Proposed
- Date: 2026-10-22
- Drafted by: agent (Claude Code, claude-opus-5-5); CC BY 4.0 for the specification and Apache-2.0 preference stated by owner

## Context

The project must be free and open: the person owns and can export everything, and the schema is open so other systems can use it (APR-0001). The specification is the source of truth, written to be implemented by others (ADR-0024). A hosted service may run the software (APR-0005).

The specification is licensed CC BY 4.0. Creative Commons advises against its licences for software: they do not address source code distribution or patents. The code needs a software licence.

The direct dependencies chosen so far are under MIT, Apache-2.0 or BSD-3-Clause licences (ADR-0022, ADR-0023). Apache-2.0 code can be combined with GPL-3.0 code, not with GPL-2.0-only code.

## Options

For the code (Go server, web UI, command line):

| Option | Others may | Others must | Issues |
|---|---|---|---|
| A. MIT | use, change, sell, host, close their changes | keep the copyright notice | No patent grant |
| B. Apache-2.0 | use, change, sell, host, close their changes | keep notices; state changes; keep the NOTICE file | Explicit patent grant from contributors, ended for anyone who sues over patents; not combinable with GPL-2.0-only code |
| C. MPL-2.0 | use, combine with closed code | publish their changes to our files when they distribute | Weak copyleft per file; hosting is not distribution |
| D. AGPL-3.0 | use, change, host | publish their changes, also when they only host the software | Strong copyleft; companies running it on premises or as a service often avoid it |

## Recommendation

B, Apache-2.0, for the code, the API contract and the JSON-LD context. CC BY 4.0 for the specification and the documentation.

- The patent grant protects implementers of the specification and contributors.
- The software can be hosted, by us or by others, without conditions (APR-0005).
- Trade-off: a hosted copy may keep its changes closed. The specification stays open whatever a copy does.

## Decision

Owner's words.

## Reason

Owner's words.

## Consequences

- `LICENSE` at the root with the Apache-2.0 text; `NOTICE` with the copyright line; `docs/LICENSE` with CC BY 4.0, or a line in the README saying which licence covers what.
- An SPDX identifier in each source file header: `SPDX-License-Identifier: Apache-2.0`.
- CI checks the licences of all dependencies, direct and indirect (e.g. `go-licenses`), and refuses copyleft or unknown licences.
- The `license` field of the API contract names the chosen licence.
- Open: how contributions are accepted, e.g. a Developer Certificate of Origin sign-off.

## References

- [APR-0001: Free and open](../apr/APR-0001-free-and-open.md)
- [APR-0005: Hosted service possible](../apr/APR-0005-hosted-service-possible.md)
- [ADR-0024: The specification is the source of truth](ADR-0024-specification-source-of-truth.md)
- Apache License 2.0: <https://www.apache.org/licenses/LICENSE-2.0>
- Apache-2.0 and GPL compatibility: <https://www.apache.org/licenses/GPL-compatibility.html>
- Creative Commons on software: <https://creativecommons.org/faq/#can-i-apply-a-creative-commons-license-to-software>
- CC BY 4.0: <https://creativecommons.org/licenses/by/4.0/>
- MPL-2.0: <https://www.mozilla.org/en-US/MPL/2.0/>
- AGPL-3.0: <https://www.gnu.org/licenses/agpl-3.0.html>
