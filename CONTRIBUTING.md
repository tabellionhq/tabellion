# Contributing

## Licences

By contributing, you agree that your contribution is licensed under the licence of the part it changes:

| Part | Licence |
|---|---|
| Specification and documentation (`docs/`) | CC BY 4.0 |
| API contract and JSON-LD context (`api/`), test data (`testdata/`) | Apache-2.0 |
| Server, web UI, command line | AGPL-3.0 |

## Sign-off

Every commit is signed off under the [Developer Certificate of Origin](https://developercertificate.org/): you certify that you wrote the change, or have the right to submit it under the project's licences.

```shell
git commit -s
```

This adds a line with your name and email:

```
Signed-off-by: Your Name <you@example.org>
```

Commits without it are refused.

## Branches

1. Branch from `main`, one branch per change: `feat/view-preview`, `fix/task-state`.
2. Keep the branch short-lived: merged within days. Hide unfinished work behind a feature flag.
3. Open a pull request. It is merged when CI passes.

Each merge to `main` publishes a release.

## Commit messages

[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/), with a [gitmoji](https://gitmoji.dev/) after the type:

```
type(scope): :emoji: subject
```

Examples:

```
feat(views): ✨ preview the facts a view shows
fix(import): 🐛 keep the line of each excerpt
docs(spec): 📝 period overlap rule for views
```

| Type | Version change on release |
|---|---|
| `fix` | patch: 1.4.2 → 1.4.3 |
| `feat` | minor: 1.4.2 → 1.5.0 |
| `!` after the type, or a `BREAKING CHANGE:` footer | major: 1.4.2 → 2.0.0 |
| `docs`, `test`, `refactor`, `ci`, `chore` | none |

Versions follow [Semantic Versioning](https://semver.org/spec/v2.0.0.html). The specification has its own numbering, as an Internet-Draft: `draft-tabellion-curriculum-NN`.

## Changes to the rules

The specification ([docs/RFC.md](docs/RFC.md)) prevails over the API contract, and the contract over the code. A change to a rule starts in the specification.
