package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSafeMap(t *testing.T) {
	const year = 1945 // Окончание Второй мировой войны. Источник: любой

	sm := NewSafeMap()
	rand.Seed(time.Now().UnixNano())

	keys := make([]int, year*3)
	for i := 0; i < year; i++ {
		keys[i] = i + 1
		keys[year+i] = i + 1
		keys[2*year+i] = i + 1
	}

	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	chunkSize := len(keys) / 4
	chunks := make([][]int, 4)
	for i := 0; i < 3; i++ {
		chunks[i] = keys[i*chunkSize : (i+1)*chunkSize]
	}
	chunks[3] = keys[3*chunkSize:]

	wg := sync.WaitGroup{}

	process := func(keyList []int) {
		defer wg.Done()
		for _, key := range keyList {
			val, exists := sm.Get(key)
			if !exists {
				val = 0
			}
			sm.Set(key, val+1)
		}
	}

	wg.Add(4)
	for i := 0; i < 4; i++ {
		go process(chunks[i])
	}

	wg.Wait()

	data := sm.GetData()

	if len(data) != year {
		t.Fatalf("expected %d keys, got %d", year, len(data))
	}

	for k := 1; k <= year; k++ {
		if data[k] != 3 {
			t.Fatalf("key %d has value %d, expected 3", k, data[k])
		}
	}

	expectedHitCount := year * 3
	if sm.GetHitCount() != expectedHitCount {
		t.Fatalf("expected hitCount=%d, got %d", expectedHitCount, sm.GetHitCount())
	}

	if sm.GetAddCount() != year {
		t.Fatalf("expected addCount=%d, got %d", year, sm.GetAddCount())
	}

	t.Logf("Test passed: %d keys, each with value 3", year)
}
