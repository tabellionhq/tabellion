# Screen 5: Import epilogue

- Status: Draft; wording to iron out
- Session: [mvp-session.md](../mvp-session.md), end of step 3
- Found by playing Anne's first visit, the interface answering as a person

The import task, finished. Its next actions come from the cv-import adapter (`potentialAction`), not from the UI.

## Content

```
① Extract ✓ — ② Confirm ✓ — ③ Find ✓ — ④ Check ✓

## Your record is saved: 81% complete.
4 occupations · 17 skills · 2 terms not matched

┌─────────────────────┐  ┌─────────────────────┐  ┌─────────────────────┐
│         📄          │  │         🛠          │  │         🧭          │
│  Export a new CV    │  │  Refine your skills │  │  Explore more       │
│  Pick what fits a   │  │  Complete levels    │  │  Share a view,      │
│  job offer.         │  │  and dates. Match   │  │  follow your        │
│                     │  │  2 terms.           │  │  progress.          │
│  [ Export a CV ]    │  │  [ Refine ]         │  │  [ Explore ]        │
└─────────────────────┘  └─────────────────────┘  └─────────────────────┘
```

## Rules found

- The epilogue of an adapter offers what to do next: three columns, an icon, a title, one sentence, one call to action each.
- At 320 CSS pixels wide, the columns stack.
- Another adapter may offer other next actions.

## Open

- "Explore more" points to features not in the MVP (sharing, progress): link to the principles page or the roadmap until they exist.
