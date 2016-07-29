package main

import "testing"

func BenchmarkAppendString(b *testing.B) {
	bs := make([]byte, 100)
	for i := 0; i < b.N; i++ {
		bs = bs[:0]
		bs = append(bs, "\n}"...)
	}
}

func BenchmarkAppendBytes(b *testing.B) {
	bs := make([]byte, 100)
	for i := 0; i < b.N; i++ {
		bs = bs[:0]
		bs = append(bs, '\\', '\n')
	}
}
