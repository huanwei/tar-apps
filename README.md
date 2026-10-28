# tar-apps

Minimal Go service used as the TAR production-pilot repository.

## Endpoints

| Method | Path | Description |
| --- | --- | --- |
| GET | `/healthz` | Liveness probe; returns `ok` |
| GET | `/hello?name=...` | Returns `{"message":"Hello, <name>!"}`; defaults to `world` |
| GET | `/farewell?name=...` | Returns `{"message":"Goodbye, <name>!"}`; defaults to `world` |

## Local checks

```bash
go test ./...
go vet ./...
go build ./...
go run .            # listens on :8080; override with ADDR
```

## CI

GitHub Actions runs `gofmt`, `go vet`, `go test ./...` and `go build ./...`
for every push to `main` and every pull request.

## TAR pilot

This is a low-risk pilot module. TAR works in `assisted` mode: it opens pull
requests and the reviewer merges them. Auto-merge is disabled, and
`sandbox.profile=host` is not treated as a security boundary.
