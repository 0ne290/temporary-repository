package main

import (
	"sync"
)

// Автор: Доровской Алексей Васильевич
// Год: 1945 — окончание Второй мировой войны

type SafeMap struct {
	mu       sync.Mutex
	data     map[int]int
	hitCount int
	addCount int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[int]int),
	}
}

func (m *SafeMap) Get(key int) (int, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.hitCount++

	val, exists := m.data[key]
	return val, exists
}

func (m *SafeMap) Set(key, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.data[key]; !exists {
		m.addCount++
	}

	m.data[key] = value
}

func (m *SafeMap) GetHitCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.hitCount
}

func (m *SafeMap) GetAddCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.addCount
}

func (m *SafeMap) GetData() map[int]int {
	m.mu.Lock()
	defer m.mu.Unlock()

	copy := make(map[int]int, len(m.data))
	for k, v := range m.data {
		copy[k] = v
	}
	return copy
}
