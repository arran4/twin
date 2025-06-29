package libtw

import (
	"os"
	"testing"
)

func TestTwMissingGetenv(t *testing.T) {
	os.Setenv("TW_TEST_ENV", "value")
	if v := TwMissingGetenv("TW_TEST_ENV"); v != "value" {
		t.Fatalf("expected value, got %q", v)
	}
	os.Unsetenv("TW_TEST_ENV")
	if v := TwMissingGetenv("TW_TEST_ENV"); v != "" {
		t.Fatalf("expected empty string for unset var, got %q", v)
	}
}

func TestTwMissingMemcmp(t *testing.T) {
	if r := TwMissingMemcmp([]byte{1, 2, 3}, []byte{1, 2, 3}, 3); r != 0 {
		t.Fatalf("expected 0, got %d", r)
	}
	if r := TwMissingMemcmp([]byte{1, 2}, []byte{1, 3}, 2); r >= 0 {
		t.Fatalf("expected negative, got %d", r)
	}
	if r := TwMissingMemcmp(nil, nil, 5); r != 0 {
		t.Fatalf("nil compare expected 0, got %d", r)
	}
}

func TestTwMissingStrdup(t *testing.T) {
	s := TwMissingStrdup("hello")
	if s != "hello" {
		t.Fatalf("unexpected strdup result %q", s)
	}
}

func TestTwMissingStrspn(t *testing.T) {
	if n := TwMissingStrspn("abcde", "abc"); n != 3 {
		t.Fatalf("expected 3, got %d", n)
	}
	if n := TwMissingStrspn("hello", ""); n != 0 {
		t.Fatalf("expected 0 with empty accept, got %d", n)
	}
}

func TestTwMissingStrstr(t *testing.T) {
	if s := TwMissingStrstr("hello world", "lo"); s != "lo world" {
		t.Fatalf("unexpected strstr result %q", s)
	}
	if s := TwMissingStrstr("hello", "x"); s != "" {
		t.Fatalf("expected empty result, got %q", s)
	}
	if s := TwMissingStrstr("abc", ""); s != "abc" {
		t.Fatalf("empty needle should return haystack, got %q", s)
	}
}

func TestTwOptionStrcmp(t *testing.T) {
	if TwOptionStrcmp("--foo", "-foo") != 0 {
		t.Fatalf("expected equal for --foo and -foo")
	}
	if TwOptionStrcmp("--foo", "--bar") == 0 {
		t.Fatalf("unexpected equal for different options")
	}
}

func TestTwOptionStrncmp(t *testing.T) {
	if TwOptionStrncmp("--option", "-option", 7) != 0 {
		t.Fatalf("expected equal")
	}
	if TwOptionStrncmp("--abc", "--abd", 5) >= 0 {
		t.Fatalf("expected negative comparison")
	}
}

func TestTwTrune(t *testing.T) {
	if r := TwTrune(0x00320000); r != 0x10000 {
		t.Fatalf("expected 0x10000, got %#x", r)
	}
	if r := TwTrune(0x00100001); r != 0x00100001 {
		t.Fatalf("expected same value, got %#x", r)
	}
}
