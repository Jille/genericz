//go:build go1.20

package mapz

// Swap swaps the value for a key and returns the previous value if any. The loaded result reports whether the key was present.
//
// Swap is only available from Go 1.20 as that is when Go added the sync.Map.Swap method.
func (m *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	p, l := m.m.Swap(key, value)
	if l {
		return p.(V), true
	}
	var zero V
	return zero, false
}

// CompareAndDelete deletes the entry for key if its value is equal to old. The old value must be of a comparable type.
//
// If there is no current value for key in the map, CompareAndDelete returns false (even if the old value is the nil interface value).
//
// CompareAndDelete is only available from Go 1.20 as that is when Go added the sync.Map.CompareAndDelete method.
func (m *SyncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.m.CompareAndDelete(key, old)
}

// CompareAndSwap swaps the old and new values for key if the value stored in the map is equal to old. The old value must be of a comparable type.
//
// CompareAndSwap is only available from Go 1.20 as that is when Go added the sync.Map.CompareAndSwap method.
func (m *SyncMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.m.CompareAndSwap(key, old, new)
}
