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
- [ ] Next quizzes on pp_rust_tracer commits
- [ ] Exercise types beyond quiz: interactive, definition, input → output
