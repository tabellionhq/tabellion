# ADR-0031: Plugin slow calls

- Status: Accepted
- Date: 2026-10-29
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

Plugins are services reached by URL (ADR-0028). The specification says what a plugin receives and returns (Section 7), not how it travels: this is the plugin HTTP binding, decided in four parts (ADR-0030 to ADR-0033). Already decided: a plugin declares its kinds and options and is called only for those (specification 4.10, 7.1); errors are problem details [RFC 9457] (7.1); TLS once a plugin is outside the host or private network (Security Considerations); a plugin keeps nothing it receives once it has answered (4.10).

### Problem

Some plugin calls are slow: term matching on a long text, by a language model, may take tens of seconds or more. How does the server wait for the result, show progress, and stop the work?

### Factors

- Stateless plugins (4.10): a plugin that holds a result until the server fetches it keeps personal data after the work is done.
- The call happens inside a task with stages, which the owner can stop (4.7), not inside the owner's request.
- How slow calls get is not known until the AI question is settled (design backlog: AI roles and search).
- Proxies between the server and a remote plugin may cut long requests.

## Options

| Option | Issues | Impacts |
|---|---|---|
| C. Synchronous: `POST`, the plugin answers when done; the server sets a timeout; closing the connection stops the work | A long call needs timeouts raised along the way | Nothing kept after the answer; stop is free; simplest for plugin authors |
| D. Asynchronous: `202 Accepted` with a job URL; the server polls, then fetches the result | The plugin holds the result until fetched: state and personal data kept after the work, against 4.10 | Survives proxy timeouts; progress per poll |
| E. Synchronous with a progress stream: the answer is a stream of progress events, then the result (e.g. NDJSON) | Streaming support needed in clients and proxies | Progress shown in the task's stages; nothing kept; stop by closing |

## Recommendation

C now; E when the AI question shows calls long enough to need progress. D is refused by the stateless rule. Timeouts and body size limits after the AI question.

## Decision

As recommended. C now; E when the AI question shows calls long enough to need progress. D is refused by the stateless rule. Timeouts and body size limits after the AI question.

## Reason

These options are simple and in line with our principles and previous ADRs.

## References

- [ADR-0028: Core and plugins as services](ADR-0028-core-and-plugins-as-services.md)
- RFC 9457, Problem Details for HTTP APIs: <https://www.rfc-editor.org/rfc/rfc9457>
