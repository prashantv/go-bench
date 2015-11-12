package main

import (
	"sync"
	"testing"
)

func doWork() {}

func BenchmarkSingleWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			doWork()
		}()

		wg.Wait()
	}
}

func BenchmarkSingleChanClose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan struct{})
		go func() {
			defer close(ch)
			doWork()
		}()

		<-ch
	}
}
