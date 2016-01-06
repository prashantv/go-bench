package main

import (
	"fmt"
	"strconv"
	"testing"
)

var res string

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = strconv.FormatInt(int64(i), 10)
	}
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = fmt.Sprint(i)
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = strconv.Itoa(i)
	}
}
