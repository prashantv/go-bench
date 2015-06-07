package main

import (
	"fmt"
	"testing"
)

// The following constants control the ratio of successful lookups where
// values exist, to lookups where the value does not exist.
const (
	successLookups = 1
	failureLookups = 1
)

func makeInnerMap(prefix string) map[string]interface{} {
	m := make(map[string]interface{})
	m[prefix+"1"] = 1
	m[prefix+"2"] = 2
	m[prefix+"3"] = 3
	return m
}

func buildMap() MapOfMap {
	m := make(MapOfMap)
	m["1"] = makeInnerMap("1")
	m["2"] = makeInnerMap("2")
	m["3"] = makeInnerMap("3")
	return m
}

func doTest(b *testing.B, f func(k1, k2 string) interface{}) {
	tests := []struct {
		k1  string
		k2  string
		val interface{}
	}{
		{"1", "11", 1},
		{"1", "13", 3},
		{"2", "21", 1},
		{"2", "22", 2},
		{"3", "32", 2},
		{"3", "33", 3},
	}

	for _, tt := range tests {
		for i := 0; i < successLookups; i++ {
			if got := f(tt.k1, tt.k2); got != tt.val {
				b.Errorf("Lookup[%v, %v] got %v, want %v", tt.k1, tt.k2, got, tt.val)
			}
		}
		for i := 0; i < failureLookups; i++ {
			if f(tt.k1+fmt.Sprint(i), tt.k2) != nil {
				b.Errorf("Lookup[%v, %v] got something other than nil", tt.k1, tt.k2)
			}
		}
	}
}

func BenchmarkGetWithOkCheck(b *testing.B) {
	m := buildMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doTest(b, m.GetWithOkCheck)
	}
}

func BenchmarkGetWithCheck(b *testing.B) {
	m := buildMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doTest(b, m.GetWithCheck)
	}
}

func BenchmarkGet(b *testing.B) {
	m := buildMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doTest(b, m.Get)
	}
}
