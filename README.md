[![GoDoc](https://pkg.go.dev/badge/github.com/silveiralexf/goflat?status.svg)](https://pkg.go.dev/github.com/silveiralexf/goflat)
[![build-and-test](https://github.com/silveiralexf/goflat/actions/workflows/build-and-test.yaml/badge.svg)](https://github.com/silveiralexf/goflat/actions/workflows/build-and-test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/silveiralexf/goflat)](https://goreportcard.com/report/github.com/silveiralexf/goflat)

# goflat

A flat, dummy and simple piece of glue-code for building apps using
[Go](https://golang.org/) + [HTMX](https://htmx.org/) + [Pocketbase](https://pocketbase.io/).

Idea here is experimenting on how to quickly spin-up web applications without
going too deep through the Javascript/CSS/Vue/React rabbit-hole.

This project is mostly for studying.

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
* run:                         Run the app
* build:css:                   Buids CSS assets
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
* test:all:                    Clear tests cache and run all tests
* test:clean:                  Clear tests cache
* test:coverage:               Run tests
* test:unit:                   Test only unit tests without coverage or E2E
* test:vuln:                   Run Go Vulnerability Check
```

This project enforces conventional commits and some further quality checks using
[pre-commit](https://pre-commit.com), to install hooks locally and test,
execute as per below:

```sh
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
task: [go:fmt] go fmt ./...
/Users/felipe.silveira/go/bin/air
task: [go:tidy] go mod tidy -x
task: [run] air

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.52.3, built with Go go1.22.4

[07:55:38] building...
task: [build:css] tailwindcss -m -i "site/static/src/input.css" -o "site/static/public/out.css"

Rebuilding...

Done in 169ms.
task: [build] rm -rf bin/goflat
task: [build] go build -trimpath -o bin/goflat .
[07:55:44] running...
2024/09/02 07:55:44 Server started at http://127.0.0.1:8090
├─ REST API: http://127.0.0.1:8090/api/
└─ Admin UI: http://127.0.0.1:8090/_/
```

## Setting-up Oauth2

For authentication, client ID and secret can be created for example, as per below:

- [Setting-up OAuth 2.0](https://support.google.com/cloud/answer/6158849?hl=en#:~:text=Go%20to%20the%20Google%20Cloud%20Platform%20Console%20Credentials%20page.,to%20add%20a%20new%20secret.)

And to have it available on Pocketbase, just fill the required information at:

- [Auth Providers](http://localhost:8090/_/#/settings/auth-providers)
