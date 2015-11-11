package main

import (
	"strings"
	"testing"
)

var stringsCompare = strings.Compare

func BenchmarkDirectMethodCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Compare("", "")
	}
}

func BenchmarkIndirectMethodCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsCompare("", "")
	}
}
