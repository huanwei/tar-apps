# Kill-switch drill

This note records the module kill-switch drill for `tar-apps` while the module
is suspended.

## Purpose

Confirm that the service can be stopped and restored quickly, and that its
absence is visible rather than silent.

## Drill steps

1. Note the baseline: `curl -s localhost:8080/healthz` returns `ok`.
2. Disengage the service by stopping the running process (or setting `ADDR`
   to an unused port and restarting, if the process must stay up).
3. Confirm the kill switch took effect: `/healthz`, `/hello` and `/farewell`
   no longer answer on the original address.
4. Restore: `go run .` (or restart the service), then re-check `/healthz`
   returns `ok` and `/hello?name=drill` returns `{"message":"Hello, drill!"}`.
5. Record the outcome and the elapsed time to restore.

## Expected result

The service is unresponsive while suspended and serves all three endpoints
again after restore, with no code changes required. Any deviation is treated
as a failed drill and followed up before the module is un-suspended.

## Notes

- The drill does not touch source, CI configuration, or stored data.
- Run `go test ./...` and `go build ./...` before and after the drill to
  confirm the workspace stays runnable.
