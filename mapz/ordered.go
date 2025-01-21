//go:build go1.23

package mapz

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

// Ordered returns an iterator over key-value pairs from m in ascending order of key.
func Ordered[K cmp.Ordered, V any](m map[K]V) iter.Seq2[K, V] {
	keys := slices.Sorted(maps.Keys(m))
	return func(yield func(K, V) bool) {
		for _, k := range keys {
			v, ok := m[k]
			if !ok {
				continue
			}
			if !yield(k, v) {
				break
			}
		}
	}
}
