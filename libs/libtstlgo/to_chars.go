package libtstl

// Errnum matches the C++ e_errnum enumeration used in to_chars.cpp.
type Errnum int

const (
	SUCCESS Errnum = iota
	NOMEMORY
	NOTABLES
	DLERROR
	SYSERROR
	INVALERROR
)

// ToCharsResult mirrors the C++ to_chars_result structure.
type ToCharsResult struct {
	Written int
	Err     Errnum
}

func toChar(val uint64) byte {
	if val < 10 {
		return '0' + byte(val)
	}
	return 'a' + byte(val-10)
}

// ToChars converts a signed integer to its string representation in the given base.
func ToChars(dst []byte, val int64, base int) ToCharsResult {
	if val >= 0 {
		return ToCharsUnsigned(dst, uint64(val), base)
	}
	if len(dst) < 2 {
		return ToCharsResult{0, NOTABLES}
	}
	dst[0] = '-'
	res := ToCharsUnsigned(dst[1:], uint64(-val), base)
	return ToCharsResult{res.Written + 1, res.Err}
}

// ToCharsUnsigned converts an unsigned integer to its string representation in the given base.
func ToCharsUnsigned(dst []byte, val uint64, base int) ToCharsResult {
	if base < 2 || base > 36 {
		return ToCharsResult{0, INVALERROR}
	}
	if len(dst) == 0 {
		return ToCharsResult{0, NOTABLES}
	}
	if val == 0 {
		dst[0] = '0'
		return ToCharsResult{1, SUCCESS}
	}
	var buf [64]byte
	i := len(buf)
	for val != 0 && i > 0 {
		i--
		buf[i] = toChar(val % uint64(base))
		val /= uint64(base)
	}
	src := buf[i:]
	if len(dst) < len(src) {
		copy(dst, src[:len(dst)])
		return ToCharsResult{len(dst), NOTABLES}
	}
	copy(dst, src)
	if val != 0 {
		return ToCharsResult{len(src), NOTABLES}
	}
	return ToCharsResult{len(src), SUCCESS}
}

// ToCharsCopy copies src into dst.
func ToCharsCopy(dst, src []byte) ToCharsResult {
	n := len(src)
	if len(dst) < n {
		copy(dst, src[:len(dst)])
		return ToCharsResult{len(dst), NOTABLES}
	}
	copy(dst, src)
	return ToCharsResult{n, SUCCESS}
}

// ToCharsLenSigned returns the number of characters needed for the signed integer in base.
func ToCharsLenSigned(val int64, base int) int {
	if val >= 0 {
		return ToCharsLenUnsigned(uint64(val), base)
	}
	return 1 + ToCharsLenUnsigned(uint64(-val), base)
}

// ToCharsLenUnsigned returns the number of characters needed for the unsigned integer in base.
func ToCharsLenUnsigned(val uint64, base int) int {
	if base < 2 || base > 36 {
		return 0
	}
	if val == 0 {
		return 1
	}
	length := 0
	for val != 0 {
		val /= uint64(base)
		length++
	}
	return length
}
