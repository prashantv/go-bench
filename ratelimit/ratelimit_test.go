package ratelimit

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestRateLimit(t *testing.T) {
	const numGoRoutines = 5

	l := newTimePeriod(50000)

	var rps int32
	go func() {
		for {
			time.Sleep(time.Second)
			v := atomic.SwapInt32(&rps, 0)
			fmt.Println("\nRPS", v)
		}
	}()

	for i := 0; i < numGoRoutines; i++ {
		go func() {
			c := 0
			for {
				l.Take()
				atomic.AddInt32(&rps, 1)
				c++
				if c == 400 {
					fmt.Printf(".")
					c = 0
				}
			}
		}()
	}
	select {}
}
