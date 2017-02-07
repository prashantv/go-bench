package main

import "testing"

var _uuidData = [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}

type uuidStr string
type uuidArray [16]byte
type uuidSlice []byte

var (
	_uuidStr   uuidStr
	_uuidArray uuidArray
	_uuidSlice uuidSlice
)

//go:noinline
func useUUIDStr(u uuidStr) {
	_uuidStr = u
}

//go:noinline
func useUUIDArray(u uuidArray) {
	_uuidArray = u
}

//go:noinline
func useUUIDSlice(u uuidSlice) {
	_uuidSlice = u
}

func BenchmarkUseUUIDStr(b *testing.B) {
	u := uuidStr(_uuidData[:])
	for i := 0; i < b.N; i++ {
		useUUIDStr(u)
	}
}

func BenchmarkUseUUIDSlice(b *testing.B) {
	u := uuidSlice(_uuidData[:])
	for i := 0; i < b.N; i++ {
		useUUIDSlice(u)
	}
}

func BenchmarkUseUUIDArray(b *testing.B) {
	u := uuidArray(_uuidData)
	for i := 0; i < b.N; i++ {
		useUUIDArray(u)
	}
}
