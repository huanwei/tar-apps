package farewell

import "testing"

func TestFarewell(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "named", in: "TAR", want: "Goodbye, TAR!"},
		{name: "empty uses default", in: "", want: "Goodbye, world!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Farewell(tt.in); got != tt.want {
				t.Errorf("Farewell(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
