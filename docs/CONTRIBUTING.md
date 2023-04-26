# Contributing to Erebor

To start, thank you for considering contributing to my little pet project. It
means a lot that anyone is even interested enough in this roguelike to help me
develop it.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Have a Question?](#have-a-question)
- [Style Guide and Coding Conventions](#style-guide-and-coding-conventions)
- [Request a Feature](#request-a-feature)
- [File a Bug Report](#file-a-bug-report)

## Code of Conduct

All project owners and contributors will be held to the code of conduct defined
in the [Code of Conduct](CODE_OF_CONDUCT.md). Please report any violations
to [cstaaben@gmail.com](mailto:cstaaben@gmail.com).

## Have a Question?

Feel free to either open an issue, removing any template that might be present,
or email [cstaaben@gmail.com](mailto:cstaaben@gmail.com).

## Style Guide and Coding Conventions

Go has built-in tooling to format and lint code. At a minimum, these need to be
run prior to committing any code changes. The preferred linter
is [`golangci-lint`](https://github.com/golangci/golangci-lint) with its default
settings. The preferred formatter
is [`goimports-reviser`](https://github.com/incu6us/goimports-reviser) with the
following flags:

```shell
-format
-rm-unused
-set-alias
-company-prefixes=github.com/cstaaben
-project-name=github.com/cstaaben/erebor
```

## Request a Feature

More details coming soon. For the time being, open an issue and provide as much
detail as you can.

## File a Bug Report

More details coming soon. For the time being, open an issue and provide as much
detail as you can. Ideally, this would include:

- Version of Erebor
- Any error messages
- Steps to reproduce
- Expected behavior, if applicable
