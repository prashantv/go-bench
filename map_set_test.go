package main

import "testing"

func BenchmarkMapSetEmptyStruct(b *testing.B) {
	m := make(map[int]struct{})

	for i := 0; i < b.N; i++ {
		want := false
		if i%5 == 0 {
			m[i] = struct{}{}
			want = true
		}
		if _, ok := m[i]; ok != want {
			panic("failed")
		}
	}
}

func BenchmarkMapSetBool(b *testing.B) {
	m := make(map[int]bool)

	for i := 0; i < b.N; i++ {
		want := false
		if i%5 == 0 {
			m[i] = true
			want = true
		}
		if _, ok := m[i]; ok != want {
			panic("failed")
		}
	}
}
