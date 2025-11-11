package main

import (
	"sync"
)

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

func (m *SafeMap) Update(key int, modifier func(int) int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.hitCount++

	val, exists := m.data[key]
	if !exists {
		val = 0
		m.addCount++
	}

	m.data[key] = modifier(val)
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

	copied := make(map[int]int, len(m.data))
	for k, v := range m.data {
		copied[k] = v
	}
	return copied
}
