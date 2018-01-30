package main

import (
	"sync"
	"testing"
)

var (
	sv       = &SyncVal{}
	mutexInt int
)

func BenchmarkPutSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sv.PutSimple(i)
	}
}

func BenchmarkPutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sv.PutDefer(i)
	}
}

func BenchmarkPutFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sv.PutFunc(i)
	}
}

func BenchmarkRLock(b *testing.B) {
	var s sync.RWMutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.RLock()
			_ = mutexInt
			s.RLock()
		}
	})
}

func BenchmarkLock(b *testing.B) {
	var s sync.Mutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Lock()
			_ = mutexInt
			s.Unlock()
		}
	})
}
