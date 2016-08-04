package main

import (
	"fmt"
	"testing"
)

var val = "value"
var _resultStr string

func BenchmarkStringAppendSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_resultStr = fmt.Sprintf("%s.appended", val)
	}
}

func BenchmarkStringAppendAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_resultStr = val + ".appended"
	}
}
