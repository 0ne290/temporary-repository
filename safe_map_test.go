package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSafeMap(t *testing.T) {
	const year = 1945 // окончание Второй мировой войны
	sm := NewSafeMap()

	rand.Seed(time.Now().UnixNano())

	wg := sync.WaitGroup{}

	// 4 горутины
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < year; j++ {
				key := rand.Intn(year) + 1 // случайный ключ от 1 до year
				sm.GetAndIncrement(key)
			}
		}()
	}

	wg.Wait()

	data := sm.GetData()

	// Проверяем, что все ключи от 1 до year == 3
	for k := 1; k <= year; k++ {
		if data[k] != 3 {
			t.Fatalf("key %d has %d, expected 3", k, data[k])
		}
	}

	if sm.GetHitCount() != year*3 {
		t.Fatalf("expected hitCount=%d, got %d", year*3, sm.GetHitCount())
	}

	if sm.GetAddCount() != year {
		t.Fatalf("expected addCount=%d, got %d", year, sm.GetAddCount())
	}
}
