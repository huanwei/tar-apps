package greeting

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "named", in: "TAR", want: "Hello, TAR!"},
		{name: "empty uses default", in: "", want: "Hello, world!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.in); got != tt.want {
				t.Errorf("Hello(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
