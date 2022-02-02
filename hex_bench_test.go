package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"testing"
)

func hexEncode(i uint64) string {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return hex.EncodeToString(b)
}

func sprint(i uint64) string {
	return fmt.Sprintf("%016x", i)
}

const spaces = "                "

func strconvFormat(i uint64) string {
	s := strconv.FormatUint(i, 16)
	if len(s) < 16 {
		s = spaces[:16-len(s)] + s
	}
	return s
}

func BenchmarkHexEncode(b *testing.B) {
	tests := []struct {
		msg string
		f   func(uint64) string
	}{
		{"hexEncode", hexEncode},
		{"sprintf", sprint},
		{"strconv", strconvFormat},
	}

	for _, tt := range tests {
		b.Run(tt.msg, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = tt.f(uint64(i))
			}
		})
	}
}

func TestAll(t *testing.T) {
	for i := uint64(math.MaxUint64); i > 0; i-- {
		i := i
		if i%2 == 0 {
			i = math.MaxUint64 - i
		}
		a := hexEncode(i)
		b := sprint(i)
		c := strconvFormat(i)

		if a == b && b == c {
			continue
		}

		fmt.Println("mismatch for", i, a, b, c)
		return
	}
}
