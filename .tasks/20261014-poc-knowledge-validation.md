# Task: POC knowledge validation of agent-written commits

- Date: 2026-10-14
- Model: Claude Opus 5.5 (claude-opus-5-5)
- Tool: Claude Code
- Status: Open

A developer must be able to explain code they merged from agent commits. POC: generate exercises from real commits, at the developer's level (beginner Rust), and record answers.

- [x] Draft questions on the last 10 commits of pp_rust_tracer
- [x] Drill down on `022318b` (main exit code) and record the answers
- [x] Fix quiz format: 1–2 code blocks, 3 questions easy → hard, code shown inline, no running code, no looking outside the shown function
- [x] Create repo with README, first quiz, first session
- [x] Record types: APR (owner principles), ADR (owner decisions), AgDR (agent decisions pending review)
- [x] APR-0001 free and open, in the owner's words
- [x] Benchmark open skills standards: `docs/research/skills-standards-benchmark.md`
- [x] ADR-0001: concepts identified by Wikidata ID; skill = action · concept · context
- [x] ADR-0002: skill level scale 1–5 (basic → expert)
- [x] ADR-0003: ESCO as professional skills referential, Wikidata for gaps
- [x] AgDR lifecycle: implementation decisions; agent prepares, reviewer sub-agent checks, owner validates; template
- [x] Review AgDR-0001 to AgDR-0003 with the reviewer sub-agent: 0001 accepted; 0002, 0003 withdrawn (storage not asked for, later behind a port)
- [ ] Owner: list of actions (ADR-0001 open point)
- [ ] Owner: concepts missing from Wikidata, where and under which licence (ADR-0001)
- [ ] Owner: how evidence becomes a level 1–5
- [ ] Owner: what a CV line shows (ESCO skill, Wikidata context, level)
- [ ] Owner: whether evidence (code snippets, possibly employer IP) is exported with the skills
- [ ] Check whether the ESCO–O*NET crosswalk maps skills (ADR-0003)
- [ ] Deferred: exchange stack (ELM or CTDL-ASN + Open Badges)
- [ ] (low priority, exercise) Tag quiz questions and session results with concept IDs and levels
- [ ] (low priority, exercise) Next quizzes on pp_rust_tracer commits
- [ ] (low priority, exercise) Exercise types beyond quiz: interactive, definition, input → output
