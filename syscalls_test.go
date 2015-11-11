package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

func BenchmarkPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Making this os.Stdout makes it take ~1us per op.
		fmt.Fprintf(ioutil.Discard, " ")
	}
}
