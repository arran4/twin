package internal

// CloneMem returns a copy of src. It returns nil if src is empty.
func CloneMem(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

// CloneStr returns a copy of s using new backing memory.
func CloneStr(s string) string {
	if s == "" {
		return ""
	}
	b := make([]byte, len(s))
	copy(b, s)
	return string(b)
}
