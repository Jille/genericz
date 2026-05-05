package mapz

// MustLoad reads a key from a map and panics if it doesn't exist.
func MustLoad[K comparable, V any](m map[K]V, k K) V {
	v, ok := m[k]
	if !ok {
		panic("map entry not found")
	}
	return v
}
