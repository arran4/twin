package libltdl

// LtMalloc allocates a byte slice of size n.
func LtMalloc(n int) []byte {
	if n <= 0 {
		return nil
	}
	return make([]byte, n)
}

// LtZalloc allocates a zeroed byte slice of size n.
func LtZalloc(n int) []byte {
	if n <= 0 {
		return nil
	}
	return make([]byte, n)
}

// LtRealloc resizes the provided slice to n bytes.
func LtRealloc(mem []byte, n int) []byte {
	if n <= 0 {
		return nil
	}
	if mem == nil {
		return make([]byte, n)
	}
	if n <= cap(mem) {
		return mem[:n]
	}
	newmem := make([]byte, n)
	copy(newmem, mem)
	return newmem
}

// LtMemdup returns a copy of the provided byte slice.
func LtMemdup(mem []byte) []byte {
	if mem == nil {
		return nil
	}
	newmem := make([]byte, len(mem))
	copy(newmem, mem)
	return newmem
}

// LtStrdup returns a duplicate of the provided string.
func LtStrdup(s string) string {
	return s
}
