package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var state = make(map[int]int)
var mutex = &sync.Mutex{}
var ops_count int64 = 0

// Grab a random value and atomically-increment our opts_count
func readAsync() {
	total := 0

	for {
		// Pick a random integer between 0 and 5
		key := rand.Intn(5)

		mutex.Lock()
		total += state[key]
		mutex.Unlock()

		// Keep track of total operations
		atomic.AddInt64(&ops_count, 1)

		// Allow other goroutines to run. This is safe because we are using a
		// mutex/semaphore/lock to restric access to our state object so that
		// only one process may read from it at a time.
		runtime.Gosched()
	}
}

func writeAsync() {
	for {
		key := rand.Intn(5)
		val := rand.Intn(100)

		mutex.Lock()
		state[key] = val
		mutex.Unlock()

		atomic.LoadInt64(&ops_count)
		runtime.Gosched()
	}
}

func main() {
	for i := 0; i < 100; i++ {
		go readAsync()
	}

	for i := 0; i < 10; i++ {
		go writeAsync()
	}

	// Give the operations a chance to complete
	time.Sleep(time.Second)

	total_ops := atomic.LoadInt64(&ops_count)
	fmt.Println("Total Operations:", total_ops)

	mutex.Lock()
	fmt.Println("State:", state)
	mutex.Unlock()
}
