// Package farewell formats user-facing farewells.
package farewell

// Farewell returns a farewell for name. An empty name defaults to "world".
func Farewell(name string) string {
	if name == "" {
		name = "world"
	}
	return "Goodbye, " + name + "!"
}

// FarewellAll returns a farewell for each name, in order. It always returns a
// non-nil slice, so an empty or nil input yields an empty slice.
func FarewellAll(names []string) []string {
	out := make([]string, len(names))
	for i, name := range names {
		out[i] = Farewell(name)
	}
	return out
}
