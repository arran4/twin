package libltdl

// LtStrlcpy copies src into dst respecting dstsize and returns the length of src.
func LtStrlcpy(dst []byte, src string, dstsize int) int {
	if dstsize == 0 {
		return len(src)
	}
	n := copy(dst[:dstsize-1], src)
	if n < dstsize {
		dst[n] = 0
	}
	return len(src)
}

// LtStrlcat appends src to dst respecting dstsize and returns the length
// of the string it tried to create (initial length of dst plus len(src)).
func LtStrlcat(dst []byte, src string, dstsize int) int {
	if dstsize == 0 {
		return len(src)
	}
	length := 0
	for length < dstsize && length < len(dst) {
		if dst[length] == 0 {
			break
		}
		length++
	}
	if length >= dstsize {
		return length + len(src)
	}
	n := copy(dst[length:dstsize-1], src)
	if length+n < dstsize {
		dst[length+n] = 0
	}
	return length + len(src)
}
