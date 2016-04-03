package main

import "fmt"

func main() {
	// Channels must be created by invoking `make(chan <type>)`
	// Good to know: `make()` returns a copy; `new()` returns a pointer
	receiver := make(chan string)

	go func() {
		// Send value into the channel
		receiver <- "hello from the other side"
	}()

	// Retrieve output from the channel. Channels are unbuffered by default
	// and require an output:	`output := <-channel_name`
	// for every input:			`channel <- "input"`
	msg := <-receiver
	fmt.Println(msg)

	receiver_2 := make(chan string)
	go func() {
		receiver_2 <- "i must have called a thousand times"
	}()

	msg2 := <-receiver_2
	fmt.Println(msg2)
}
