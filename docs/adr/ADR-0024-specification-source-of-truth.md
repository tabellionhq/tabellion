# ADR-0024: The specification is the source of truth

- Status: Accepted
- Date: 2026-10-22
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

The skills curriculum is our standard; existing standards cover the rest (APR-0003).

Several clients implement the curriculum rules: the core web UI, the command line, agents, and apps built by others that work offline with their own rules on the device (US-10). The server is the authority: it checks every change it receives.

Today the model and the rules are spread over design notes (`data-model.md`, `resources.md`, user stories), the API contract (`api/openapi.yaml`) and the JSON-LD context (`api/contexts/curriculum.jsonld`). A client implementer has no single text to follow, and the documents can drift apart.

## Options

| Option | Source of truth | Issues |
|---|---|---|
| A. The code | The Go implementation | Others must read Go to implement a client; the rules change with every refactoring |
| B. The API contract | `api/openapi.yaml` | Describes HTTP only: no rule for a client working offline, or for another binding; rules such as frozen facts fit poorly in a schema |
| C. A specification | A written specification of the model and the rules, in RFC style | One more document to keep; must be kept in step with the contract |

## Recommendation

C.

## Decision

C. The specification, beta for now, is the source of truth. It uses the requirement levels MUST, SHOULD and MAY (RFC 2119, RFC 8174). The OpenAPI contract is an interpretation of it: it follows the obligations and chooses among the options. The specification may cite the contract as a link, for reference.

## Reason

This solves many problems, and we want it to be really open. An RFC is presumptuous, but it is the best way found so far to be really interoperable up front. Plus, in an agent-enhanced development approach, it provides a textual source of truth.

## Consequences

- Order of precedence: the specification, then the OpenAPI contract, then the code. A conflict is a bug in the lower one.
- The specification holds the data model, the rules (statuses, frozen facts, deleting, view semantics, completing a proposed fact) and the vocabulary; the JSON-LD context implements the vocabulary.
- The design notes become its draft; once in the specification, a rule is no longer restated in them, only linked.
- The specification is [docs/RFC.md](../RFC.md), named and numbered as an Internet-Draft (`draft-tabellion-curriculum-00`): it is not an RFC, and RFCs are not versioned with semantic versions. Licence: CC BY 4.0.
- Conformance: the test data ([testdata](../../testdata/README.md)) becomes the start of a conformance suite for servers and clients.

## References

- [APR-0003: Skills curriculum standard, existing standards for the rest](../apr/APR-0003-curriculum-standard-defer-the-rest.md)
- [ADR-0015: API maturity, Richardson level 3](ADR-0015-api-maturity-richardson-level-3.md)
- [US-10: Install an app, run no server](../design/user-stories/US-10-hosted-mobile.md)
- RFC 2119, key words for requirement levels: <https://www.rfc-editor.org/rfc/rfc2119>
- RFC 8174, uppercase key words only: <https://www.rfc-editor.org/rfc/rfc8174>
