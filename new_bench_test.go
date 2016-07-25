package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type User struct {
	name string
	age  int
	foo  string
}

func BenchmarkNewKeyword(b *testing.B) {
	var u *User
	for i := 0; i < b.N; i++ {
		u = new(User)
	}
	fmt.Fprintln(ioutil.Discard, u)
}

func BenchmarkNewInline(b *testing.B) {
	var u *User
	for i := 0; i < b.N; i++ {
		u = &User{}
	}
	fmt.Fprintln(ioutil.Discard, u)
}
