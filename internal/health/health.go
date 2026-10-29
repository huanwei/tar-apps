// Package health reports service liveness status.
package health

// Check returns the current health status.
func Check() string {
	return "ok"
}
