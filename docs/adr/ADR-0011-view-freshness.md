# ADR-0011: View freshness

- Status: Accepted
- Date: 2026-10-18
- Drafted by: agent (Claude Code, claude-opus-5-5) and owner; reason by owner

## Context

A reader keeps a copy of a view (US-07: consumers store and search at scale). When the person changes the view or its facts, the copy becomes stale. The reader needs to know whether its copy is still fresh.

Access control (who may read, revoking a reader) is a question of its own (ADR-0012).

## Options

| Option | Description | Pros | Cons |
|---|---|---|---|
| A. HTTP caching semantics | The view carries a version (ETag), a last-modified date and an expiry (max-age); readers check validity cheaply (HEAD, conditional GET answering 304 Not Modified) | Standard, known by every HTTP client and proxy; works for pull; metadata can also sit in an exported file | Readers must poll |
| B. Feed style | Views are feeds of facts with an update date and a refresh hint (RSS `ttl`, iCalendar `REFRESH-INTERVAL`) | Easy catch-up from the last entry read | Partly overlaps A; no standard for "entry withdrawn" |
| C. Push | The system notifies readers of changes (webhooks, the event stream of ADR-0008) | Fresh without polling | Readers must expose an endpoint; more moving parts; contradicts a simple start (APR-0002) |
| D. Expiry only | Each view copy expires after a set time | Simplest | No way to know earlier that a copy is stale |

## Decision

A view carries freshness metadata, based on HTTP caching semantics: version, last modified, expiry. Readers verify validity before reusing a copy. The same metadata sits in the header of an exported file (level 0, APR-0002), so a file says how old it is and when to ask again.

Push (C) may come later, on top of A.

## Reason

Alignment with the APRs: don't reinvent the wheel; HTTP is a great way to expose resources.

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [APR-0003: Skills curriculum standard, existing standards for the rest](../apr/APR-0003-curriculum-standard-defer-the-rest.md)
- RFC 9111, HTTP Caching: <https://www.rfc-editor.org/rfc/rfc9111>
- RFC 9110, HTTP Semantics (ETag, conditional requests, HEAD): <https://www.rfc-editor.org/rfc/rfc9110>
