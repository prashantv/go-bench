package main

import (
	"fmt"
	"testing"
)

func BenchmarkMapNoContains(b *testing.B) {
	for _, nKeys := range []int{100, 1000, 10000, 100000} {
		b.Run(fmt.Sprint(nKeys), func(b *testing.B) {
			m := make(map[int]struct{}, nKeys)
			for i := 0; i < b.N; i++ {
				k := i % nKeys
				if _, ok := m[k]; ok {
					continue
				}
				m[k] = struct{}{}
			}
		})
	}
}

func BenchmarkMapContains(b *testing.B) {
	for _, nKeys := range []int{100, 1000, 10000, 100000} {
		b.Run(fmt.Sprint(nKeys), func(b *testing.B) {
			m := make(map[int]struct{}, nKeys)
			for i := 0; i < b.N; i++ {
				k := i % nKeys
				m[k] = struct{}{}
			}
		})
	}

}
