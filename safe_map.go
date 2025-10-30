package main

import (
	"sync"
)

// Автор: Иванов Иван Иванович
// Год: 1945 — окончание Второй мировой войны (источник: Википедия)

// SafeMap — потокобезопасная мапа.
type SafeMap struct {
	mu             sync.Mutex
	data           map[int]int
	hitCount       int // счётчик обращений к ключам
	addCount       int // счётчик добавлений новых ключей
}

// NewSafeMap создаёт новый экземпляр SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[int]int),
	}
}

// GetAndIncrement получает значение по ключу, создаёт его при отсутствии и увеличивает на 1.
func (m *SafeMap) GetAndIncrement(key int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.hitCount++

	if _, exists := m.data[key]; !exists {
		m.data[key] = 0
		m.addCount++
	}

	m.data[key]++
}

// GetHitCount возвращает количество обращений.
func (m *SafeMap) GetHitCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.hitCount
}

// GetAddCount возвращает количество добавленных ключей.
func (m *SafeMap) GetAddCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.addCount
}

// GetData возвращает копию текущей мапы.
func (m *SafeMap) GetData() map[int]int {
	m.mu.Lock()
	defer m.mu.Unlock()

	copy := make(map[int]int, len(m.data))
	for k, v := range m.data {
		copy[k] = v
	}
	return copy
}
