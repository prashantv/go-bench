package main

import (
	"fmt"
	"strings"
	"testing"
)

// getHostStd uses strings.IndexByte to find the location of the colon
// and returns the string up to that point.
func getHostStd(hostPort string) string {
	if idx := strings.IndexByte(hostPort, ':'); idx > 0 {
		return hostPort[:idx]
	}
	return hostPort
}

// getHostLoop uses a for loop to find the location of the colon and
// returns the string up to that point.
func getHostLoop(hostPort string) string {
	for i := 0; i < len(hostPort); i++ {
		if hostPort[i] == ':' {
			return hostPort[:i]
		}
	}
	return hostPort
}

func BenchmarkGetHostStd(b *testing.B) {
	var host string
	for i := 0; i < b.N; i++ {
		host = getHostStd("127.0.0.1:2134")
	}
	fmt.Println(host)
}

func BenchmarkGetHostLoop(b *testing.B) {
	var host string
	for i := 0; i < b.N; i++ {
		host = getHostLoop("127.0.0.1:2134")
	}
	fmt.Println(host)
}
