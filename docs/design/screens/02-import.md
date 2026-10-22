# Screen 2: Import a CV

- Status: Draft; wording to iron out
- Session: [mvp-session.md](../mvp-session.md), step 2
- Found by playing Anne's first visit, the interface answering as a person

## Content, in order

```
## Upload your CV one last time.
With Tabellion, finding where you saved the last version won't be a problem anymore.

Older version? Upload it, then add what's missing in the next step.

🌐  Source: online file
    The link must be public.
    [ https://…                ] [Go]

💾  Source: local file
    [ Choose a file… ] [Go]

📝  Source: raw text
    [ text area                ] [Go]

Supported formats: pdf, docx, odt, txt.

──────────
🗄 Data: … · [Export]
⚠ Demo mode: no login. Not for production.
```

Layout: one row per source; icon column 15 %, content 85 %; each row has its own input and its own "Go".

## Rules found

- The headline carries the home page's promise ("for the last time") from one screen to the next.
- A raw text area alone is not a call to action: sources are listed, each with the input it needs.
- Copying from a PDF is not a usable path: files are read as they are.
- "One last time" was read as "no updates afterwards": a line says what to do with an older version, and points to the next step.

## Open

- MVP scope: reading pdf, docx and odt needs extraction adapters (Extraction port). docx and odt are zipped XML, readable in pure Go; PDF text extraction is harder; a scanned PDF has no text.
- Online file: the server fetches the URL; restricted to public HTTPS addresses, never internal ones (server-side request forgery). A public Google Docs link returns a page, not the file, unless its export form is used.
