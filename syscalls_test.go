package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

func BenchmarkFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.HasPrefix("a", "asd")
		strings.HasPrefix("a", "asd")
		strings.HasPrefix("a", "asd")
		strings.HasPrefix("a", "asd")
		strings.HasPrefix("a", "asd")
	}
}

func BenchmarkPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Making this os.Stdout makes it take ~1us per op.
		fmt.Fprintf(ioutil.Discard, " ")
	}
}
