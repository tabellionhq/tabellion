# ADR-0010: Person identifier

- Status: Accepted
- Date: 2026-10-18
- Drafted by: agent (Claude Code, claude-opus-5-5) and owner; reason by owner

## Context

The exchange format is the data model, serialised (US-07). Facts from one person travel to other systems, so a person needs an identifier that is unique across systems.

## Options

| Option | Pros | Cons |
|---|---|---|
| A. UUID v4 | Simple; universal; opaque; never changes | Proves nothing by itself |
| B. DID (W3C Decentralized Identifier) | Ownership provable by signature; no server needed (did:key) | Keys and wallets to manage; key loss is identity loss; young tooling |
| C. WebID (Solid project) | Resolvable URL; linked data | Needs hosting; trust in the host; adoption not verified |
| D. Email | Familiar | Personal data; changes over time; owned by the provider |

## Decision

A: a UUID v4 is the person's only identifier; facts carry only it.

Label and contact belong to **identities**: a person has several, and chooses one per view. They never appear on facts or in the identifier.

- **Identifier**: the UUID; never changes; on every fact.
- **Identity**: a label and an optional contact, presented to a reader (known in privacy terminology as a partial identity).

The identifier is negotiated at the contract's handshake; the receiving software keeps it with its access credential. Proof and cryptography belong to access control, not to the identifier.

## Reason

Every fact would otherwise bear transient information. What if Anne was Andrew three years before? What if I use multiple emails and want to communicate one or another to a recruiter?

## References

- [APR-0001: Free and open](../apr/APR-0001-free-and-open.md)
- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- Pfitzmann, Hansen: A terminology for talking about privacy by data minimization (partial identities)
- W3C Decentralized Identifiers: <https://www.w3.org/TR/did-core/>
- Solid project: <https://solidproject.org>
