package slicez

// MustIndex is like [slices.Index], but panics instead of returning -1.
func MustIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	panic("slicez.MustIndex: element not found")
}

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
	if len(a) <= 50 {
		return unique_slice(a)
	}
	return unique_map(a)
}

func unique_slice[T comparable](a []T) []T {
	var out []T
outer:
	for _, e := range a {
		for _, s := range out {
			if e == s {
				continue outer
			}
		}
		out = append(out, e)
	}
	return out
}

func unique_map[T comparable](a []T) []T {
	var out []T
	seen := make(map[T]struct{}, len(a))
	for _, e := range a {
		if _, s := seen[e]; !s {
			out = append(out, e)
			seen[e] = struct{}{}
		}
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

// Map returns a new slice with every element from `s` converted by `fn`.
func Map[T, U any](s []T, fn func(e T) U) []U {
	ret := make([]U, len(s))
	for i, e := range s {
		ret[i] = fn(e)
	}
	return ret
}

// MakePointers creates a slice of pointers and allocates the entries.
//
// All Ts are allocated in a single batch, and won't be garbage collected until the entire batch is no longer referenced.
func MakePointers[T any](n int) []*T {
	data := make([]T, n)
	ptrs := make([]*T, n)
	for i := range data {
		ptrs[i] = &data[i]
	}
	return ptrs
}
