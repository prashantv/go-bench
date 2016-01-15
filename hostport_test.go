package main

import (
	"fmt"
	"strings"
	"testing"
)

const printHost = false

// getHostStd uses strings.IndexByte to find the location of the colon
// and returns the string up to that point.
func getHostStd(hostPort string) string {
	if idx := strings.IndexByte(hostPort, ':'); idx > 0 {
		return hostPort[:idx]
	}
	return hostPort
}

// getHostIndexLoop uses a for loop to find the location of the colon and
// returns the string up to that point.
func getHostIndexLoop(hostPort string) string {
	for i := 0; i < len(hostPort); i++ {
		if hostPort[i] == ':' {
			return hostPort[:i]
		}
	}
	return hostPort
}

// getHostRangeLoop uses a for loop to find the location of the colon and
// returns the string up to that point.
func getHostRangeLoop(hostPort string) string {
	for i, c := range hostPort {
		if c == ':' {
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
	if printHost {
		fmt.Println(host)
	}
}

func BenchmarkGetHostIndexLoop(b *testing.B) {
	var host string
	for i := 0; i < b.N; i++ {
		host = getHostIndexLoop("127.0.0.1:2134")
	}
	if printHost {
		fmt.Println(host)
	}
}

func BenchmarkGetHostRangeLoop(b *testing.B) {
	var host string
	for i := 0; i < b.N; i++ {
		host = getHostRangeLoop("127.0.0.1:2134")
	}
	if printHost {
		fmt.Println(host)
	}
}
