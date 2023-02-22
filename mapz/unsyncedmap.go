package mapz

// UnsyncedMap is a map protected with a mutex. Its interface closely resembles sync.Map.
// It's indended as a drop-in replacement for MutexMap and SyncMap when synchronization is no longer needed (but might become so in the future).
// The zero value is valid.
type UnsyncedMap[K comparable, V any] struct {
	M map[K]V
}

// Load returns the value stored in the map for a key. The ok result indicates whether value was found in the map.
func (m *UnsyncedMap[K, V]) Load(key K) (V, bool) {
	v, ok := m.M[key]
	return v, ok
}

// Load returns the value stored in the map for a key, or zero if no value is present. This is the same as Load() but ignoring the second result.
func (m *UnsyncedMap[K, V]) LoadOrZero(key K) V {
	return m.M[key]
}

// Store sets the value for a key.
func (m *UnsyncedMap[K, V]) Store(key K, value V) {
	if m.M == nil {
		m.M = map[K]V{}
	}
	m.M[key] = value
}

// Delete deletes the value for a key.
func (m *UnsyncedMap[K, V]) Delete(key K) {
	delete(m.M, key)
}

// LoadAndDelete deletes the value for a key, returning the previous value if any. The second result reports whether the key was present.
func (m *UnsyncedMap[K, V]) LoadAndDelete(key K) (V, bool) {
	v, ok := m.M[key]
	if ok {
		delete(m.M, key)
	}
	return v, ok
}

// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value. The loaded result is true if the value was loaded, false if stored.
func (m *UnsyncedMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, ok := m.M[key]
	if ok {
		return v, true
	}
	if m.M == nil {
		m.M = map[K]V{}
	}
	m.M[key] = value
	return value, false
}

// Swap swaps the value for a key and returns the previous value if any. The loaded result reports whether the key was present.
func (m *UnsyncedMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	v, ok := m.M[key]
	if m.M == nil {
		m.M = map[K]V{}
	}
	m.M[key] = value
	return v, ok
}

/* TODO: We'd need to restrict V to be comparable, which isn't supported in Go 1.18 yet.
// CompareAndDelete deletes the entry for key if its value is equal to old. The old value must be of a comparable type.
//
// If there is no current value for key in the map, CompareAndDelete returns false.
func (m *UnsyncedMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	v, ok := m.M[key]
	if ok && v == old {
		delete(m.M, key)
		return true
	}
	return false
}

// CompareAndSwap swaps the old and new values for key if the value stored in the map is equal to old.
func (m *UnsyncedMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	if m.M[key] == old {
	if m.M == nil {
		m.M = map[K]V{}
	}
		m.M[key] = new
		return true
	}
	return false
}
*/

// CompareAndDelete deletes the entry for key if its value is equal to old.
//
// If there is no current value for key in the map, CompareAndDelete returns false.
//
// This is a function rather than a method because Go 1.18 doesn't allow restricting a method's type parameters more than the base type (yet?).
func UnsyncedMapCompareAndDelete[K, V comparable](m *UnsyncedMap[K, V], key K, old V) (deleted bool) {
	v, ok := m.M[key]
	if ok && v == old {
		delete(m.M, key)
		return true
	}
	return false
}

// CompareAndSwap swaps the old and new values for key if the value stored in the map is equal to old. The old value must be of a comparable type.
//
// This is a function rather than a method because Go 1.18 doesn't allow restricting a method's type parameters more than the base type (yet?).
func UnsyncedMapCompareAndSwap[K, V comparable](m *UnsyncedMap[K, V], key K, old, new V) bool {
	if m.M[key] == old {
		if m.M == nil {
			m.M = map[K]V{}
		}
		m.M[key] = new
		return true
	}
	return false
}

// WithLock calls f. The name is based on the MutexMap method, but UnsyncedMap has no mutex.
func (m *UnsyncedMap[K, V]) WithLock(f func(m map[K]V)) {
	if m.M == nil {
		m.M = map[K]V{}
	}
	f(m.M)
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m *UnsyncedMap[K, V]) Range(f func(key K, value V) bool) {
	for k, v := range m.M {
		if !f(k, v) {
			return
		}
	}
}

// Len returns the number of elements in the map.
func (m *UnsyncedMap[K, V]) Len() int {
	return len(m.M)
}
