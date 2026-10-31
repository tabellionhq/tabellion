# ADR-0028: Core and plugins as services

- Status: Accepted
- Date: 2026-10-23
- Drafted by: agent (Claude Code, claude-opus-5-5); decision and reason by owner

## Context

Anne's session needs three jobs besides the core: reading a CV (PDF, DOCX, ODT), matching its terms to ESCO skills and Wikidata concepts (embeddings, later a language model), and rendering a CV from JSON Resume with one of its themes. The JSON Resume themes are npm packages, rendered by Node.js. Embeddings and language models run best in native code (ADR-0014, ADR-0020).

ADR-0017 compiles adapters into `tabellion` and rejects plugins. ADR-0014 and ADR-0020 already allow heavy adapters to run as separate processes, in any language.

Shipping every job in one image makes the download heavy, mostly because of the model weights, for people who may need only part of it.

## Options

| Option | How | Pros | Cons |
|---|---|---|---|
| A. Adapters compiled in (ADR-0017) | Every adapter in the Go binary | One binary, one toolchain | No Node.js themes; native models break cross-compilation (ADR-0020); every download carries everything |
| B. Plugins as services | `tabellion` is the core; each plugin is an optional service reached by URL, in the language that fits its job, behind a contract in the specification | Each image is downloaded only when needed; any provider can implement a plugin; the core stays cross-compiled | A contract to specify and version; several toolchains, CIs and images; calls between services need auth once they leave localhost |
| C. Executable plugins (ADR-0017 B) | Versioned executables in a directory, started by `tabellion` | No network between core and plugins | Needs a runtime on the host for each language; a sandbox for trust; does not run as containers |

## Recommendation

B.

| Service | Language | Role | Required |
|---|---|---|---|
| `tabellion` | Go | Core: facts, rules, tasks, API, web UI; calls plugins by URL | yes |
| `cv` | TypeScript | CV adapters: text from PDF, DOCX, ODT; JSON Resume rendered with a theme | no |
| `model` | Rust | Embeddings and term matching; a language model later | no |
| other plugins | any | Other providers of the same contracts | no |

- The specification defines a plugin and a client application, and the contracts between them and the core.
- Plugins hold no state of their own, or a cache they can rebuild (ADR-0017): a feature with its own store stays an external system.
- The core stays the authority (RFC, Section 4.9): a plugin proposes, the core decides.
- The core works without any plugin: facts entered by hand, record exported as JSON Resume.

## Decision

B. The CV adapters are in their own image, optional in the long run, written in TypeScript to fit their role. `tabellion` is the pluggable core, in Go. The basic model is in Rust; other plugins may use other providers.

The specification describes what a plugin is and what a client application is. The contracts are in the specification. We build plugins, not client applications: client applications are left to others.

The demo compose file serves our use case: the core and two services, `cv` and `model`.

## Reason

Plugins keep the core clean and adapt it to different use cases.

## Consequences

- ADR-0017: plugins are no longer rejected; its rule stays: a feature that needs its own store is an external system.
- ADR-0016: the web UI stays in the core; it shows a plugin's work through task stages and actions (RFC, Section 4.7).
- The specification gains the definitions of plugin and client application, and a contract per plugin kind: CV extraction, term matching, CV rendering.
  Updated 2026-10-28 by the owner: CVs are an adapter concern. The specification keeps term matching, on text; CV extraction and CV rendering move to the CV adapter contracts ([cv-adapter.md](../design/cv-adapter.md)).
- Each plugin has its own image, CI, licence check and release.
- The web UI says what a missing plugin prevents, and how to start it.

## References

- [APR-0002: Simple adoption curve](../apr/APR-0002-simple-adoption-curve.md)
- [APR-0003: Skills curriculum standard, existing standards for the rest](../apr/APR-0003-curriculum-standard-defer-the-rest.md)
- [ADR-0014: Language](ADR-0014-language.md)
- [ADR-0016: GUI in the main binary](ADR-0016-gui-in-the-binary.md)
- [ADR-0017: Where a feature belongs](reversed/ADR-0017-where-a-feature-belongs.md)
- [ADR-0020: Target platforms](ADR-0020-target-platforms.md)
- [ADR-0024: Specification as the source of truth](ADR-0024-specification-source-of-truth.md)
- JSON Resume themes: <https://jsonresume.org/themes>
