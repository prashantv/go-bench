package main

import (
	"fmt"
	"io/ioutil"
)

type emptyType struct{}

func (emptyType) Call() {
	fmt.Fprintln(ioutil.Discard, "1")
}

func (e *emptyType) CallPointer() {
	fmt.Fprintln(ioutil.Discard, "2")
}
