package slicez

// Diff returns `a` with all elements occurring in `b` removed.
func Diff[T comparable](a, b []T) []T {
	drop := make(map[T]struct{}, len(b))
	for _, e := range b {
		drop[e] = struct{}{}
	}
	var out []T
	for _, e := range a {
		if _, ok := drop[e]; !ok {
			out = append(out, e)
		}
	}
	return out
}

// Filter returns `a` with only the elements for which the callback returned true.
func Filter[T any](a []T, cb func(a T) bool) []T {
	var out []T
	for _, e := range a {
		if cb(e) {
			out = append(out, e)
		}
	}
	return out
}

// Unique returns `a` with all duplicate elements removed. Order is preserved.
func Unique[T comparable](a []T) []T {
	seen := make(map[T]struct{}, len(a))
	var out []T
	for _, e := range a {
		if _, s := seen[e]; !s {
			out = append(out, e)
		}
		seen[e] = struct{}{}
	}
	return out
}

// Concat concatenates multiple slices together into a single slice.
// The input slices are not modified, order is preserved and the returned slice is exactly sized to hold its elements.
func Concat[T any](slices ...[]T) []T {
	n := 0
	for _, s := range slices {
		n += len(s)
	}
	ret := make([]T, 0, n)
	for _, s := range slices {
		ret = append(ret, s...)
	}
	return ret
}
