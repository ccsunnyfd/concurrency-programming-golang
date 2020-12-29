package rwmap

import (
	"fmt"
	"sync"
)

// RWMap is
type RWMap struct {
	sync.RWMutex
	m map[int]int
}

// NewRWMap is
func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}

// Get is
func (m *RWMap) Get(key int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[key]
	return v, existed
}

// Set is
func (m *RWMap) Set(key, value int) {
	m.Lock()
	defer m.Unlock()
	m.m[key] = value
}

// Delete is
func (m *RWMap) Delete(key int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, key)
}

// Len is
func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

// Each is
func (m *RWMap) Each(f func(k, v int) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
		fmt.Printf("k: %d, v: %d\n", k, v)
	}
}
