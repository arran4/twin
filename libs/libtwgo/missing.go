package libtw

import (
	"bytes"
	"os"
	"strings"
)

// TwMissingGetenv returns the value of the environment variable name.
func TwMissingGetenv(name string) string {
	return os.Getenv(name)
}

// TwMissingMemcmp compares the first n bytes of s1 and s2.
// It returns -1, 0 or 1 like bytes.Compare and never panics on short slices.
func TwMissingMemcmp(s1, s2 []byte, n int) int {
	if n <= 0 {
		return 0
	}
	if len(s1) < n {
		n = len(s1)
	}
	if len(s2) < n {
		n = len(s2)
	}
	return bytes.Compare(s1[:n], s2[:n])
}

// TwMissingStrdup returns a duplicate of the provided string.
func TwMissingStrdup(s string) string {
	if s == "" {
		return ""
	}
	// Force a copy to mimic C's strdup semantics.
	return string([]byte(s))
}

// TwMissingStrspn returns the length of the initial segment of s
// containing only bytes from accept.
func TwMissingStrspn(s, accept string) int {
	if accept == "" {
		return 0
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if strings.IndexByte(accept, s[i]) < 0 {
			break
		}
		count++
	}
	return count
}

// TwMissingStrstr returns the substring of haystack starting from the first
// occurrence of needle. If needle is empty, haystack is returned.
func TwMissingStrstr(haystack, needle string) string {
	if needle == "" {
		return haystack
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return ""
	}
	return haystack[idx:]
}

// TwOptionStrcmp compares two option strings treating "--foo" like "-foo".
func TwOptionStrcmp(s1, s2 string) int {
	if len(s1) > 2 && strings.HasPrefix(s1, "--") {
		s1 = s1[1:]
	}
	if len(s2) > 2 && strings.HasPrefix(s2, "--") {
		s2 = s2[1:]
	}
	return strings.Compare(s1, s2)
}

// helper implementing C's strncmp semantics for ASCII strings.
func cStrncmp(s1, s2 string, n int) int {
	b1 := []byte(s1)
	b2 := []byte(s2)
	for i := 0; i < n; i++ {
		var c1, c2 byte
		if i < len(b1) {
			c1 = b1[i]
		}
		if i < len(b2) {
			c2 = b2[i]
		}
		if c1 != c2 {
			if c1 < c2 {
				return -1
			}
			return 1
		}
	}
	return 0
}

// TwOptionStrncmp compares up to n bytes of two option strings.
func TwOptionStrncmp(s1, s2 string, n int) int {
	if n > 2 && len(s1) > 2 && strings.HasPrefix(s1, "--") {
		s1 = s1[1:]
		n--
	}
	if len(s2) > 2 && strings.HasPrefix(s2, "--") {
		s2 = s2[1:]
	}
	if n <= 0 {
		return 0
	}
	return cStrncmp(s1, s2, n)
}

type Tcell uint32

type Trune uint32

const utf21Size = 0x110000

// TwTrune converts a tcell to a trune.
func TwTrune(cell Tcell) Trune {
	c := cell & 0x001FFFFF
	if c >= utf21Size {
		c -= utf21Size
	}
	return Trune(c)
}
