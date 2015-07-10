package main

import (
	"bytes"
	"flag"
	"io"
	"testing"
)

// The different versions:
// Indirect = uses the writePr func with a closure
// Direct = doesn't use the writePtr func
// Val suffix = uses a value receiver
// Ptr suffix = uses  a pointer receiver
// Val_Val: Creates a Val, and uses the Val method
// Val_Ptr: Creates a Val, but uses the Ptr method.

var (
	dataToWrite []byte
	numBytes    = flag.Int("numBytes", 50, "Number of bytes to write")
)

func init() {
	flag.Parse()

	const source = "hello world"
	lenSource := len(source)
	for i := 0; i < *numBytes; i++ {
		dataToWrite = append(dataToWrite, source[i%lenSource])
	}
}

type addClose struct {
	io.Writer
}

func (addClose) Close() error {
	return nil
}

func writerCloser(writer io.Writer) io.WriteCloser {
	return addClose{writer}
}

func Writer(writer io.Writer) (io.WriteCloser, error) {
	return writerCloser(writer), nil
}

func BenchmarkWriteIndirectVal_Val(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		newArgWriterVal(Writer(&buf)).WriteIndirectVal(dataToWrite)
		buf.Reset()
	}
}

func BenchmarkWriteIndirectVal_Ptr(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		newArgWriterPtr(Writer(&buf)).WriteIndirectVal(dataToWrite)
		buf.Reset()
	}
}

func BenchmarkWriteIndirectPtr(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		newArgWriterPtr(Writer(&buf)).WriteIndirectPtr(dataToWrite)
		buf.Reset()
	}
}

func BenchmarkWriteDirectVal_Val(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		newArgWriterVal(Writer(&buf)).WriteDirectVal(dataToWrite)
		buf.Reset()
	}
}

func BenchmarkWriteDirectVal_Ptr(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		newArgWriterPtr(Writer(&buf)).WriteDirectVal(dataToWrite)
		buf.Reset()
	}
}

func BenchmarkWriteDirectPtr(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		newArgWriterPtr(Writer(&buf)).WriteDirectPtr(dataToWrite)
		buf.Reset()
	}
}

func BenchmarkNoWriter(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		writer, err := Writer(&buf)
		if err == nil {
			_, err = writer.Write(dataToWrite)
			if err == nil {
				err = writer.Close()
			}
		}
	}
}
