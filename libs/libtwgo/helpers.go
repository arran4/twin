package libtw

// Tw_CloneMem creates a copy of src and returns it.
func Tw_CloneMem(src []byte) []byte {
	if src == nil {
		return nil
	}
	out := make([]byte, len(src))
	copy(out, src)
	return out
}

// Tw_CloneStr returns a duplicated string.
func Tw_CloneStr(s string) string {
	return string([]byte(s))
}
