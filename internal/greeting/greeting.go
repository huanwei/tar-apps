// Package greeting formats user-facing greetings.
package greeting

// Hello returns a greeting for name. An empty name defaults to "world".
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return "Hello, " + name + "!"
}

// HelloAll returns a greeting for each name, preserving order.
func HelloAll(names []string) []string {
	greetings := make([]string, len(names))
	for i, name := range names {
		greetings[i] = Hello(name)
	}
	return greetings
}
