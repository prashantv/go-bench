package main

import (
	"testing"
	"time"

	"golang.org/x/net/context"
)

var ctx context.Context

func BenchmarkContextWithDeadline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		ctx = newCtx
		cancel()
	}
}
