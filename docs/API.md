# API reference

The service is a minimal Go HTTP server. By default it listens on `:8080`; set
the `ADDR` environment variable to override (for example `ADDR=:9000`).

All endpoints respond to `GET` only. Unknown paths and unsupported methods
return `404 Not Found`.

| Method | Path | Query parameters | Response type |
| --- | --- | --- | --- |
| GET | `/healthz` | none | `text/plain` |
| GET | `/hello` | `name` (optional) | `application/json` |
| GET | `/farewell` | `name` (optional) | `application/json` |

## GET /healthz

Liveness probe. Returns the literal text `ok` followed by a newline. Useful for
container and load-balancer health checks.

- **Status:** `200 OK`
- **Content-Type:** `text/plain; charset=utf-8`
- **Body:** `ok\n`

Example request:

```bash
curl http://localhost:8080/healthz
```

Example response:

```
ok
```

## GET /hello

Returns a greeting. The optional `name` query parameter is interpolated into the
message; when it is omitted or empty, `world` is used.

- **Status:** `200 OK`
- **Content-Type:** `application/json`
- **Body:** `{"message":"Hello, <name>!"}`

Example request:

```bash
curl "http://localhost:8080/hello?name=TAR"
```

Example response:

```json
{"message":"Hello, TAR!"}
```

Example request using the default name:

```bash
curl http://localhost:8080/hello
```

Example response:

```json
{"message":"Hello, world!"}
```

## GET /farewell

Returns a farewell. The optional `name` query parameter is interpolated into the
message; when it is omitted or empty, `world` is used.

- **Status:** `200 OK`
- **Content-Type:** `application/json`
- **Body:** `{"message":"Goodbye, <name>!"}`

Example request:

```bash
curl "http://localhost:8080/farewell?name=TAR"
```

Example response:

```json
{"message":"Goodbye, TAR!"}
```

Example request using the default name:

```bash
curl http://localhost:8080/farewell
```

Example response:

```json
{"message":"Goodbye, world!"}
```

## Running locally

```bash
go run .            # listens on :8080
ADDR=:9000 go run . # listen on a custom address
```

The behavior above is covered by the tests in `main_test.go`,
`internal/greeting/greeting_test.go`, and `internal/farewell/farewell_test.go`.
