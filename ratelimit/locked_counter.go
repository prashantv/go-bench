package ratelimit

import (
	"sync"
	"time"
)

// lockedCounter implements rate limiting using a token bucket.
// The bucket is locked using a mutex, and uses sync.Cond to
// signal waiters.
type lockedCounter struct {
	sync.Mutex
	cond *sync.Cond
	n    int
	rps  int
}

func newLockedCounter(rps int) *lockedCounter {
	c := &lockedCounter{rps: rps}
	c.cond = sync.NewCond(c)
	go c.filler()
	return c
}

func (c *lockedCounter) filler() {
	const numPeriods = 1000

	increment := c.rps / numPeriods
	sleepFor := time.Second / numPeriods

	for i := 0; true; i++ {
		if i == numPeriods {
			i = 0
			c.Lock()
			c.n = increment
			c.Unlock()

		} else {
			c.Lock()
			c.n += increment
			c.Unlock()
		}

		c.cond.Broadcast()
		time.Sleep(sleepFor)
	}
}

func (c *lockedCounter) Take() {
	c.Lock()
	defer c.Unlock()

	if c.n > 0 {
		c.n--
		return
	}

	for c.n <= 0 {
		c.cond.Wait()
	}
	c.n--
}
