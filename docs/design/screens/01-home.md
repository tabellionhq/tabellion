# Screen 1: Home, first visit

- Status: Draft; wording to iron out
- Session: [mvp-session.md](../mvp-session.md), step 0
- Found by playing Anne's first visit, the interface answering as a person

## Content, in order

```
# Update your CV for the last time.

Record your skills once, as you go. Export a CV tailored to each job offer.
Share what you choose with whom you choose. Follow your progress, year after year.

[Import a CV]   [Start from scratch]

🖥 Local                       🔒 Safe                           🔓 Open
Runs on your machine. Your    No telemetry, no account. Only    Export everything, any time,
CV is read here, then         skill names are looked up on      in an open format.
deleted.                      wikidata.org.

[Read our principles →]   (website)

──────────
🗄 Data: Docker volume tabellion_data, project tabellion · [Export]
⚠ Demo mode: no login. Not for production.
```

## Rules found

- One promise as the headline; one paragraph saying what the software is for; then the two actions.
- Trust comes before the details: three blocks (icon, header, one sentence); the details are on the principles page of the website.
- Text uses verbs; no reassurance or selling.
- The data location is a persistent component, on every screen: where the data is, and how to export it. Docker: volume and compose project names, passed by the compose file; binary: the data path (ADR-0020).
- Demo mode shows a discreet banner on every screen.
- Icons are decorative; the headers carry the meaning.

## Open

- The paragraph names features not in the MVP (sharing, progress): show them as coming, or leave them out until they exist.
- "No telemetry" becomes a rule, in the specification or an ADR.
- GDPR on a personal machine: the household exemption, Art. 2(2)(c); to check. The hosted service holds data: the GDPR applies.
