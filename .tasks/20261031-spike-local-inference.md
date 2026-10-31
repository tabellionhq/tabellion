# Task: Spike, local inference

- Date: 2026-10-31
- Model: Claude Opus 5.5 (claude-opus-5-5)
- Tool: Claude Code
- Status: Open, not started
- For: [ADR-0037 default AI services](../docs/adr/ADR-0037-default-ai-services.md), proposed
- Kind: research, throwaway code in `spikes/`, run in Docker

## Goal

Evidence for the default AI services of an install (ADR-0037, option 2): can local models serve term matching (F01, specification 7.2) and search (F02, F05) on an ordinary machine?

## Hypotheses (owner)

| # | Hypothesis | Measured by | Refuted when |
|---|---|---|---|
| H1 | Downloading embedding and inference models is too much for a demo | Download size of each model, and of both; time to first answer from a fresh install | Total download and first-answer time under the thresholds |
| H2 | Local language model text is OK: term matching on a CV gives usable facts | Import of `testdata/anne/cv.txt` compared with `testdata/anne/facts.yaml`: facts found, missed, invented; excerpts found in their line (7.2) | Quality under the threshold |
| H3 | Embeddings search is efficient but long | Efficient: related concepts found (US-02 acceptance: related, no false friends, any language, empty is valid). Long: time to build the vectors of ESCO and Wikidata labels, and time per query | Quality under the threshold, or time under the threshold |

## Setup

- Machine: to set by the owner, e.g. 16 GB RAM, no GPU, amd64 and arm64.
- Models: a few candidates per function, small to medium; their licences recorded (an escalation item before any is adopted).
- Runtime: in Docker; no model provider account.
- Data: `testdata/anne` only; no personal data.

## Thresholds (owner)

- [ ] H1: download size, time to first answer.
- [ ] H2: share of facts found, invented facts allowed, import time.
- [ ] H3: quality on the US-02 acceptance cases, index build time, query time.
- [ ] H3: "long" means index build time, query time, or both.

## Steps

- [ ] Owner sets machine and thresholds.
- [ ] Pick candidate models and runtimes; record licences.
- [ ] Measure H1, H2, H3; results as YAML in `spikes/local-inference/README.md`.
- [ ] Report to the owner; ADR-0037 decided from the results.
