package main

import (
	"math/rand"
	"sync"
	"testing"
)

func TestSafeMap(t *testing.T) {
	const year = 1945 // Окончание ВМВ

	sm := NewSafeMap()

	allKeys := make([]int, year*3)
	for i := 0; i < year; i++ {
		allKeys[i*3] = i + 1
		allKeys[i*3+1] = i + 1
		allKeys[i*3+2] = i + 1
	}

	rand.Shuffle(len(allKeys), func(i, j int) {
		allKeys[i], allKeys[j] = allKeys[j], allKeys[i]
	})

	tasks := make(chan int, len(allKeys))
	for _, key := range allKeys {
		tasks <- key
	}
	close(tasks)

	wg := sync.WaitGroup{}
	process := func() {
		defer wg.Done()
		for key := range tasks {
			sm.Update(key, func(v int) int {
				return v + 1
			})
		}
	}

	wg.Add(4)
	for i := 0; i < 4; i++ {
		go process()
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

	t.Logf("✅ Test passed: %d keys, each with value 3", year)
}
