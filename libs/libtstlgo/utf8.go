package libtstl

import (
	"encoding/binary"
	"unicode/utf8"
)

// Utf8 represents a single UTF-8 encoded Unicode code point.
type Utf8 struct {
	b [4]byte
}

// NewUtf8FromRune converts r to UTF-8.
func NewUtf8FromRune(r rune) Utf8 {
	return Utf8{fromRune(r)}
}

// Rune returns the Unicode code point represented by u.
func (u Utf8) Rune() rune { return toRune(u.b) }

// Size returns the length in bytes of u.
func (u Utf8) Size() int { return toSize(u.b) }

// Bytes returns the UTF-8 byte sequence.
func (u Utf8) Bytes() Chars { return Chars{u.b[:u.Size()]} }

// Parse decodes the first UTF-8 sequence in c. It stores the
// sequence in u and returns the remaining bytes and whether the
// sequence was valid.
func (u *Utf8) Parse(c Chars) (Chars, bool) {
	src := c.b
	if len(src) == 0 {
		*u = NewUtf8FromRune('\uFFFD')
		return c, false
	}
	var x [4]byte
	x[0] = src[0]
	n := 1
	if x[0]&0xC0 == 0xC0 {
		max := len(src)
		if max > 4 {
			max = 4
		}
		for ; n < max; n++ {
			ch := src[n]
			if ch&0xC0 != 0x80 {
				break
			}
			x[n] = ch
		}
	}
	ok := valid(x)
	if ok {
		u.b = x
	} else {
		u.b = fromRune('\uFFFD')
	}
	if n > len(src) {
		n = len(src)
	}
	return Chars{src[n:]}, ok
}

// Peek returns the next UTF-8 sequence without consuming bytes.
func (u *Utf8) Peek(c Chars) (Utf8, bool) {
	var tmp Utf8
	_, ok := tmp.Parse(c)
	if ok {
		return tmp, true
	}
	return tmp, false
}

// internal helpers

func toSize(x [4]byte) int {
	if x[1] == 0 {
		return 1
	} else if x[2] == 0 {
		return 2
	} else if x[3] == 0 {
		return 3
	}
	return 4
}

func toRune(x [4]byte) rune {
	if x[1] == 0 {
		return rune(x[0])
	} else if x[2] == 0 {
		return rune(x[0]&0x1F)<<6 | rune(x[1]&0x3F)
	} else if x[3] == 0 {
		return rune(x[0]&0x0F)<<12 | rune(x[1]&0x3F)<<6 | rune(x[2]&0x3F)
	}
	return rune(x[0]&0x07)<<18 | rune(x[1]&0x3F)<<12 | rune(x[2]&0x3F)<<6 | rune(x[3]&0x3F)
}

func fromRune(r rune) (seq [4]byte) {
	if r > 0x10FFFF || r == 0xFFFE || r == 0xFFFF || (r >= 0xD800 && r <= 0xDFFF) {
		r = '\uFFFD'
	}
	n := utf8.EncodeRune(seq[:], r)
	for i := n; i < 4; i++ {
		seq[i] = 0
	}
	return
}

func valid(x [4]byte) bool {
	r := toRune(x)
	y := fromRune(r)
	return x == y
}

func less(x, y [4]byte) bool {
	return binary.BigEndian.Uint32(x[:]) < binary.BigEndian.Uint32(y[:])
}
