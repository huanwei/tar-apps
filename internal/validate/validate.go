// Package validate provides input validation helpers.
package validate

import (
	"errors"
	"strings"
)

// Name trims surrounding whitespace from name and returns the result. It
// returns an error if the trimmed name is empty.
func Name(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "", errors.New("name must not be empty")
	}
	return trimmed, nil
}
