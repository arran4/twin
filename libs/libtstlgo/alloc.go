package libtstl

// Alloc returns a pointer to a new zero-value T.
func Alloc[T any]() *T {
	return new(T)
}

// AllocN allocates a zeroed slice of length n.
func AllocN[T any](n int) []T {
	if n <= 0 {
		return nil
	}
	return make([]T, n)
}

// Grow resizes the slice to newLen, copying existing data.
// It mimics realloc semantics: nil input allocates a new slice,
// newLen <= 0 returns nil.
func Grow[T any](s []T, newLen int) []T {
	if newLen <= 0 {
		return nil
	}
	if s == nil {
		return make([]T, newLen)
	}
	if newLen <= cap(s) {
		return s[:newLen]
	}
	t := make([]T, newLen)
	copy(t, s)
	return t
}

// Clone returns a copy of the provided slice.
func Clone[T any](src []T) []T {
	if src == nil {
		return nil
	}
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}
