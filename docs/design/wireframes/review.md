# Wireframe: review proposed facts

- Status: Superseded by the screens found in role-play ([screens/](../screens/)); fidelity removed from the model
- Story: [US-09 review imported facts](../user-stories/US-09-review-imported-facts.md); UI functions U04 to U09 ([ui-functions.md](../ui-functions.md))
- Accessibility: WCAG 2.2 AA (ADR-0027)
- Data: Anne's import ([testdata/anne](../../../testdata/anne/))

Anne goes through what the import proposed, one fact at a time, at her own pace. Review is not a gate: she can leave and come back.

## Page

`GET /ui/review` — title: "Review (12 to go) · Tabellion"

```
┌──────────────────────────────────────────────────────────────────────┐
│ [Skip to the facts]                                                  │
│ Tabellion   Home · Record · Views · Import                  EN ▾     │
├──────────────────────────────────────────────────────────────────────┤
│ # Review what the import proposed                                    │
│ 12 facts to review, from cv.txt, imported 22 October 2026.           │
│ You can correct any fact later.                                      │
│                                                                      │
│ ## Backend developer, ACME Logistics · March 2021 – today            │
│ ┌──────────────────────────────────────────────────────────────────┐ │
│ │ ### Occupation: Backend developer, ACME Logistics                │ │
│ │ ESCO occupation  software developer                              │ │
│ │ Kind             work                                            │ │
│ │ Period           March 2021 – today                              │ │
│ │ Fidelity         ●●● good                                        │ │
│ │ From the CV      "03/2021 – today  Backend developer, ACME…"     │ │
│ │                  cv.txt, line 7                                  │ │
│ │ [Accept]  [Edit]  [Reject]                                       │ │
│ └──────────────────────────────────────────────────────────────────┘ │
│ ┌──────────────────────────────────────────────────────────────────┐ │
│ │ ### Develop data processing applications · Kafka pipeline ·     │ │
│ │     Rust and Go                                                  │ │
│ │ Skill            develop data processing applications            │ │
│ │ Concept          pipeline, with Apache Kafka                     │ │
│ │ Context          Rust and Go (two languages: consider splitting) │ │
│ │ Period           March 2021 – today                              │ │
│ │ Level            not set                                         │ │
│ │ Fidelity         ●●○ average                                     │ │
│ │ From the CV      "Designed and built the Kafka data pipelines    │ │
│ │                  of the order service, in Rust and Go."          │ │
│ │                  cv.txt, line 8                                  │ │
│ │ [Accept]  [Edit]  [Reject]  [Split]  [Add facts from this line]  │ │
│ └──────────────────────────────────────────────────────────────────┘ │
│ …                                                                    │
│                                                                      │
│ ## Without an occupation                                             │
│ ┌──────────────────────────────────────────────────────────────────┐ │
│ │ ### Computer programming · Rust                                  │ │
│ │ Period           start date needed                               │ │
│ │ Level            not set                                         │ │
│ │ Fidelity         ●○○ low                                         │ │
│ │ From the CV      "Backend developer. Rust, Go, C. Event-driven…" │ │
│ │                  cv.txt, line 3                                  │ │
│ │ Needs a start date before it can be accepted.                    │ │
│ │ [Complete]  [Reject]                                             │ │
│ └──────────────────────────────────────────────────────────────────┘ │
│                                                                      │
│ ## Terms not understood (2)                                          │
│ ┌──────────────────────────────────────────────────────────────────┐ │
│ │ ### Zephyr RTOS                                                  │ │
│ │ From the CV      "Firmware on Zephyr RTOS, in C, for air-quality │ │
│ │                  sensors." cv.txt, line 13                       │ │
│ │ [Find what it means]  [Leave it out]                             │ │
│ └──────────────────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────────────────────────────┘
```

Order: occupations from the most recent; in each, the occupation first, then its practices in CV order; then practices without an occupation; then terms not understood.

The buttons on a card are the actions the server gives for that fact. "Accept" is absent while a start date is missing; "Complete" (an edit) is offered instead, with the reason in text.

## Actions

| Button | Request | Result on the page |
|---|---|---|
| Accept | `PATCH` the fact, `status: accepted` | The card shrinks to one line: "Accepted: …". Focus moves to the next card. |
| Reject | `PATCH` the fact, `status: rejected` | The card shrinks to one line: "Rejected: …". Focus moves to the next card. |
| Edit, Complete | `GET` the edit form for the fact | The card becomes a form in place (below). |
| Split | `GET` the split form | The card becomes a form: one row per new fact, prefilled from the fact; the excerpt is kept. |
| Add facts from this line | `GET` the expand form | A form below the card: empty rows; the fact stays as it is. |
| Find what it means | `GET` the resolve form | A search for a skill and a concept, prefilled with the term. |
| Leave it out | none: hidden in the UI | The term is hidden; nothing is recorded. |

### Edit form

```
┌──────────────────────────────────────────────────────────────────┐
│ ### Edit: Develop data processing applications · Kafka pipeline  │
│ Skill    [develop data processing applications      ] (search)   │
│ Concept  [pipeline                ] with [Apache Kafka    ] [+]  │
│ Context  [Rust                    ]                              │
│ From     [2021-03-01]   To [          ] (empty: ongoing)         │
│ Level    ( ) not set ( ) 1 Basic ( ) 2 Intermediate              │
│          ( ) 3 Average (•) 4 Advanced ( ) 5 Expert               │
│ From the CV  "Designed and built the Kafka data pipelines…"      │
│ [Save and accept]  [Save]  [Cancel]                              │
└──────────────────────────────────────────────────────────────────┘
```

### Split form

```
┌──────────────────────────────────────────────────────────────────┐
│ ### Split: "Designed and built the Kafka data pipelines of the   │
│     order service, in Rust and Go."                              │
│ Fact 1  develop data processing applications · Kafka pipeline ·  │
│         Rust    Level [4 Advanced ▾]                [Remove]     │
│ Fact 2  develop data processing applications · Kafka pipeline ·  │
│         Go      Level [not set ▾]                   [Remove]     │
│ [Add a fact]                                                     │
│ [Replace with these facts]  [Cancel]                             │
└──────────────────────────────────────────────────────────────────┘
```

"Add a fact" and "Remove" are requests that return a row: no script of our own for the rows.

## Accessibility

| WCAG 2.2 | On this page |
|---|---|
| 1.3.1 Info and relationships | Headings: `h1` the page, `h2` an occupation or a group, `h3` a fact. Each card is an `article` labelled by its `h3`. Field names and values are a description list. The excerpt is a `blockquote` with its source. |
| 1.4.1 Use of colour | Fidelity is a word (low, average, good) next to the dots; status is a word ("Accepted", "Rejected"). |
| 1.4.3 Contrast | Text at least 4.5:1, including the dots' label and the "not set" values. |
| 1.4.10 Reflow | At 320 CSS pixels wide, each card is one column; buttons wrap. |
| 2.1.1 Keyboard | Every action is a `button` or a link; forms are plain HTML forms. |
| 2.4.1 Bypass blocks | "Skip to the facts" link first. |
| 2.4.3 Focus order | Skip link, navigation, `h1`, then the cards in order; in a card: Accept, Edit, Reject, then the other actions. |
| 2.4.6 Headings and labels | A fact's `h3` reads as a sentence: skill · concept · context. Buttons carry the fact in their accessible name, e.g. "Accept: develop data processing applications, Kafka pipeline, Rust". |
| 2.4.11 Focus not obscured | No sticky header covering the focused card. |
| 2.5.8 Target size | Buttons at least 24 by 24 CSS pixels. |
| 3.2.2 On input | Choosing a level or a date changes nothing until Save. |
| 3.3.1, 3.3.3 Errors | A refused change (problem details) is shown in the card, next to the field, in words, and focus moves to it. |
| 4.1.3 Status messages | One polite live region: "Accepted: develop data processing applications, Kafka pipeline, Rust. 11 to go." |

Focus after a swap: on Accept or Reject, the next card's `h3` (made focusable with `tabindex="-1"`); on opening a form, its first field; on Cancel, the button that opened the form. This is the only script of our own on the page, a few lines on the HTMX `afterSwap` event.

## Open

1. Direction, to redraw minimal: review all at once. The import places each fact in a bucket, "ok", "to check" or "reject", from its fidelity; Anne moves the exceptions, then sends the review in one step.
2. Accepting all the facts of an occupation (U07, not in the MVP).
