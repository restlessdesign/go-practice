package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// time.Unix() and time.UnixNano() return
	// elapsed time since epoch
	secs := now.Unix()
	nanos := now.UnixNano()

	// Milliseconds don’t exist within Unix
	// methods, so we must convert them manually
	millis := nanos / 1000000

	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
