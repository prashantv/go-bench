package main

import (
	"sync"
	"testing"
)

func BenchmarkGoroutineCost(b *testing.B) {
	var wg sync.WaitGroup

	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go wg.Done()
	}
	wg.Wait()
}

func BenchmarkNoGoroutine(b *testing.B) {
	var wg sync.WaitGroup

	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		wg.Done()
	}
	wg.Wait()
}
