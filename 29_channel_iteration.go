package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "foo"
	queue <- "baz"
	close(queue)

	// The queue was closed, but we can still grab values out of it!
	// If we didn’t close the channel, we would get deadlock because the loop
	// would be waiting on a third value
	for message := range queue {
		fmt.Println(message)
	}
}
