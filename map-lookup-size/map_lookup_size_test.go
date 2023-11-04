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
	strs := getRandStrings(1000000)
	perms := rand.Perm(len(strs))

	for _, n := range []int{
		1000, 10000, 100000, 100000,
	} {
		m := make(map[string]int, n)
		for i := 0; i < n; i++ {
			m[strs[i]] = i
		}

		b.Run(fmt.Sprint(n), func(b *testing.B) {
			lookup(b, m, strs, perms)
		})
	}
}

func lookup(b *testing.B, m map[string]int, allStrs []string, perms []int) {
	for i := 0; i < b.N; i++ {
		permIdx := i % (len(perms) - 1000)
		strs := allStrs[permIdx : permIdx+1000]
		for _, s := range strs {
			v := m[s]
			call(v)
		}
	}
}

func call(v int) {}
