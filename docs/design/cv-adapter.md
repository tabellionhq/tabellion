# CV adapter contracts

The CV adapter is the `cv` plugin (ADR-0028). It reads CV documents and renders CVs. Its contracts are outside the specification; they follow the common plugin rules of the specification (Sections 4.10 and 7.1) and may become a specification of their own.

## CV extraction

Turns a document into text.

| Direction | Field | Type | Requirement |
|---|---|---|---|
| Sent | document | bytes | required |
| Sent | media type | string, e.g. application/pdf | required |
| Returned | lines | list of strings, in reading order | required |

- The plugin declares the media types it reads.
- The line numbers of the text are the positions in `lines`, from 1. Evidence refers to them. The text then goes to term matching (specification, Section 7.2).

## CV rendering

Turns a CV document into a file to read or print.

| Direction | Field | Type | Requirement |
|---|---|---|---|
| Sent | resume | JSON Resume document, built by the server from a view | required |
| Sent | theme | theme name | required |
| Sent | format | media type, e.g. text/html, application/pdf | required |
| Returned | file | bytes, of the media type asked | required |

- The plugin declares the themes and the formats it renders.
- The server builds the resume only from the facts and fields the view shows (specification, Section 4.5).

## References

- JSON Resume schema: <https://jsonresume.org/schema>
- [ADR-0028: Core and plugins as services](../adr/ADR-0028-core-and-plugins-as-services.md)
