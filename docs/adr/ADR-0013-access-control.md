# ADR-0013: Access control

- Status: Accepted
- Date: 2026-10-18
- Drafted by: agent (Claude Code, claude-opus-5-5) from owner's design; reason by owner

## Context

A view says what a person shares (US-06). Several readers may read the same view, and the person must be able to stop one reader without affecting the others. The API is all HTTP (ADR-0012). Access must stay invisible at first (APR-0002): users don't need to know about tokens and external providers early.

## Options

| Option | Description | Pros | Cons |
|---|---|---|---|
| A. Grantee inside the view | Each view names its reader | Simple | One view per reader; stopping a reader means editing or deleting the view |
| B. Grants, separate from views | A view says what; a grant says who reads it, how they authenticate, until when | Many readers per view; revoke one reader; audit per reader | One more object |
| C. Delegate entirely to an external system | An identity provider decides who reads what | No access code of our own | Mandatory infrastructure from day one; contradicts APR-0002 |

## Decision

B. Access control is a layer of its own, pluggable like databases do (e.g. PostgreSQL's `pg_hba.conf`).

- **View**: what is shared (selection and projection of facts).
- **Grant**: who reads a view: one per reader; status active or revoked; optional expiry. Revoking a grant stops that reader only.
- **Auth adapters**, by adoption level (APR-0002):

| Level | Access | The person needs to know |
|---|---|---|
| 0 | Exported file: no access control, no revocation; it is a copy | what they share |
| 1 | Unix socket owned by the user: full access to their own record | nothing |
| 2 | Bearer token per reader | "share with Michaela" |
| 3 | OpenID Connect, for organisations | only if their organisation needs it |

- **Remote access** (TCP) always uses encryption (TLS) (ADR-0012).
- **Handshake**: the reader receives the person's identifier (ADR-0010) and keeps it with its credential.
- **Access log**: who read which view, when. It holds data about readers.
- Revoking stops future reads; copies already taken stay with the reader. Freshness of copies is ADR-0011.

## Reason

Separating grants from views offers flexibility and control. It matches our principles: self control and easy adoption, delegating to existing protocols.

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [ADR-0007: The person decides what is exposed](ADR-0007-person-decides-exposed-views.md)
- [ADR-0010: Person identifier](ADR-0010-person-identifier.md)
- [ADR-0011: View freshness](ADR-0011-view-freshness.md)
- [ADR-0012: HTTP API](ADR-0012-http-api.md)
- [US-06: Share a view](../design/user-stories/US-06-share-a-view.md)
