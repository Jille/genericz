package mapz

import (
	"sync"
	"sync/atomic"
)

// AppendMap is a concurrency-safe append-only map. Its interface resembles a subset of sync.Map.
// Loads on existing values only need an atomic read. Loads of non-existent and recently stored values pick up a mutex. Stores pick up a mutex.
// The zero value is valid.
type AppendMap[K comparable, V any] struct {
	fast     atomic.Pointer[map[K]V]
	l        sync.Mutex
	truth    map[K]V
	slowHits int
}

// Load returns the value stored in the map for a key. The ok result indicates whether value was found in the map.
func (m *AppendMap[K, V]) Load(key K) (V, bool) {
	if f := m.fast.Load(); f != nil {
		if v, ok := (*f)[key]; ok {
			return v, true
		}
	}
	m.l.Lock()
	if m.truth == nil {
		m.l.Unlock()
		if f := m.fast.Load(); f != nil {
			if v, ok := (*f)[key]; ok {
				return v, true
			}
		}
	} else if v, ok := m.truth[key]; ok {
		m.slowHits++
		m.considerPromotion()
		m.l.Unlock()
		return v, true
	} else {
		m.l.Unlock()
	}
	var zero V
	return zero, false
}

// LoadOrZero returns the value stored in the map for a key, or zero if no value is present. This is the same as Load() but ignoring the second result.
func (m *AppendMap[K, V]) LoadOrZero(key K) V {
	v, _ := m.Load(key)
	return v
}

// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value. The loaded result is true if the value was loaded, false if stored.
func (m *AppendMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	if f := m.fast.Load(); f != nil {
		if v, ok := (*f)[key]; ok {
			return v, true
		}
	}
	m.l.Lock()
	defer m.l.Unlock()
	if m.truth == nil {
		if f := m.fast.Load(); f != nil {
			if v, ok := (*f)[key]; ok {
				return v, true
			}
			m.truth = make(map[K]V, len(*f))
			for k, v := range *f {
				m.truth[k] = v
			}
		} else {
			m.truth = map[K]V{}
		}
	} else if v, ok := m.truth[key]; ok {
		m.slowHits++
		m.considerPromotion()
		return v, true
	}
	m.truth[key] = value
	return value, false
}

func (m *AppendMap[K, V]) considerPromotion() {
	if m.slowHits >= len(m.truth) {
		m.promote()
	}
}

func (m *AppendMap[K, V]) promote() {
	f := m.truth
	m.fast.Store(&f)
	m.truth = nil
	m.slowHits = 0
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
//
// Range does not block other methods on the receiver; even f itself may call any method on m.
//
// Range iterates over a consistent copy of the map, snapshotted before the first callback.
func (m *AppendMap[K, V]) Range(f func(key K, value V) bool) {
	m.l.Lock()
	if m.truth != nil {
		m.promote()
	}
	m.l.Unlock()
	if fm := m.fast.Load(); fm != nil {
		for k, v := range *fm {
			if !f(k, v) {
				return
			}
		}
	}
}

// Len returns the number of elements in the map.
func (m *AppendMap[K, V]) Len() int {
	m.l.Lock()
	if m.truth != nil {
		defer m.l.Unlock()
		return len(m.truth)
	}
	m.l.Unlock()
	if f := m.fast.Load(); f != nil {
		return len(*f)
	}
	return 0
}
