package ratelimit

import (
	"sync"
	"time"
)

type timePeriod struct {
	sync.Mutex
	last       time.Time
	sleepFor   time.Duration
	perRequest time.Duration
}

func newTimePeriod(rps int) *timePeriod {
	return &timePeriod{
		perRequest: time.Second / time.Duration(rps),
	}
}

func (t *timePeriod) Take() {
	t.Lock()
	defer t.Unlock()

	// If this is our first request, then we allow it.
	cur := time.Now()
	if t.last.IsZero() {
		t.last = cur
		return
	}

	// sleepFor calculates how much time we should sleep based on
	// the perRequest budget and how long the last request took.
	// Since the request may take longer than the budget, this number
	// can get negative, and is summed across requests.
	t.sleepFor = t.sleepFor + t.perRequest - cur.Sub(t.last)
	t.last = cur

	// We shouldn't allow sleepFor to get too negative, since it would mean that
	// a service that slowed down a lot for a short period of time would get
	// a much higher RPS following that.
	if t.sleepFor < -time.Second {
		t.sleepFor = time.Second
	}

	// If sleepFor is positive, then we should sleep now.
	if t.sleepFor > 0 {
		time.Sleep(t.sleepFor)
		t.last = cur.Add(t.sleepFor)
		t.sleepFor = 0
	}
}
