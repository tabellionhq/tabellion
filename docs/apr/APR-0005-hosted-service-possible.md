# APR-0005: Hosted service possible

- Status: Accepted
- Date: 2026-10-21

## Principle

The fastest way to adoption is a hosted service, free of charge. It stays free software, but it is easier to test and run.

The design must keep this possible: we deliver subdomains, and hosted instances that wake up when called and scale to zero by default. Money and service might be acceptable this way.

Storage in the service might end up encrypted, with the key held by the service or by the user (BYOK).

Exposing views is the part that requires persistence, and therefore a paying account.

Eventually, this could be the same service that large companies deploy on premises or in their cloud.

A company can propose views to its employees. The employee manages them: accepts, edits, or removes them. The company collects the views the employee shares, and cannot modify facts against the employee's will.
