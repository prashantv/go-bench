package main

import "testing"

func BenchmarkCall(b *testing.B) {
	var e emptyType
	for i := 0; i < b.N; i++ {
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
		e.Call()
	}
}

func BenchmarkCallPointer(b *testing.B) {
	e := &emptyType{}
	for i := 0; i < b.N; i++ {
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
		e.CallPointer()
	}
}
