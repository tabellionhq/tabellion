# UI to API map

- Status: Draft
- Screens: [screens/](screens/); contract: [api/openapi.yaml](../../api/openapi.yaml); specification: [RFC.md](../RFC.md)

Every control of the MVP screens, mapped to an operation of the contract.

- **ok**: the operation exists.
- **change**: the operation exists and needs a change.
- **missing**: no operation yet.
- **UI**: no server call; static content or navigation.
- **open**: what the control does is not decided.

Screens 03 and 06 were drafted in the role-play and are not saved as screens yet.

## 01 Home

| Control | Operation | Status |
|---|---|---|
| Headline, paragraph, trust blocks, principles link | | UI |
| Import a CV | to screen 02 | UI |
| Start from scratch | `getPerson`, then `createOccupation`, `createPractice` | ok |
| Data location bar: volume, project or path | | missing: server information (mode, data location) |
| Data location bar: Export | | missing: export of the whole record |
| Demo banner | | missing: server information (mode) |

## 02 Import

| Control | Operation | Status |
|---|---|---|
| Raw text, Go | `createImport`, `text/plain` | ok |
| Local file, Go (pdf, docx, odt, txt) | `createImport` with a file body; extraction by the cv plugin (RFC 7.2) | change |
| Online file, Go | `createImport` with a URL; public HTTPS only | change |
| Supported formats | the media types the cv plugin declares (RFC 7.1) | missing: plugin information |

## 03 Reading and confirm

| Control | Operation | Status |
|---|---|---|
| Progress bar | `getTask`: `stages` | ok |
| Progress while waiting | `getTask`, polled | ok |
| Extracted text | `getTask`: `text` | ok |
| Yes, read it | `changeTask` `{stage: find}` | ok |
| Stop | `changeTask` `{state: abandoned}` | ok |

## 04 Check

| Control | Operation | Status |
|---|---|---|
| Progress bar, stage ④ Check | `getTask`: `stages` | change: cv-import has no check stage (extract, confirm, find) |
| Record completion | `getPerson`: `completion` | ok |
| Identity line | `listIdentities` | ok |
| Identity from the CV | term matching (RFC 7.3) | change: matching returns no identity |
| Occupations list | `listOccupations` | ok |
| Skill cards, buckets Complete and To complete | `listPractices`, grouped by `completion` | ok |
| Not matched cards | `getTask`: `result.unresolved` | ok |
| Leave out a term | | missing |
| Resolve a term into a practice | `createPractice`; the term leaves `result.unresolved` | change: nothing removes the term |
| Card, click: modal | `getPractice`, with its evidence | ok |
| Modal: skill blank | `searchSkills` | ok |
| Modal: concept blank, More details | `searchConcepts` | ok |
| Modal: occupation blank | `listOccupations` | ok |
| Modal: dates, level | `reviewPractice` | ok; date precision open |
| Modal: Save | `reviewPractice` | ok |
| Drag a card to another bucket | | open: buckets follow completion; what a drop changes is not decided |
| Save and proceed | accept the proposed facts; end the task | missing: accept many at once (or one call per fact); an action that ends the task |
| Delete a fact | | missing: required by the specification (4.1), not on any screen yet |

## 05 Import epilogue

| Control | Operation | Status |
|---|---|---|
| Progress bar, all done | `getTask` | ok |
| Completion, counts | `getPerson`, `listOccupations`, `listPractices`, `getTask` | ok |
| The three columns | the finished task's `potentialAction` | change: the task offers "Review the proposed…", not Export, Refine, Explore |
| Export a CV | to screen 06 | UI |
| Refine | `listPractices` filtered by incomplete facts | change: no completion filter |
| Explore | link to the website | UI |

## 06 Export

| Control | Operation | Status |
|---|---|---|
| ① Who: identities | `listIdentities` | ok |
| ① Add another identity | `createIdentity` | ok |
| ② What: occupation ticks, period, show levels, show excerpts, practice ticks | `createView` on entering, `replaceView` on each change | ok |
| ② Count and list, updated live | `getViewFacts` | ok |
| ② Focus on a job | | open |
| ③ Look: installed and available themes | the themes the cv plugin declares | missing |
| ③ Install a theme | | missing |
| ③ Preview | rendering by the cv plugin (RFC 7.4), HTML | missing |
| ④ Save as | preferences | missing |
| ④ File: JSON Resume | `getViewResume` | ok |
| ④ File: PDF, HTML | rendering by the cv plugin | missing |
| Build and download: progress, Stop | a cv-export task: facts selected, CV written, rendered, downloaded | missing: task kind, operation to start it, file download |
| My CVs | preferences | missing |

## Across screens

| Need | Status |
|---|---|
| Which plugins are reachable, and what a missing one prevents (RFC 4.10) | missing |
| Language of labels | ok: `language` on searches; `Accept-Language` |
| Opening the UI: one-time token, session cookie | UI adapter, not in the contract |

## Build order

Each step ends with something Anne can use.

1. **Record, no plugin.**
   - Covers: start from scratch, screen 04 without an import, record completion.
   - Operations: `getEntry`, `getPerson`, occupations, practices, identities, `searchSkills`, `searchConcepts`.
   - Contract changes first: delete a fact; server information (mode, data location) for the bar and the banner.
2. **CV out, no plugin.**
   - Covers: screen 06 steps ① and ②, and ④ as JSON Resume.
   - Operations: views, `getViewFacts`, `getViewResume`.
   - Result: the round trip works by hand, from entered facts to a JSON Resume file.
3. **Import from text, with the model plugin.**
   - Covers: screen 02 raw text, then 03, 04 and 05.
   - Operations: `createImport`, `getTask`, `changeTask`.
   - Contract changes:
     - the plugin binding (term matching)
     - the check stage
     - accepting proposed facts and ending the task
     - leaving out a term
     - the epilogue actions
   - Matching can be a stub at first, returning the facts of `testdata/anne`.
4. **Files and rendering, with the cv plugin.**
   - Covers: screen 02 file upload; screen 06 steps ③ and ④ as PDF and HTML.
   - Contract changes: plugin binding (extraction, rendering, theme lists); the cv-export task.
5. **Later.**
   - Preferences and My CVs.
   - Installing themes.
   - Import from a URL.
   - Export of the whole record.
   - Drag and drop.
   - Focus on a job.
