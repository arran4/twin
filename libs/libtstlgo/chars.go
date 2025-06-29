package libtstl

import "bytes"

// Chars represents an immutable view on a byte slice.
type Chars struct {
	b []byte
}

// NewChars creates a view on b.
func NewChars(b []byte) Chars {
	return Chars{b}
}

// FromCString converts a NUL terminated C string to a Chars view.
func FromCString(cstr []byte) Chars {
	if i := bytes.IndexByte(cstr, 0); i >= 0 {
		return Chars{cstr[:i]}
	}
	return Chars{cstr}
}

// FromCStringMaxlen converts a possibly NUL terminated C string limited to maxLen.
func FromCStringMaxlen(cstr []byte, maxLen int) Chars {
	if len(cstr) > maxLen {
		cstr = cstr[:maxLen]
	}
	if i := bytes.IndexByte(cstr, 0); i >= 0 {
		return Chars{cstr[:i]}
	}
	return Chars{cstr}
}

// Size returns the number of bytes in the sequence.
func (c Chars) Size() int { return len(c.b) }

// Data returns the underlying byte slice.
func (c Chars) Data() []byte { return c.b }

// View returns a sub slice of the chars between start and end.
func (c Chars) View(start, end int) Chars {
	if start < 0 {
		start = 0
	}
	if end > len(c.b) {
		end = len(c.b)
	}
	if start > end {
		start = end
	}
	return Chars{c.b[start:end]}
}

// Find returns the index of substr in c or -1.
func (c Chars) Find(substr Chars) int {
	if len(substr.b) == 0 {
		return 0
	}
	return bytes.Index(c.b, substr.b)
}

// StartsWith reports whether c begins with substr.
func (c Chars) StartsWith(substr Chars) bool {
	if len(substr.b) == 0 {
		return true
	}
	if len(substr.b) > len(c.b) {
		return false
	}
	return bytes.Equal(c.b[:len(substr.b)], substr.b)
}

// String converts the chars to string.
func (c Chars) String() string { return string(c.b) }
