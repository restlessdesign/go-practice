// Be careful with the sync/atomic package! As per the Go documentation:
//
// “Except for special, low-level applications, synchronization is better done
// with channels or the facilities of the sync package. Share memory by
// communicating; don't communicate by sharing memory.”
package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				// Incrementing by reference for atomic operation
				atomic.AddUint64(&ops, 1)

				// Allow other Goroutines to continue as normal
				runtime.Gosched()
			}
		}()
	}

	// Let the operations collection for a bit…
	time.Sleep(time.Second)

	// We want to load this atomically to ensure we get the final value after
	// goroutines have had time to hammer on incrementing this value.
	ops_final := atomic.LoadUint64(&ops)

	fmt.Println("Ops:", ops_final)
}
