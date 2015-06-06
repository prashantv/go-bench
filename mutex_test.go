package main

import "testing"

var sv = &SyncVal{}

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
