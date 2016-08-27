package main

import (
	"strings"
	"testing"
)

var (
	_emptyString    = ""
	_notEmptyString = strings.Repeat("test", 1000)
	_boolResult     bool
)

func BenchmarkCompareStringEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_boolResult = (_emptyString == "")
		_boolResult = (_notEmptyString == "")
	}
}

func BenchmarkCompareStringLen0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_boolResult = (len(_emptyString) == 0)
		_boolResult = (len(_notEmptyString) == 0)
	}
}

func BenchmarkCompareStringLenGreater0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_boolResult = (len(_emptyString) > 0)
		_boolResult = (len(_notEmptyString) > 0)
	}
}

func BenchmarkCompareStringLenNot0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_boolResult = (len(_emptyString) != 0)
		_boolResult = (len(_notEmptyString) != 0)
	}
}

func BenchmarkCompareStringNotEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_boolResult = (_emptyString != "")
		_boolResult = (_notEmptyString != "")
	}
}
