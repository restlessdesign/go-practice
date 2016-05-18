package main

import "fmt"

func main() {
	// Refresher: Channels must be created with make(chan <type>)
	// The second parameter (in this case, 5) specifies the size of the
	// channel’s buffer.
	jobs := make(chan int, 5)

	// This will be used to await (or “sync” with) the result of the goroutines
	// sent out as jobs
	done := make(chan bool)

	go func() {
		// There’s no for clause here!? This is actually shorthand for:
		//
		//   for true {
		//       // etc.
		//   }
		//
		// In other words, this is an infinite loop! We return when there are
		// no more jobs to process.
		//
		// I *think* we could also write this like so:
		//
		//   for j, more := <-jobs; more {
		//     // etc.
		//   }
		//
		// …but then we wouldn’t be able to test against `if more` … `else` …
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("No more jobs.")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sending job", j)
	}

	// Prevent more jobs from being sent
	close(jobs)
	fmt.Println("Sent all jobs.")

	<-done
}
