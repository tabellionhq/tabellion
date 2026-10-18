# US-05: Certification

- Status: Draft
- Actor: Anne ([personas](../personas.md))
- Catalogue: UC6 signed skills

## Story

As Anne, I want to add to my curriculum the skills I learned as part of the CKAD certification (Certified Kubernetes Application Developer), with a link to my certificate, so that whoever reads my curriculum can check it.

## Data in, data out

- In: the certification's name, its period, the skills learned (in Anne's words), a link to the certificate.
- Out: facts in Anne's record: the preparation as an occupation (study), Practices at that occupation; each with Anne as author and the certificate link as evidence.

```yaml
evidence:
  kind: certificate
  name: Certified Kubernetes Application Developer
  url: https://…
```

## Acceptance

- Skills come from Anne's words (F02).
- The certificate link is stored as evidence and returned as is; the system never checks it (ADR-0004). Readers follow the link to the issuer.
- Levels are optional: set by Anne, or empty.

## Functions required

- F02 Find the skill meant by words.
- F03 Propose facts to a record.
- F04 Review proposed facts.

## Data model impacts

- None: OccupationHeld (study), Practice, author and evidence cover it.

## Notes

A signature or a certificate proves who made a claim, not that it is true; verification is by practice (assessments). Trust is out of scope (ADR-0004).
