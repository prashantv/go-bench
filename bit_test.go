package main

import "testing"

// Package level variables to ensure compiler doesn't optimize away benchmark
var (
	_ignored1 bool
	_ignored2 bool
	_ignored3 bool
	_ignored4 bool
)

const (
	check1 = 1 << iota
	check2
	check3
	check4
)

func BenchmarkBitCheckInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := int(i)
		_ignored1 = (v & check1) != 0
		_ignored2 = (v & check2) != 0
		_ignored3 = (v & check3) != 0
		_ignored4 = (v & check4) != 0
	}
}

func BenchmarkBitCheckUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := uint8(i)
		_ignored1 = (v & check1) != 0
		_ignored2 = (v & check2) != 0
		_ignored3 = (v & check3) != 0
		_ignored4 = (v & check4) != 0
	}
}

func BenchmarkBitCheckUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := uint16(i)
		_ignored1 = (v & check1) != 0
		_ignored2 = (v & check2) != 0
		_ignored3 = (v & check3) != 0
		_ignored4 = (v & check4) != 0
	}
}

func BenchmarkBitCheckUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := uint32(i)
		_ignored1 = (v & check1) != 0
		_ignored2 = (v & check2) != 0
		_ignored3 = (v & check3) != 0
		_ignored4 = (v & check4) != 0
	}
}

func BenchmarkBitCheckUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := uint64(i)
		_ignored1 = (v & check1) != 0
		_ignored2 = (v & check2) != 0
		_ignored3 = (v & check3) != 0
		_ignored4 = (v & check4) != 0
	}
}
