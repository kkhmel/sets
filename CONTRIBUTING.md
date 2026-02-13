# Contributing

Contributions are welcome!

If you'd like to add new exported APIs, please [open an issue][open-issue] describing your proposal first. Discussing API changes ahead of time makes pull request review much smoother.

## Setup

1. [Fork][fork] the repository
2. Clone your fork:

```bash
mkdir -p $GOPATH/src/github.com/kkhmel
cd $GOPATH/src/github.com/kkhmel
git clone git@github.com:your_github_username/sets.git
cd sets
git remote add upstream https://github.com/kkhmel/sets.git
git fetch upstream
```

## Making Changes

1. Create a new branch for your changes:

```bash
git checkout main
git fetch upstream
git rebase upstream/main
git checkout -b feature/amazing-feature
```

2. Make your changes.

3. Run tests and linters locally:

```bash
make lint
make test
make cover # 100% code coverage is required
```

4. Commit your changes with a [descriptive message][commit-message]:

```bash
git commit -m 'Add amazing feature'
```

5. Push to your fork:

```bash
git push origin feature/amazing-feature
```

6. Open a Pull Request via the GitHub UI.

## Quality Requirements

Your pull request will be automatically tested by GitHub Actions, which verify:

- **Compatibility**: Code must work on Go 1.23 and latest stable version
- **Dependencies**: No external dependencies allowed (stdlib only)
- **Linting**: Code must pass golangci-lint checks
- **Tests**: All tests must pass (including race detector)
- **Coverage**: 100% code coverage is required

[open-issue]: https://github.com/kkhmel/sets/issues/new
[fork]: https://github.com/kkhmel/sets/fork
[commit-message]: http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html
