package gostl

import (
	"fmt"
	"unicode/utf8"
)

// String wraps a Vector[byte] providing helpers similar to the C++ String class.
type String struct {
	Vector[byte]
}

// NewString creates an empty String with the given capacity.
func NewString(n int) *String {
	return &String{Vector: *NewVector[byte](n)}
}

// AssignString replaces the string contents with s.
func (s *String) AssignString(str string) {
	s.Vector.Assign([]byte(str))
}

// AppendString appends a string to this String.
func (s *String) AppendString(str string) {
	s.Vector.AppendSlice([]byte(str))
}

// AppendRune appends r encoded as UTF-8.
func (s *String) AppendRune(r rune) {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	s.Vector.AppendSlice(buf[:n])
}

// AppendByte appends a single byte.
func (s *String) AppendByte(b byte) { s.Vector.Append(b) }

// MakeCString ensures the underlying buffer is null terminated without
// affecting the logical length of the string.
func (s *String) MakeCString() {
	l := s.Size()
	s.Reserve(l + 1)
	s.data = s.data[:l+1]
	s.data[l] = 0
	s.data = s.data[:l]
}

// Formatf assigns the result of fmt.Sprintf to the string.
func (s *String) Formatf(format string, a ...any) {
	s.AssignString(fmt.Sprintf(format, a...))
}

// String returns the contents as a built-in string.
func (s *String) String() string { return string(s.data) }
