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

func TestHelloAll(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{name: "single", in: []string{"TAR"}, want: []string{"Hello, TAR!"}},
		{
			name: "multiple preserves order",
			in:   []string{"TAR", "pilot"},
			want: []string{"Hello, TAR!", "Hello, pilot!"},
		},
		{
			name: "empty name uses default",
			in:   []string{""},
			want: []string{"Hello, world!"},
		},
		{
			name: "mixed empty and named",
			in:   []string{"TAR", "", "pilot"},
			want: []string{"Hello, TAR!", "Hello, world!", "Hello, pilot!"},
		},
		{name: "empty slice", in: []string{}, want: []string{}},
		{name: "nil slice", in: nil, want: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HelloAll(tt.in)
			if len(got) != len(tt.want) {
				t.Fatalf("HelloAll(%q) = %q, want %q", tt.in, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("HelloAll(%q)[%d] = %q, want %q", tt.in, i, got[i], tt.want[i])
				}
			}
		})
	}
}
