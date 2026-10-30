# US-11: Preferences

- Status: Draft
- Actors: Anne, Thomas, Claude ([personas](../personas.md))
- Principles: [APR-0001](../../apr/APR-0001-free-and-open.md), [APR-0004](../../apr/APR-0004-international.md), [APR-0005](../../apr/APR-0005-hosted-service-possible.md), [APR-0006](../../apr/APR-0006-accessibility.md)
- Decision: [ADR-0035 preferences store](../../adr/ADR-0035-preferences-store.md)

## Story

As a person, I want the system to remember how I like to work, so that I do not set it again on each visit, each device or each plugin call.

## Scenarios

### S1: Languages

Anne reads in French and English. On first visit the UI follows her browser. She sets French, then English: labels of skills and occupations come in French, English when French is missing (specification, Section 9).

- Stored: an ordered list of BCP 47 language tags.
- Rule: the stored list prevails over `Accept-Language`, which applies only when nothing is stored (first visit, an agent, an app).

### S2: Accessibility

A person sets the UI to their needs before any screen is designed (APR-0006).

- Stored: nothing for now. The UI follows browser and OS settings (contrast, motion, text size). An entry is added when a need appears that these settings cannot express.

### S3: CV filters and aliases

Anne makes a CV for an offer (US-08). The `cv` plugin proposes the filters and aliases she used last time. She changes one; next time the change is there.

- Stored: in the `cv` kind's table, written by the server from the plugin's result (ADR-0035, way 1).
- Rule: Anne sees a kind's keys and values, and resets them all at once. Editing stays in the plugin's own flow.

### S4: Another device

Thomas uses the hosted instance from his phone and his laptop (US-10). His language is the same on both.

- Stored on the server, encrypted with his record's key (ADR-0035).
- Rule: on a shared instance, clients use the remote preferences by default, encrypted. A client may keep its own preferences locally instead (US-10). Synchronising the two is left to the client.

### S5: Export and move

Anne exports her data and moves to another server. Her preferences come with her.

- Out: a zip holding `Records.sqlite` and `Preferences.sqlite` (ADR-0035).
- Rule: the zip is a feature of this server, not of the specification. Another implementation may export differently.
- Rule: the zip holds the files as stored, encrypted; the key opens them. Nothing leaves in clear.

### S6: Consent

Anne agrees to send her CV text to a plugin outside her machine (ADR-0032). The server records her consent, dated, in Records.

- Stored: a consent table in Records (ADR-0035).
- Needs consent, dated and withdrawable:
  - calling a plugin outside the host or private network (specification, Section 4.10; ADR-0032);
  - a hosted service's terms (APR-0005);
  - giving someone access to a view (ADR-0007, ADR-0013);
  - facts proposed by others, such as signed skills (UC6; specification, Section 4.7).
- Rule: against click fatigue, a person may agree by default to a kind of consent. That agreement is itself a consent, in Records: it follows the record, not the install.
- Rule: giving access to a view is one act: the grant records the consent at its date; revoking the grant withdraws it.

### S7: An agent

Claude assesses Anne's recent commits for her (UC1). It answers in her language.

- Read: Anne's languages, by a client acting for her.
- Rule: Claude reads them as Anne's client, acting for her (UC1). A requester with a grant on a view does not read preferences. Exposing them is a specific need, close to an export, decided case by case.

## Resources this suggests

| Need | Resource | Scenarios |
|---|---|---|
| Read and change the core's preferences | `GET`, `PATCH /persons/{personId}/preferences` | S1, S4, S7 |
| Read and reset a plugin kind's preferences | `GET`, `DELETE /persons/{personId}/preferences/{kind}` | S3 |
| Carry a kind's preferences in a plugin call, and changes in its result | a `preferences` member in plugin requests and results (Section 7.1) | S3 |
| Export and import | `GET`, `PUT /persons/{personId}/export` (zip) | S5 |
| Give and withdraw consent, or agree by default to a kind | `GET`, `POST /persons/{personId}/consents`, `DELETE /consents/{consentId}`; a view grant records its own consent | S6 |

## Open

- None.
