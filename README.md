[![GoDoc](https://pkg.go.dev/badge/github.com/silveiralexf/goflat?status.svg)](https://pkg.go.dev/github.com/silveiralexf/goflat)
[![build-and-test](https://github.com/silveiralexf/goflat/actions/workflows/build-and-test.yaml/badge.svg)](https://github.com/silveiralexf/goflat/actions/workflows/build-and-test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/silveiralexf/goflat)](https://goreportcard.com/report/github.com/silveiralexf/goflat)

# goflat

A flat, dummy and simple wrapper for building web components with pure go.

Idea here is exploring [gomponents](https://github.com/maragudk/gomponents) + [HTMX](https://htmx.org/) with some CSS
on top, as a better alternative for quickly building simple front-end applications.

Mostly for studying and testing purposes.

## Local Setup

Various tooling and tasks are automated within the [Taskfile runner](https://github.com/go-task/task/).
Configurations are done inside [Taskfile.yml](Taskfile.yml) file.

A list of tasks available can be viewed with `task -l`, as shown below:

```sh
$ task -l
task: Available tasks for this project:
* build:                       Build binary.
* default:                     Lists available commands
* hooks:                       Setup git hooks locally
* list:                        Lists available commands
* precommit:                   Verifies and fix requirements for new commits
* run:                         Run a controller from your host.
* app:clean:                   Clears built files and tests
* app:run:                     Runs goflat
* docker:build:                Build docker image with manager
* docker:push:                 Push docker image with manager to local registry
* docs:changelog:              Generates rudimentary changelog
* go:download:                 Downloads dependencies and removes unused ones
* go:fmt:                      Run go fmt against code.
* go:lint:                     Run golangci-lint linter and perform fixes.
* go:tidy:                     Run Go mod tidy
* go:vet:                      Run go vet against code.
* install:air:                 Install Air (hot-reload)
* install:all:                 Install necessary tools.
* install:goimports:           Install Go Imports
* install:golangci-lint:       Install golangci-lint
* install:vuln:                Install Go Vulnerabilities Check
* test:clean:                  Clear tests cache
* test:coverage:               Run tests
* test:unit:                   Test only unit tests without coverage or E2E
* test:vuln:                   Run Go Vulnerability Check
```

This project enforces conventional commits and some further quality checks using [pre-commit](https://pre-commit.com), to install hooks locally and test, execute as per below:

```sh
$ task hooks
$ task precommit
task: [precommit] scripts/hooks/pre-commit
task: [docs:changelog] scripts/changelog.sh
trim trailing whitespace.................................................Passed
fix end of files.........................................................Passed
check for added large files..............................................Passed
go-task-installed........................................................Passed
go-fmt...................................................................Passed
go-vet...................................................................Passed
go-lint..................................................................Passed
go-tests-unit............................................................Passed
```

## Run Locally

```sh
$ task run
/Users/silveiralexf/go/bin/air
task: [app:run] air

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.52.2, built with Go go1.22.4

[13:55:44] mkdir /Users/silveiralexf/go/src/github.com/silveiralexf/goflat/tmp
[13:55:44] !exclude bin
[13:55:44] !exclude build
[13:55:44] !exclude pkg
[13:55:44] !exclude scripts
[13:55:44] !exclude tmp
[13:55:44] !exclude web
[13:55:44] building...
task: [build] go build -trimpath -o bin/goflat main.go
[13:55:48] running...
{"time":"2024-06-08T13:55:48.787446+01:00","level":"INFO","msg":"starting server","host":"localhost:3000"}
```
