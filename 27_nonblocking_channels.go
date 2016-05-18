package main

import "fmt"

func main() {
	chan_1 := make(chan string)

	go func() {
		chan_1 <- "First!"
	}()

	select {
	case result := <-chan_1:
		fmt.Println(result, "I come first because I block")
	}

	fmt.Println("I come second because I was blocked above")

	chan_2 := make(chan string)

	go func() {
		chan_2 <- "Third!"
	}()

	select {
	case result := <-chan_2:
		fmt.Println(result, "I don’t fire")
	default:
		// fmt.Println("Skip")
	}

	fmt.Println("I come third because the above select is non-blocking due to its default case")
}
