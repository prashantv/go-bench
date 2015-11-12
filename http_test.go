package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func setupServer(t testing.TB, handlerFunc func()) string {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("net.Listen failed: %v", err)
	}

	helloWorldBytes := []byte("Hello world")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerFunc()
		w.Write(helloWorldBytes)

	})
	go http.Serve(ln, handler)

	return fmt.Sprintf("http://%v/test", ln.Addr().String())
}

func doGet(t testing.TB, client *http.Client, url string) {
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if _, err := io.Copy(ioutil.Discard, resp.Body); err != nil {
		t.Fatalf("ReadAll failed: %v", err)
	}
}

func BenchmarkHTTPCall(b *testing.B) {
	client := &http.Client{}
	url := setupServer(b, func() {})
	// Create a connection that will be reused.
	doGet(b, client, url)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doGet(b, client, url)
	}
}

func BenchmarkHTTPCallWithCtx(b *testing.B) {
	client := &http.Client{}
	url := setupServer(b, func() {
		_, cancel := context.WithTimeout(context.Background(), time.Second)
		cancel()
	})
	// Create a connection that will be reused.
	doGet(b, client, url)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, cancel := context.WithTimeout(context.Background(), time.Second)
		doGet(b, client, url)
		cancel()
	}

}
