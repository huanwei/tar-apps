// Package greeting formats user-facing greetings.
package greeting

// Hello returns a greeting for name. An empty name defaults to "world".
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return "Hello, " + name + "!"
}
