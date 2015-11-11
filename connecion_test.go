package main

import (
	"io"
	"net"
	"testing"
)

const packetSize = 16

func connectToServer(t testing.TB) net.Conn {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}

	go acceptLoop(t, ln)

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("net.Dial failed: %v", err)
	}

	return conn
}

func acceptLoop(t testing.TB, ln net.Listener) {
	defer ln.Close()
	conn, err := ln.Accept()
	if err != nil {
		t.Fatalf("Accept failed: %v", err)
	}

	data := make([]byte, packetSize)
	for {
		if readBytes(t, conn, data) {
			return
		}
		if readBytes(t, conn, data) {
			return
		}
		writeBytes(t, conn, data)
		writeBytes(t, conn, data)
	}
}

func readBytes(t testing.TB, r io.Reader, data []byte) bool {
	if _, err := io.ReadFull(r, data); err != nil {
		if err == io.EOF {
			return true
		}

		t.Fatalf("Failed to read data: %v", err)
	}
	return false
}

func writeBytes(t testing.TB, w io.Writer, data []byte) {
	if _, err := w.Write(data); err != nil {
		t.Fatalf("Failed to write data: %v", err)
	}
}

// BenchmarkConnectionRoundtrip measures the latency for sending a packet and getting a response.
// The caller writes 2 packets with packetSize, the receiver reads these packets and immediately
// sends them back. There is no work performed on the reciever.
// There is no memory allocated during the send/receive.
func BenchmarkConnectionRoundtripSerial(b *testing.B) {
	conn := connectToServer(b)
	defer conn.Close()

	bs := make([]byte, packetSize)
	bs[0] = 1
	bs[packetSize-1] = 2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeBytes(b, conn, bs)
		writeBytes(b, conn, bs)
		readBytes(b, conn, bs)
		readBytes(b, conn, bs)
	}
}

func benchmarkConnectionRoundtripParallel(b *testing.B, parallelism int) {
	b.SetParallelism(parallelism)
	b.RunParallel(func(pb *testing.PB) {
		conn := connectToServer(b)
		defer conn.Close()

		bs := make([]byte, packetSize)
		bs[0] = 1
		bs[packetSize-1] = 2

		for pb.Next() {
			writeBytes(b, conn, bs)
			writeBytes(b, conn, bs)
			readBytes(b, conn, bs)
			readBytes(b, conn, bs)
		}
	})
}

func BenchmarkConnectionRoundtripParallel1(b *testing.B) {
	benchmarkConnectionRoundtripParallel(b, 1)
}

func BenchmarkConnectionRoundtripParallel2(b *testing.B) {
	benchmarkConnectionRoundtripParallel(b, 2)
}

func BenchmarkConnectionRoundtripParallel4(b *testing.B) {
	benchmarkConnectionRoundtripParallel(b, 4)
}

func BenchmarkConnectionRoundtripParallel8(b *testing.B) {
	benchmarkConnectionRoundtripParallel(b, 8)
}
