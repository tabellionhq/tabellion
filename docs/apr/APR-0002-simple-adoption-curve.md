# APR-0002: Simple adoption curve

- Status: Accepted
- Date: 2026-10-18

## Principle

**Bootstrapping the user well is the main priority.** If we don't manage to create the database fast, make it easy to plug in with agents, while making the user feel safe and satisfied, we wrote code for nothing.

The project aims for a simple adoption curve. It must aim for simple interfaces, hiding away the potential complexity and making it discoverable later.

Case in point: the first level of export is a file that's human and machine readable, without any cryptography or notion of access or revocation for the receiving part. Same for access: users don't need to know about tokens and external providers early.
