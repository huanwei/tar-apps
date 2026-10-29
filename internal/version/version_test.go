package version

import (
	"regexp"
	"testing"
)

func TestVersionIsSemver(t *testing.T) {
	semver := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	if !semver.MatchString(Version) {
		t.Errorf("Version = %q, want semantic version like 0.1.0", Version)
	}
}
