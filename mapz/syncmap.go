package mapz

import "sync"

type SyncMap[K comparable, V any] struct {
	m sync.Map
}

// Load returns the value stored in the map for a key. The ok result indicates whether value was found in the map.
func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	v, ok := m.m.Load(key)
	if ok {
		return v.(V), true
	}
	var zero V
	return zero, false
}

// Load returns the value stored in the map for a key, or zero if no value is present. This is the same as Load() but ignoring the second result.
func (m *SyncMap[K, V]) LoadOrZero(key K) V {
	v, _ := m.Load(key)
	return v
}

// Store sets the value for a key.
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

// Delete deletes the value for a key.
func (m *SyncMap[K, V]) Delete(key K) {
	m.m.Delete(key)
}

// LoadAndDelete deletes the value for a key, returning the previous value if any. The second result reports whether the key was present.
func (m *SyncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	v, ok := m.m.LoadAndDelete(key)
	if ok {
		return v.(V), true
	}
	var zero V
	return zero, false
}

// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value. The loaded result is true if the value was loaded, false if stored.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, l := m.m.LoadOrStore(key, value)
	if l {
		return a.(V), true
	}
	return value, false
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's contents: no key will be visited more than once, but if the value for any key is stored or deleted concurrently (including by f), Range may reflect any mapping for that key from any point during the Range call. Range does not block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns false after a constant number of calls.
func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}
