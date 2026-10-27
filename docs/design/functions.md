# Functions

- Status: Draft

Functions the system provides, independent of technology. Each is traced to the user stories that require it. Functions are added as stories are analysed.

| ID | Function | In | Out | Stories |
|---|---|---|---|---|
| F01 | Extract statements from a document | document | statements (occupation, period, terms), each with its excerpt; plugins: CV extraction, term matching (specification, Section 7) | US-01 |
| F02 | Find the skill meant by words | words, any language; optional hint | ranked candidates (ESCO, Wikidata): label, description, source, why; or nothing found | US-01, US-05 |
| F03 | Propose facts to a record | person, facts, evidence | facts with status proposed | US-01, US-05 |
| F04 | Review proposed facts | proposed facts | facts augmented, filtered, validated by the person | US-01, US-05 |
| F05 | Find competencies in a record related to a text | text, record (or a view of it) | related competencies with level, date, why | US-02, US-04, US-06, US-07 |
| F06 | Write a view | owner's words or picks, grantee | view (rules, allow list, deny list, projection), confirmed by the owner | US-03, US-06 |
| F07 | Read a record through a view | grantee, view | visible facts, hidden fields removed | US-03, US-04, US-06, US-07 |
| F10 | Serve a view to another system | view, reader | the view's facts in the exchange format | US-07 |
