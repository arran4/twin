package libltdl

import (
	"path/filepath"
	"strings"
)

// CanonicalizePath returns the cleaned absolute representation of path.
func CanonicalizePath(path string) (string, error) {
	if path == "" {
		return "", nil
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Clean(abs), nil
}

// ArgzizePath splits a search path on colon separators after cleaning each
// element.
func ArgzizePath(path string) []string {
	if path == "" {
		return nil
	}
	parts := strings.Split(path, ":")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p == "" {
			continue
		}
		if c, err := CanonicalizePath(p); err == nil {
			out = append(out, c)
		}
	}
	return out
}
