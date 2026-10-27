// Package farewell formats user-facing farewells.
package farewell

// Farewell returns a farewell for name. An empty name defaults to "world".
func Farewell(name string) string {
	if name == "" {
		name = "world"
	}
	return "Goodbye, " + name + "!"
}
