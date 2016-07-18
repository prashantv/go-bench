package main

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Rand interface {
	Intn(n int) int
}

func benchmarkRand(b *testing.B, r Rand) {
	b.ResetTimer()
	b.SetParallelism(20)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r.Intn(4)
			r.Intn(10)
			r.Intn(20)
			r.Intn(1000)
		}
	})
}

type globalRand struct{}

func (globalRand) Intn(n int) int {
	return rand.Intn(n)
}

type roundRobinRand struct {
	rands []*rand.Rand
	cur   int32
}

type lockedSource struct {
	lk  sync.Mutex
	src rand.Source
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

func newRoundRobinRand(count int) Rand {
	rr := &roundRobinRand{
		rands: make([]*rand.Rand, count),
	}
	for i := 0; i < count; i++ {
		rr.rands[i] = rand.New(
			&lockedSource{src: rand.NewSource(time.Now().UnixNano())},
		)
	}
	return rr
}

func (r *roundRobinRand) Intn(n int) int {
	i := atomic.AddInt32(&r.cur, 1)
	return r.rands[int(i)%len(r.rands)].Intn(n)
}

func BenchmarkGlobalRand(b *testing.B) {
	benchmarkRand(b, globalRand{})
}

func BenchmarkRoundRobinRand10(b *testing.B) {
	benchmarkRand(b, newRoundRobinRand(10))
}

func BenchmarkRoundRobinRand100(b *testing.B) {
	benchmarkRand(b, newRoundRobinRand(100))
}

func BenchmarkRandSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(4)
		rand.Intn(10)
		rand.Intn(20)
		rand.Intn(1000)
	}
}

type timeRand struct{}

// Probably not a good idea when n > 10.
func (timeRand) Intn(n int) int {
	return int(time.Now().UnixNano()) % n
}

func BenchmarkTimeNowRand(b *testing.B) {
	benchmarkRand(b, timeRand{})
}

// Get processor-local counter?
// Atomic operations for processors (faster!)
// Fast rand would

// TODO:
// what if there was a processor-local atomic?
// basically, you get returned the processor ID.
// and you can manage the array safely.
// maybe it should always increment and return?
