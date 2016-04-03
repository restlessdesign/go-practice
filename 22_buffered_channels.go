package main

import "fmt"

func main() {

	// the second parameter specifies that the channel may buffer
	// up to two values
	buffered_chan := make(chan string, 2)

	buffered_chan <- "hello"
	buffered_chan <- "is it me you’re looking for?"

	fmt.Println(<-buffered_chan)
	fmt.Println(<-buffered_chan)

	buffered_chan <- "i can see it in your eyes"
	buffered_chan <- "i can see it in your smile"

	// this will throw because the buffer is full!
	// buffered_chan <- "you’re all I've ever wanted"

	// flush one entry from the buffer
	fmt.Println(<-buffered_chan)

	// now we’re free to add one more entry to the buffer, taking us back up to
	// two buffered entries. let’s flush both of them!
	buffered_chan <- "you’re all I’ve ever wanted"
	fmt.Println(<-buffered_chan)
	fmt.Println(<-buffered_chan)
}
