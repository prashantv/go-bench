package main

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

// var testHostPort = "127.0.0.1:1234"
var testHostPort = "127.127.127.127:1234"

const printHost = false

// getHostStd uses strings.IndexByte to find the location of the colon
// and returns the string up to that point.
func getHostIndexByte(hostPort string) string {
	if idx := strings.IndexByte(hostPort, ':'); idx > 0 {
		return hostPort[:idx]
	}
	return hostPort
}

func getHostSplitHostPort(hostPort string) string {
	if host, _, err := net.SplitHostPort(hostPort); err == nil {
		return host
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

func BenchmarkGetHostIndexByte(b *testing.B) {
	var host string
	for i := 0; i < b.N; i++ {
		host = getHostIndexByte(testHostPort)
	}
	if printHost {
		fmt.Println(host)
	}
}

func BenchmarkGetHostSplitHostPort(b *testing.B) {
	var host string
	for i := 0; i < b.N; i++ {
		host = getHostSplitHostPort(testHostPort)
	}
	if printHost {
		fmt.Println(host)
	}
}

func BenchmarkGetHostIndexLoop(b *testing.B) {
	var host string
	for i := 0; i < b.N; i++ {
		host = getHostIndexLoop(testHostPort)
	}
	if printHost {
		fmt.Println(host)
	}
}

func BenchmarkGetHostRangeLoop(b *testing.B) {
	var host string
	for i := 0; i < b.N; i++ {
		host = getHostRangeLoop(testHostPort)
	}
	if printHost {
		fmt.Println(host)
	}
}
