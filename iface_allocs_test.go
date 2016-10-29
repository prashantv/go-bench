package main

import "testing"

var _ifaceLast interface{}

type emptyStruct struct{}

type nonemptyStruct struct{ Name string }

func BenchmarkEmptyStructInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ifaceLast = emptyStruct{}
	}
}

func BenchmarkEmptyStructPtrInterface(b *testing.B) {
	s := &emptyStruct{}
	for i := 0; i < b.N; i++ {
		_ifaceLast = s
	}
}

func BenchmarkSmallStructInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ifaceLast = nonemptyStruct{}
	}
}

func BenchmarkSmallStructPtrInterface(b *testing.B) {
	s := &nonemptyStruct{}
	for i := 0; i < b.N; i++ {
		_ifaceLast = s
	}
}
