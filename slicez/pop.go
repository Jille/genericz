package slicez

// MustPop returns the last element of a slice, and shrinks the slice by one.
// MustPop panics if the slice is empty.
func MustPop[T any](s *[]T) T {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

// Pop returns the last element of a slice, and shrinks the slice by one.
// Pop returns _, false iff the given slice is empty.
func Pop[T any](s *[]T) (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	return MustPop(s), true
}
