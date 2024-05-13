package genericz

// Ternary returns `a` if `cond` is true, `b` otherwise.
func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
