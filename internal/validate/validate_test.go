package validate

import "testing"

func TestName(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{name: "simple", in: "TAR", want: "TAR"},
		{name: "trims whitespace", in: "  TAR \t\n", want: "TAR"},
		{name: "empty", in: "", wantErr: true},
		{name: "whitespace only", in: " \t\n ", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Name(tt.in)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("Name(%q) error = nil, want error", tt.in)
				}
				return
			}
			if err != nil {
				t.Fatalf("Name(%q) error = %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("Name(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
