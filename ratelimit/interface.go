package ratelimit

// RateLimiter is the interface implemented by a rate limiter.
type RateLimiter interface {
	// Take should block to make sure that the RPS is met.
	Take()
}
