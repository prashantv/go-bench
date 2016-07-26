package main

import (
	"reflect"
	"testing"
	"unsafe"
)

func hash(bs []byte) int {
	sum := 0
	for _, b := range bs {
		sum += int(b)
	}
	return sum
}

func BenchmarkHashStringCast(b *testing.B) {
	s := "str"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash([]byte(s))
	}
}

func BenchmarkHashStringUnsafe(b *testing.B) {
	s := "str"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash(unsafeStrToByte(s))
	}
}

func BenchmarkHashBytes(b *testing.B) {
	s := []byte("str")
	for i := 0; i < b.N; i++ {
		hash(s)
	}
}

func unsafeStrToByte(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	byteHeader := reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&byteHeader))
}
