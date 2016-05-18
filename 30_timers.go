package main

import (
	"fmt"
	"time"
)

func main() {
	// This won’t be called since the below channel call will block. That’s
	// because we aren’t deferring the code out to a goroutine!
	// fmt.Println("Starting timer...")
	// timer_1 := time.NewTimer(time.Second * 1)

	// block on the timer’s channel
	// <-timer_1.C
	// fmt.Println("...done!")

	fmt.Println("before goroutine")

	timer_2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer_2.C
		fmt.Println("inside goroutine")
	}()

	stop_result := timer_2.Stop()
	if stop_result {
		fmt.Println("timer stopped")
	}

	fmt.Println("after goroutine")
}
