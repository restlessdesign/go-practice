package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `Notify` accepts a channel and a variadic list of Unix signals that
	// should pass through to the channel. In this example, listen for the
	// common SIGINT (ctrl+c) and SIGTERM (`$ kill <pid>`).
	//
	// Remember: SIGKILL (`$ kill -9 <pid>`) cannot be ignored by a process,
	// so there’s no point adding that to our notifier!
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// This will block until the channel’s buffer is filled by SIGINT or SIGTERM
		sig := <-sigs

		fmt.Println()
		fmt.Println(sig)

		done <- true
	}()

	fmt.Println("Awaiting signal…")

	// This will also block until it receives the `true` value within our
	// goroutine above, which itself is blocked until a signal is received.
	<-done

	fmt.Println("Exiting.")
}
