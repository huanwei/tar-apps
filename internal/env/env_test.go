package env

import (
	"os"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  string
		def  string
		want string
	}{
		{name: "set overrides default", key: "ENV_TEST_STRING_SET", val: "value", def: "fallback", want: "value"},
		{name: "empty uses default", key: "ENV_TEST_STRING_EMPTY", val: "", def: "fallback", want: "fallback"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(tt.key, tt.val)
			if got := String(tt.key, tt.def); got != tt.want {
				t.Errorf("String(%q, %q) = %q, want %q", tt.key, tt.def, got, tt.want)
			}
		})
	}
}

func TestStringUnset(t *testing.T) {
	const key = "ENV_TEST_STRING_TRULY_UNSET"
	if _, ok := os.LookupEnv(key); ok {
		t.Skipf("%s is already set in this environment", key)
	}
	if got := String(key, "fallback"); got != "fallback" {
		t.Errorf("String(%q, %q) = %q, want %q", key, "fallback", got, "fallback")
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  string
		def  int
		want int
	}{
		{name: "valid overrides default", key: "ENV_TEST_INT_VALID", val: "-7", def: 42, want: -7},
		{name: "empty uses default", key: "ENV_TEST_INT_EMPTY", val: "", def: 42, want: 42},
		{name: "invalid uses default", key: "ENV_TEST_INT_INVALID", val: "not-a-number", def: 42, want: 42},
		{name: "surrounding whitespace is invalid", key: "ENV_TEST_INT_SPACE", val: " 7 ", def: 42, want: 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(tt.key, tt.val)
			if got := Int(tt.key, tt.def); got != tt.want {
				t.Errorf("Int(%q, %d) = %d, want %d", tt.key, tt.def, got, tt.want)
			}
		})
	}
}

func TestIntUnset(t *testing.T) {
	const key = "ENV_TEST_INT_TRULY_UNSET"
	if _, ok := os.LookupEnv(key); ok {
		t.Skipf("%s is already set in this environment", key)
	}
	if got := Int(key, 42); got != 42 {
		t.Errorf("Int(%q, %d) = %d, want %d", key, 42, got, 42)
	}
}
