package health

import "testing"

func TestCheck(t *testing.T) {
	if got, want := Check(), "ok"; got != want {
		t.Errorf("Check() = %q, want %q", got, want)
	}
}
