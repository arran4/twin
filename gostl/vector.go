package gostl

// Vector is a dynamically growable slice similar to the C++ Vector class.
type Vector[T any] struct {
	data []T
}

// NewVector allocates a Vector with the given size.
func NewVector[T any](n int) *Vector[T] {
	v := &Vector[T]{}
	if n > 0 {
		v.data = make([]T, n)
	}
	return v
}

// Data returns the underlying slice.
func (v *Vector[T]) Data() []T { return v.data }

// Size returns the number of elements stored in the vector.
func (v *Vector[T]) Size() int { return len(v.data) }

// Capacity returns the allocated capacity of the vector.
func (v *Vector[T]) Capacity() int { return cap(v.data) }

// Empty reports whether the vector is empty.
func (v *Vector[T]) Empty() bool { return len(v.data) == 0 }

// Clear removes all elements without releasing the buffer.
func (v *Vector[T]) Clear() { v.data = v.data[:0] }

// Resize changes the size of the vector, growing with zero values when needed.
func (v *Vector[T]) Resize(n int) {
	if n <= cap(v.data) {
		if n <= len(v.data) {
			v.data = v.data[:n]
		} else {
			extra := make([]T, n-len(v.data))
			v.data = append(v.data, extra...)
		}
		return
	}
	v.reserveInternal(n)
	v.data = v.data[:n]
}

func (v *Vector[T]) reserveInternal(n int) {
	cap2 := cap(v.data)
	if cap2 >= 10 {
		cap2 *= 2
	} else {
		cap2 = 10
	}
	if n > cap2 {
		cap2 = n
	}
	newData := make([]T, len(v.data), cap2)
	copy(newData, v.data)
	v.data = newData
}

// Reserve ensures that the vector has at least the given capacity.
func (v *Vector[T]) Reserve(n int) {
	if cap(v.data) < n {
		v.reserveInternal(n)
	}
}

// Append adds one element to the vector.
func (v *Vector[T]) Append(elem T) {
	v.Reserve(len(v.data) + 1)
	v.data = append(v.data, elem)
}

// AppendSlice appends all elements from src to the vector.
func (v *Vector[T]) AppendSlice(src []T) {
	if len(src) == 0 {
		return
	}
	v.Reserve(len(v.data) + len(src))
	v.data = append(v.data, src...)
}

// Assign replaces the vector's content with a copy of src.
func (v *Vector[T]) Assign(src []T) {
	v.Resize(len(src))
	copy(v.data, src)
}

// Swap exchanges the contents of two vectors.
func (v *Vector[T]) Swap(other *Vector[T]) {
	*v, *other = *other, *v
}

// Release returns the underlying slice and clears the vector.
func (v *Vector[T]) Release() []T {
	d := v.data
	v.data = nil
	return d
}
