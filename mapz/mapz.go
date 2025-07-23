package mapz

import (
	"sort"
	"sync"

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
		best = k
		break
	}
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
		best = k
		break
	}
	for k := range m {
		if best < k {
			best = k
		}
	}
	return best
}

// Keys gets the keys of the given map as a slice.
// It is more efficient than slices.Collect(maps.Keys()) because it preallocates the slice correctly.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	return maps.Keys(m)
}

// Values gets the values of the given map as a slice.
// It is more efficient than slices.Collect(maps.Values()) because it preallocates the slice correctly.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	return maps.Values(m)
}

// KeysSorted gets the keys of the given map, sorts them and returns them.
// KeysSorted may fail to sort correctly when sorting slices of floating-point numbers containing Not-a-number (NaN) values.
func KeysSorted[M ~map[K]V, K constraints.Ordered, V any](m M) []K {
	keys := Keys(m)
	slices.Sort(keys)
	return keys
}

// ValuesSorted gets the values of the given map, sorts them and returns them.
// ValuesSorted may fail to sort correctly when sorting slices of floating-point numbers containing Not-a-number (NaN) values.
func ValuesSorted[M ~map[K]V, K comparable, V constraints.Ordered](m M) []V {
	values := Values(m)
	slices.Sort(values)
	return values
}

// ValuesSortedByKey gets the values of the given map, sorts them by their key and returns them.
// ValuesSortedByKey may fail to sort correctly when sorting slices of floating-point numbers containing Not-a-number (NaN) values.
func ValuesSortedByKey[M ~map[K]V, K constraints.Ordered, V any](m M) []V {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	sort.Sort(keysAndValues[K, V]{keys, values})
	return values
}

type keysAndValues[K constraints.Ordered, V any] struct {
	keys   []K
	values []V
}

func (s keysAndValues[K, V]) Len() int {
	return len(s.keys)
}

func (s keysAndValues[K, V]) Less(i, j int) bool {
	return s.keys[i] < s.keys[j]
}

func (s keysAndValues[K, V]) Swap(i, j int) {
	s.keys[i], s.keys[j] = s.keys[j], s.keys[i]
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

// StoreWithLock grabs l and then does m[key] = value. Useful one-liner for in a defer.
func StoreWithLock[K comparable, V any](l sync.Locker, m map[K]V, key K, value V) {
	l.Lock()
	defer l.Unlock()
	m[key] = value
}

// DeleteWithLock grabs l and then calls delete(m, key). Useful one-liner for in a defer.
func DeleteWithLock[K comparable, V any](l sync.Locker, m map[K]V, key K) {
	l.Lock()
	defer l.Unlock()
	delete(m, key)
}
