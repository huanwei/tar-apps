// Package env reads configuration from the environment with defaults.
package env

import (
	"os"
	"strconv"
)

// String returns the value of the environment variable named key, or def when
// the variable is unset or empty.
func String(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// Int returns the integer value of the environment variable named key, or def
// when the variable is unset, empty, or not a valid integer.
func Int(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}
