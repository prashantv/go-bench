package main

import (
	"fmt"
	"testing"
)

func BenchmarkMod(b *testing.B) {
	var got int
	for i := 0; i < b.N; i++ {
		got = i % 2
	}
	fmt.Println(got)
}

func BenchmarkAnd(b *testing.B) {
	var got int
	for i := 0; i < b.N; i++ {
		got = i & 1
	}
	fmt.Println(got)
}
