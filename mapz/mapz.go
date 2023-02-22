package mapz

import "golang.org/x/exp/constraints"

// MinKey returns the lowest key of the map m.
// It panics when m is empty.
func MinKey[M ~map[K]V, K constraints.Ordered, V any](m M) K {
	if len(m) == 0 {
		panic("MinKey: map is empty")
	}
	var best K
	for k := range m {
		if best > k {
			best = k
		}
	}
	return best
}

// MaxKey returns the highest key of the map m.
// It panics when m is empty.
func MaxKey[M ~map[K]V, K constraints.Ordered, V any](m M) K {
	if len(m) == 0 {
		panic("MaxKey: map is empty")
	}
	var best K
	for k := range m {
		if best < k {
			best = k
		}
	}
	return best
}
