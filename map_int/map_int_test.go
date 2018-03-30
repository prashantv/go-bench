package main

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"testing"
)

var _read Obj

type Obj struct {
	Name string
}

func makeMap(n int) (_ map[int]Obj) {
	m := make(map[int]Obj, n)
	for i := 0; i < n; i++ {
		k := int(crc32.ChecksumIEEE([]byte(fmt.Sprint(i))))
		m[k] = Obj{"test"}
	}
	return m
}

func keys(m map[int]Obj) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func shuffle(keys []int) []int {
	return keys
}

func BenchmarkMapInt(b *testing.B) {
	for _, mapSize := range []int{100, 1000, 5000, 10000} {
		b.Run("Size-"+strconv.Itoa(mapSize), func(b *testing.B) {

			m := makeMap(mapSize)
			ks := keys(m)

			b.Run("Read", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					k := ks[i%len(ks)]
					_read = m[k]
				}
			})

			new := Obj{"new"}
			b.Run("ReadWrite", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					k := ks[i%len(ks)]
					_read = m[k]
					m[k] = new
				}
			})
		})
	}
}
