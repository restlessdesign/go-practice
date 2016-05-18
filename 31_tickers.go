package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 1000)

	go func() {
		for tick := range ticker.C {
			fmt.Println("Tick", tick)
		}
	}()

	time.Sleep(time.Millisecond * 5001)
	ticker.Stop()
	fmt.Println("The mouse ran up the clock.")
}
