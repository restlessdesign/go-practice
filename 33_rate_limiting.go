package main

import (
	"fmt"
	"time"
)

func mockRequests(total_requests int, rate_limiter <-chan time.Time) {
	requests := make(chan int, total_requests)
	for i := 1; i <= total_requests; i++ {
		requests <- i
	}
	close(requests)

	for request := range requests {
		<-rate_limiter
		fmt.Println("Request", request, time.Now())
	}
}

// Basic rate limiter which throttles all requests based on a delay value
// supplied to the method (in milliseconds)
func basicRateLimiter(delay int) <-chan time.Time {
	rate_limiter := time.Tick(time.Millisecond * time.Duration(delay))
	return rate_limiter
}

func burstRateLimiter(request_threshold int, delay int) <-chan time.Time {
	rate_limiter := make(chan time.Time, request_threshold)

	// Fill up the channel’s buffer
	//
	// Remember that buffered channels will only block a sender once the buffer
	// fills up!
	//
	// So we’re filling the buffer up 3 “slots” that will pop-off immediately
	// due to the time.Now()
	//
	// After that, the for loop exists and the buffer continues to fill up with
	// “slots” delayed by the calls to time.Tick().
	for i := 0; i < request_threshold; i++ {
		rate_limiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * time.Duration(delay)) {
			rate_limiter <- t
		}
	}()

	return rate_limiter
}

func main() {
	// advanced_rate_limiter := burstRateLimiter(3, 200)
	// No! ^^^^^^^
	//
	// This breaks up here because the burstRateLimiter() function triggers
	// a goroutine which kicks off calls to time.Tick() that eventually fill
	// up the entire buffer. Then, because the basic rate limiter test runs,
	// it blocks on the logic below it from immediately executing so that,
	// by the time the requests against the advanced rate limiter are sent in,
	// the rate limiter’s buffer has ticked at least 3 times (600ms).
	//
	// 1. Basic Rate Limiter Declared
	//
	// 2. Advanced Rate Limiter Declared
	//     Initial State: [ 0ms ] [ 0ms ] [ 0ms ] [200ms] [200ms]
	//
	// 3. Basic Requests Made (5, rate-limited to 1 per 200ms)
	//     +----+-----------------------------------------+
	//     | #  | Advanced Rate Limiter State             |
	//     +----+-----------------------------------------+
	//     | 1  | [ 0ms ] [ 0ms ] [ 0ms ] [200ms] [200ms] |
	//     | 2  | [ 0ms ] [ 0ms ] [200ms] [200ms] [200ms] |
	//     | 3  | [ 0ms ] [200ms] [200ms] [200ms] [200ms] |
	//     | 4  | [200ms] [200ms] [200ms] [200ms] [200ms] |
	//     | 5  | [200ms] [200ms] [200ms] [200ms] [200ms] |
	//     +----+-----------------------------------------+
	//
	// 4. Advanced Requests Made
	//     OH NO! Burst rate limiter failed because the buffer of its channel
	//     has already filled up with 200ms ticks!

	basic_rate_limiter := basicRateLimiter(200)
	fmt.Println("Making requests with basic rate limiting…")
	mockRequests(5, basic_rate_limiter)
	fmt.Println("Done.")

	fmt.Println("")

	advanced_rate_limiter := burstRateLimiter(3, 200)
	fmt.Println("Making requests with advanced rate limiting…")
	mockRequests(5, advanced_rate_limiter)
	fmt.Println("Done.")
}
