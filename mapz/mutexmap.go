package mapz

import "sync"

// MutexMap is a map protected with a mutex. Its interface closely resembles sync.Map.
// The zero value is valid.
type MutexMap[K comparable, V any] struct {
	L sync.Mutex
	M map[K]V
}

// Load returns the value stored in the map for a key. The ok result indicates whether value was found in the map.
func (m *MutexMap[K, V]) Load(key K) (V, bool) {
	m.L.Lock()
	defer m.L.Unlock()
	v, ok := m.M[key]
	return v, ok
}

// Load returns the value stored in the map for a key, or zero if no value is present. This is the same as Load() but ignoring the second result.
func (m *MutexMap[K, V]) LoadOrZero(key K) V {
	m.L.Lock()
	defer m.L.Unlock()
	return m.M[key]
}

// Store sets the value for a key.
func (m *MutexMap[K, V]) Store(key K, value V) {
	m.L.Lock()
	defer m.L.Unlock()
	if m.M == nil {
		m.M = map[K]V{}
	}
	m.M[key] = value
}

// Delete deletes the value for a key.
func (m *MutexMap[K, V]) Delete(key K) {
	m.L.Lock()
	defer m.L.Unlock()
	delete(m.M, key)
}

// LoadAndDelete deletes the value for a key, returning the previous value if any. The second result reports whether the key was present.
func (m *MutexMap[K, V]) LoadAndDelete(key K) (V, bool) {
	m.L.Lock()
	defer m.L.Unlock()
	v, ok := m.M[key]
	if ok {
		delete(m.M, key)
	}
	return v, ok
}

// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value. The loaded result is true if the value was loaded, false if stored.
func (m *MutexMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	m.L.Lock()
	defer m.L.Unlock()
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
func (m *MutexMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	m.L.Lock()
	defer m.L.Unlock()
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
func (m *MutexMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	m.L.Lock()
	defer m.L.Unlock()
	v, ok := m.M[key]
	if ok && v == old {
		delete(m.M, key)
		return true
	}
	return false
}

// CompareAndSwap swaps the old and new values for key if the value stored in the map is equal to old.
func (m *MutexMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	m.L.Lock()
	defer m.L.Unlock()
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
func MutexMapCompareAndDelete[K, V comparable](m *MutexMap[K, V], key K, old V) (deleted bool) {
	m.L.Lock()
	defer m.L.Unlock()
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
func MutexMapCompareAndSwap[K, V comparable](m *MutexMap[K, V], key K, old, new V) bool {
	m.L.Lock()
	defer m.L.Unlock()
	if m.M[key] == old {
		if m.M == nil {
			m.M = map[K]V{}
		}
		m.M[key] = new
		return true
	}
	return false
}

// WithLock calls f while holding the lock. f can manipulate the given map at will, but can't use m's regular functions because it's already holding the lock itself.
func (m *MutexMap[K, V]) WithLock(f func(m map[K]V)) {
	m.L.Lock()
	defer m.L.Unlock()
	if m.M == nil {
		m.M = map[K]V{}
	}
	f(m.M)
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
//
// Range does not block other methods on the receiver; even f itself may call any method on m.
//
// Range repeatedly picks up and drops the mutex so f() won't be called with the mutex held. Use WithLock if you need more performance at the cost of blocking other users.
func (m *MutexMap[K, V]) Range(f func(key K, value V) bool) {
	m.L.Lock()
	for k, v := range m.M {
		m.L.Unlock()
		if !f(k, v) {
			return
		}
		m.L.Lock()
	}
	m.L.Unlock()
}

// Len returns the number of elements in the map.
func (m *MutexMap[K, V]) Len() int {
	m.L.Lock()
	defer m.L.Unlock()
	return len(m.M)
}
