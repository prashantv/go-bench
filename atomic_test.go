package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicInc(b *testing.B) {
	var v int32
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt32(&v, 1)
		}
	})
}

func checkZero(v int32) {
	if v != 0 {
		panic("number should be 0")
	}
}

func BenchmarkAtomicRead(b *testing.B) {
	var v int32
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			checkZero(atomic.LoadInt32(&v))
		}
	})
}

func BenchmarkNonAtomicRead(b *testing.B) {
	var v int32
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			checkZero(v)
		}
	})
}

func benchmarkReadWrite(b *testing.B, readPercent int, readFn func() int32, writeFn func(v int32)) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i == 100 {
				i = 0
			}
			if i < readPercent {
				readFn()
			} else {
				writeFn(1)
			}
		}
	})
}

func BenchmarkReadWriteAtomicRead90(b *testing.B) {
	var v int32
	benchmarkReadWrite(b, 90, func() int32 {
		return atomic.LoadInt32(&v)
	}, func(newV int32) {
		atomic.StoreInt32(&v, newV)
	})
}

func BenchmarkReadWriteLockRead90(b *testing.B) {
	var mu sync.Mutex
	var v int32
	benchmarkReadWrite(b, 90, func() int32 {
		mu.Lock()
		v := v
		mu.Unlock()
		return v
	}, func(newV int32) {
		mu.Lock()
		v = newV
		mu.Unlock()
	})
}
