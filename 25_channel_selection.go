package main

import (
	"fmt"
	"time"
)

func main() {
	chan_a := make(chan string, 1)
	chan_b := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		chan_a <- "sleep one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		chan_b <- "sleep two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case message_a := <-chan_a:
			fmt.Println("Received on channel A:", message_a)
		case message_b := <-chan_b:
			fmt.Println("Received on channel B", message_b)
		}
	}
}
