package main

import "testing"

type Interface interface {
	Get() string
}

type Impl struct{}

func (i *Impl) Get() string { return "test" }

func getI(i Interface) {
	i.Get()
}

func BenchmarkInterfaceConvInLoop(b *testing.B) {
	v := &Impl{}
	for i := 0; i < b.N; i++ {
		getI(v)
	}
}

func BenchmarkInterfaceConvOnce(b *testing.B) {
	v := &Impl{}
	var vInterface Interface = v

	for i := 0; i < b.N; i++ {
		getI(vInterface)
	}
}
