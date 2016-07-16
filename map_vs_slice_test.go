package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
)

func randKey() string {
	const n = 10
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bs := make([]byte, n)
	for i := range bs {
		bs[i] = letters[rand.Intn(len(letters))]
	}
	return string(bs)
}

func pickN(keys []string, n int) []string {
	newKeys := make([]string, n)
	for i := range newKeys {
		newKeys[i] = keys[rand.Intn(len(keys))]
	}
	return newKeys
}

var v string

func benchmarkMap(b *testing.B, n int) {
	m := make(map[string]string, n)
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		k := randKey()
		m[k] = k

		keys[i] = k
	}
	keys = pickN(keys, 10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v = m["not-found"]
		for _, k := range keys {
			v = m[k]
		}
	}
	b.StopTimer()
	fmt.Fprintf(ioutil.Discard, v)
}

type entry struct {
	key, value string
}

func lookup(entries []entry, key string) string {
	for i := range entries {
		if entries[i].key == key {
			return entries[i].value
		}
	}
	return ""
}

func benchmarkSlice(b *testing.B, n int) {
	keys := make([]string, n)
	s := make([]entry, n)
	for i := range s {
		s[i].key = randKey()
		s[i].value = s[i].key
		keys[i] = s[i].key
	}
	keys = pickN(keys, 10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v = lookup(s, "not-found")
		for _, k := range keys {
			v = lookup(s, k)
		}
	}
	b.StopTimer()
	fmt.Fprintf(ioutil.Discard, v)
}

func BenchmarkMap1(b *testing.B)   { benchmarkMap(b, 1) }
func BenchmarkSlice1(b *testing.B) { benchmarkSlice(b, 1) }

func BenchmarkMap2(b *testing.B)   { benchmarkMap(b, 2) }
func BenchmarkSlice2(b *testing.B) { benchmarkSlice(b, 2) }

func BenchmarkMap4(b *testing.B)   { benchmarkMap(b, 4) }
func BenchmarkSlice4(b *testing.B) { benchmarkSlice(b, 4) }

func BenchmarkMap10(b *testing.B)   { benchmarkMap(b, 10) }
func BenchmarkSlice10(b *testing.B) { benchmarkSlice(b, 10) }

func BenchmarkMap20(b *testing.B)   { benchmarkMap(b, 20) }
func BenchmarkSlice20(b *testing.B) { benchmarkSlice(b, 20) }

func BenchmarkMap40(b *testing.B)   { benchmarkMap(b, 40) }
func BenchmarkSlice40(b *testing.B) { benchmarkSlice(b, 40) }

func BenchmarkMap100(b *testing.B)   { benchmarkMap(b, 100) }
func BenchmarkSlice100(b *testing.B) { benchmarkSlice(b, 100) }

func BenchmarkMap500(b *testing.B)   { benchmarkMap(b, 500) }
func BenchmarkSlice500(b *testing.B) { benchmarkSlice(b, 500) }
