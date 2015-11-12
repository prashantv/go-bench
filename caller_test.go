package main

import (
	"fmt"
	"runtime"
	"testing"
)

func f5() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func f4() string {
	return f5()
}

func f3() string {
	return f4()
}

func f2() string {
	return f3()
}

func f1() string {
	return f2()
}

func BenchmarkCaller(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = f1()
	}
	fmt.Println("s", s)
}
