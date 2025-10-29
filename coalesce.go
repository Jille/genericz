package genericz

// Coalesce returns the first argument that's not zero.
func Coalesce[T comparable](alternatives ...T) T {
	var zero T
	for _, v := range alternatives {
		if v != zero {
			return v
		}
	}
	return zero
}

// CoalesceSlice returns the first argument that has a len() of more than 0.
func CoalesceSlice[T ~[]V, V any](alternatives ...T) T {
	for _, v := range alternatives {
		if len(v) > 0 {
			return v
		}
	}
	var zero T
	return zero
}

// CoalesceMap returns the first argument that has a len() of more than 0.
func CoalesceMap[T ~map[K]V, K comparable, V any](alternatives ...T) T {
	for _, v := range alternatives {
		if len(v) > 0 {
			return v
		}
	}
	var zero T
	return zero
}
