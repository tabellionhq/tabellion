# Screen 3: Confirm and read the CV

- Status: Draft; wording to iron out
- Session: [mvp-session.md](../mvp-session.md), step 2; story [US-01](../user-stories/US-01-import-cv.md)
- Found by playing Anne's first visit, the interface answering as a person

The import task, at its confirm stage, then at its find stage. Each stage is the task, rendered at its own URL.

## Confirm

```
① Extract ✓ — ② Confirm — ③ Find jobs and skills — ④ Check

## Is this your CV?

┌──────────────────────────────────────────────┐
│ Anne Martin — anne@example.org               │
│ Backend developer. Rust, Go, C. Event-driven │
│ systems.                                     │
│ …                                            │
└──────────────────────────────────────────────┘

[Yes, read it]   [No, try another source]
```

## Reading

```
① Extract ✓ — ② Confirm ✓ — ③ Find — ④ Check

## Reading your CV

✓ You: name, email, phone, address, date of birth
✓ Jobs: 3
◌ Skills: 7 so far
○ Names looked up on wikidata.org

Read by a language model running on this machine, in the container tabellion-model. [Details]

[Stop]
```

## Rules found

- With raw text, ① shows as skipped and the bar starts at ②.
- The extracted text stays in the task until the task ends.
- Where the CV is read, and by what, is said on the screen; details are one click away.

## Open

- Identity holds a name and one contact; the CV gives name, email, phone, address, date of birth. Picking them up is the CV adapter's job.
- The model is not chosen: the line and [Details] need its name.
