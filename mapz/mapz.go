package mapz

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

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

// KeysSorted gets the keys of the given map, sorts them and returns them.
// KeysSorted may fail to sort correctly when sorting slices of floating-point numbers containing Not-a-number (NaN) values.
func KeysSorted[M ~map[K]V, K constraints.Ordered, V any](m M) []K {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}

// ValuesSorted gets the values of the given map, sorts them and returns them.
// ValuesSorted may fail to sort correctly when sorting slices of floating-point numbers containing Not-a-number (NaN) values.
func ValuesSorted[M ~map[K]V, K comparable, V constraints.Ordered](m M) []V {
	values := maps.Values(m)
	slices.Sort(values)
	return values
}
