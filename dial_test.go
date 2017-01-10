package main

import (
	"net"
	"testing"
	"time"
)

type addr struct {
	hp    string
	raddr *net.TCPAddr
}

func BenchmarkDial(b *testing.B) {
	tests := []struct {
		name   string
		dialFn func(a addr) (net.Conn, error)
	}{
		{
			name: "net.Dial",
			dialFn: func(a addr) (net.Conn, error) {
				return net.Dial("tcp", a.hp)
			},
		},
		{
			name: "net.DialTCP",
			dialFn: func(a addr) (net.Conn, error) {
				return net.DialTCP("tcp", nil, a.raddr)
			},
		},
		{
			name: "net.DialTimeout",
			dialFn: func(a addr) (net.Conn, error) {
				return net.DialTimeout("tcp", a.hp, time.Second)
			},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			ln, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				panic(err)
			}

			go func() {
				for {
					c, err := ln.Accept()
					if err != nil {
						return
					}

					c.Close()
				}
			}()

			hp := ln.Addr().String()
			raddr, err := net.ResolveTCPAddr("tcp", hp)
			if err != nil {
				panic(err)
			}

			a := addr{hp, raddr}
			for i := 0; i < b.N; i++ {
				c, err := tt.dialFn(a)
				if err != nil {
					b.Errorf("Dial failed: %v", err)
				} else {
					c.Close()
				}
			}
		})
	}
}
