# Screen 4: Check what the import found

- Status: Draft; wording to iron out
- Session: [mvp-session.md](../mvp-session.md), step 3; story [US-09](../user-stories/US-09-review-imported-facts.md)
- Found by playing Anne's first visit, the interface answering as a person

Opens by itself when the import has read the CV: no summary screen in between.

## Panel

```
① Extract ✓ — ② Confirm ✓ — ③ Find ✓ — ④ Check

## Check what was found · Your record: 74% complete

You
Anne Martin · anne@example.org · +33 6… · Lyon · 1992-04

### Occupations
Label                                         Title                                 Kind      Dates
MSc Computer Science, Rhône University        —                                     study     2015 – 2017
Embedded software engineer, Nordwind Sensors  embedded systems software developer   work      09/2017 – 02/2021
Maintainer of mqttlite                        software developer                    pleasure  05/2019 – today
Backend developer, ACME Logistics             software developer                    work      03/2021 – today

### Skills
✅ Complete (5)        cards, by date, first to last
✏️ To complete (12)    cards, by date; cards without a start date last
⚠ Not matched (2)      cards: the term and its occupation

[Save and proceed]
```

Card, text only:

```
┌───────────────────────────┐ ┌───────────────────────────┐
│ Data processing apps      │ │ Programming               │
│ Kafka pipeline · Rust     │ │ Rust                      │
│ ACME · 2021 – today       │ │ ⚠ When?                   │
│ Advanced                  │ │                           │
└───────────────────────────┘ └───────────────────────────┘
```

- Lines: the skill (shortest ESCO label), the concept and context, the occupation and period, the level name when set.
- Nothing is shown for an empty level; a missing start date reads "When?".

## Modal: the sentence is the form

```
Rust                                                                 ✕

❝ Backend developer. Rust, Go, C. Event-driven systems. ❞
  CV_Anne_Martin_2024.pdf, line 3

I did [computer programming] on [Rust] at [no job] from [2022] to [today], at level [▾].

[More details]   [Save]
```

Level list, meaning first, name second:

```
I need guidance · Basic
I apply known solutions · Intermediate
I solve problems · Experienced
I define the problems · Advanced
I find new problems · Expert
```

## Rules found

- Our objects, in order: who you are, occupations as a list, skills as cards.
- Skills are in buckets by completion: an import fills only what the CV states, so an uncertain skill lacks a date or a level and waits in "To complete".
- Cards move between buckets by drag and drop; a click opens the modal; a series of tips explains the concepts on a first visit, and an ⓘ next to each concept explains it.
- The modal is a sentence; each blank is a real, labelled form control styled as text: type or click in place; Tab moves from one blank to the next.
- Level and completion are words, never dots; colour and icons only repeat them.
- Composition of a concept and its context are behind "More details".
- After a save, a live region says what was saved and the new completion.

## Open

- Drag and drop needs a non-drag way to move a card (WCAG 2.5.7), e.g. in the modal.
- Drag and drop, the modal and the tips need more script than HTMX alone: a small library to choose.
- The vocabulary of the sentence ("did", "on", "at") needs wording.
