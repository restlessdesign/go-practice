package main

import (
	"fmt"
)

// Only accepts input
func ping(input chan<- string, message string) {
	input <- message
}

// Passes through message from ping channel
func pong(ping <-chan string, output chan<- string) {
	message := <-ping
	output <- message
}

func main() {
	in := make(chan string, 1)
	out := make(chan string, 1)

	ping(in, "a man, a plan, a canal — panama")
	pong(in, out)
	fmt.Println(<-out)
}
