package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func getRandStrings(n int) []string {
	r := rand.New(rand.NewSource(1))
	ss := make([]string, 0, n)
	unique := make(map[string]bool, n)
	for i := 0; i < n; i++ {
		s := getRandString(r)
		for unique[s] {
			s = getRandString(r)
		}

		ss = append(ss, s)
		unique[s] = true
	}

	return ss
}

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func getRandString(r *rand.Rand) string {
	const length = 10

	buf := make([]byte, length)
	for i := range buf {
		buf[i] = chars[r.Intn(len(chars))]
	}
	return string(buf)
}

func BenchmarkLookups(b *testing.B) {
	strs := getRandStrings(100000)

	for _, n := range []int{
		1000, 10000, 100000,
	} {
		m := make(map[string]int, n)
		for i := 0; i < n; i++ {
			m[strs[i]] = i
		}

		b.Run(fmt.Sprint(n), func(b *testing.B) {
			lookup(b, m, strs[:1000])
		})
	}
}

func lookup(b *testing.B, m map[string]int, strs []string) {
	for i := 0; i < b.N; i++ {
		for _, s := range strs {
			v := m[s]
			call(v)
		}
	}
}

func call(v int) {}
