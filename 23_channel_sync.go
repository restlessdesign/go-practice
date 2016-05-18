package main

import (
	"fmt"
	"time"
)

func delay(done chan bool) {
	fmt.Println("Delaying execution…")
	time.Sleep(time.Second)
	fmt.Println("Done.")

	done <- true
}

func chain(cur_val int, resolve chan int) {
	next_val := cur_val + 1
	fmt.Printf("Sleeping %d seconds\n", next_val)
	time.Sleep(time.Second * time.Duration(next_val))

	resolve <- next_val
}

func main() {
	fmt.Println("Program start")
	done := make(chan bool, 1)
	go delay(done)

	// Very important to call on the result (<-done), even if we don’t need to
	// assign it to a variable. Without it, the program would exit before the
	// goroutine had a chance to run!
	notified := <-done
	fmt.Println(notified)

	// Let’s make things a little more interesting and recursively pass the
	// result back into the handler. This could also be passed to another
	// method ala Promise chains in JavaScript
	chained_channel := make(chan int, 1)
	go chain(0, chained_channel)
	for _, val := range []int{1, 2, 3, 4} {
		fmt.Println(val)
		go chain(<-chained_channel, chained_channel)
	}

	result := <-chained_channel
	fmt.Println(result)

	fmt.Println("Program exit")
}
