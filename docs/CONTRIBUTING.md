# Contributing

Thanks for contributing to `tar-apps`. Before opening a pull request, run the
checks below from the repository root. They are the same checks CI enforces, so
a clean local run means CI should pass.

## Prerequisites

- Go 1.22 or newer (`go.mod` pins the language version).

## Required checks

### Formatting — `gofmt`

```bash
gofmt -l .
```

This lists files that are not formatted. The output must be empty. To fix
formatting automatically:

```bash
gofmt -w .
```

### Static analysis — `go vet`

```bash
go vet ./...
```

Reports suspicious constructs such as printf mismatches and unreachable code.
It must exit successfully.

### Tests — `go test`

```bash
go test ./...
```

Runs the full test suite. All tests must pass.

## Run everything at once

```bash
gofmt -l . && go vet ./... && go test ./...
```

`gofmt -l .` prints unformatted files but exits 0, so CI wraps it as
`test -z "$(gofmt -l .)"` to fail on any output. To mirror CI exactly:

```bash
test -z "$(gofmt -l .)" && go vet ./... && go test ./... && go build ./...
```

## Build

```bash
go build ./...
```

Confirms the module compiles. CI also runs this step.

## Pull requests

- Keep changes focused and include tests for new behavior.
- Make sure the required checks above pass locally before requesting review.
- The repository is maintained in `assisted` mode: pull requests are merged by
  a reviewer, not automatically.
