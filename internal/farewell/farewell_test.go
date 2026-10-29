package farewell

import (
	"reflect"
	"testing"
)

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

func TestFarewellAll(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{name: "multiple names", in: []string{"TAR", "Go"}, want: []string{"Goodbye, TAR!", "Goodbye, Go!"}},
		{name: "empty names default individually", in: []string{"", "TAR"}, want: []string{"Goodbye, world!", "Goodbye, TAR!"}},
		{name: "single name", in: []string{"TAR"}, want: []string{"Goodbye, TAR!"}},
		{name: "empty slice", in: []string{}, want: []string{}},
		{name: "nil slice", in: nil, want: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FarewellAll(tt.in)
			if got == nil {
				t.Fatalf("FarewellAll(%v) = nil, want non-nil slice", tt.in)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FarewellAll(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
