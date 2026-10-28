# ADR-0025: Licences

- Status: Accepted
- Date: 2026-10-22
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

The project must be free and open: the person owns and can export everything, and the schema is open so other systems can use it (APR-0001). The specification is the source of truth, written to be implemented by others (ADR-0024). The software may be run as a hosted service, by us or by others (APR-0005).

The goals pull in two directions: the implementation should be free as in freedom, and the work should be reusable, commercial uses included.

Creative Commons advises against its licences for software: they do not address source code distribution or patents.

The Go toolchain links libraries statically: a library imported by the server is part of the distributed binary.

## Options

| Option | Specification, docs | Contract, JSON-LD context, test data | Server, web UI | Issues |
|---|---|---|---|---|
| A. All permissive | CC BY 4.0 | Apache-2.0 | MIT or Apache-2.0 | A copy may be modified and hosted, its changes closed |
| B. Split, copyleft server | CC BY 4.0 | Apache-2.0 | AGPL-3.0 | Some companies do not use AGPL software; they can still implement the specification |
| C. All copyleft | CC BY-SA 4.0 | GPL-3.0 | AGPL-3.0 | Clients and other implementations built on the contract must be copyleft too |

Contributions:

| Option | How | Issues |
|---|---|---|
| DCO | Each commit is signed off: the contributor certifies the right to submit it under the project licence | The licence cannot be changed later without every contributor's consent |
| CLA | Contributors grant the project rights over their contributions | Allows relicensing and dual licensing; contributors often distrust it |

## Recommendation

B, with a DCO.

- The specification lets anyone build an implementation, open or closed, and sell it.
- The contract, the vocabulary and the test data carry no copyleft: anyone builds clients and implementations against them.
- Our server stays free: whoever modifies it, also when only hosting it, shares the changes.
- With a DCO, the server cannot be closed later: a promise to contributors.

## Decision

B, the split with a copyleft server, AGPL-3.0, and a DCO for contributions.

## Reason

This might be far-fetched for now, but it is a sensible approach.

## Consequences

- Licences by part:
  - `docs/RFC.md` and the documentation: CC BY 4.0.
  - `api/` (OpenAPI contract, JSON-LD context) and `testdata/`: Apache-2.0.
  - Server, web UI, command line: AGPL-3.0.
- Files: `LICENSE` (AGPL-3.0) at the root; a licence file in `api/`, `testdata/` and `docs/`; the README says which licence covers what.
  Updated 2026-10-28 by the owner: the licence texts live in `licences/`, one file per licence.
- An SPDX identifier in each source file header, e.g. `SPDX-License-Identifier: AGPL-3.0-only` or `AGPL-3.0-or-later`: to choose when the first source file is written.
- Dependencies compiled into the server: licences compatible with AGPL-3.0 only. MIT, BSD and Apache-2.0 are; GPL-2.0-only is not. Programs run as separate processes are not affected.
- CI checks the licences of all Go dependencies, direct and indirect (e.g. `go-licenses`), and refuses incompatible or unknown licences.
- Every commit carries a `Signed-off-by` line (DCO); CI refuses commits without it.
- The `license` field of the API contract: Apache-2.0.

## References

- [APR-0001: Free and open](../apr/APR-0001-free-and-open.md)
- [APR-0005: Hosted service possible](../apr/APR-0005-hosted-service-possible.md)
- [ADR-0020: Target platforms](ADR-0020-target-platforms.md)
- [ADR-0024: The specification is the source of truth](ADR-0024-specification-source-of-truth.md)
- GNU AGPL-3.0: <https://www.gnu.org/licenses/agpl-3.0.html>
- Apache License 2.0: <https://www.apache.org/licenses/LICENSE-2.0>
- Apache-2.0 and GPL compatibility: <https://www.apache.org/licenses/GPL-compatibility.html>
- CC BY 4.0: <https://creativecommons.org/licenses/by/4.0/>
- Creative Commons on software: <https://creativecommons.org/faq/#can-i-apply-a-creative-commons-license-to-software>
- Developer Certificate of Origin: <https://developercertificate.org/>
