package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkWG(b *testing.B) {
	var wg sync.WaitGroup
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg.Add(1)
			wg.Done()
		}
	})
}

func BenchmarkAtomic(b *testing.B) {
	var n int32
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt32(&n, 1)
			atomic.AddInt32(&n, -1)
		}
	})
	if n != 0 {
		panic("fail")
	}
}
