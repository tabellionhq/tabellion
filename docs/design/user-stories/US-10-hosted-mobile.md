# US-10: Install an app, run no server

- Status: Draft
- Catalogue: UC12 use as a service
- Actor: Thomas ([personas](../personas.md))
- Principle: [APR-0005 hosted service possible](../../apr/APR-0005-hosted-service-possible.md)

## Story

As a potential user, Thomas wants to install an app so that he doesn't have to run a server.

## Rules

- The app is a client of the HTTP API, separate from the server: it has its own rules on the device (what it shows, what it stores, offline use, encryption), with Thomas's agreement.
- The server runs as a hosted instance: on its own subdomain, woken on request, scaled to zero (APR-0005).
- The server is the authority: it checks every change it receives and refuses what breaks its rules, with a problem detail (RFC 9457).
- The server publishes its rules for client implementers: the API contract, and a specification of the curriculum rules (statuses, frozen facts, view semantics).

## Data in, data out

- In: Thomas's actions in the app.
- Out: Thomas's record on his hosted instance; what the app keeps on the device is the app's choice.

## Data model impacts

- A client that works offline needs the changes since its last synchronisation: the event stream (adapters), e.g. a feed of new and changed facts.

## Open

- Sign up and sign in to the hosted instance: passkeys, OIDC (ADR-0013).
- Who builds the app: this project, or others from the published rules.
