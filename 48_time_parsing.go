package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Behold! Our various and glorious time formats.
	fmt.Println(now.Format(time.RFC1123))
	fmt.Println(now.Format(time.RFC3339))
	fmt.Println(now.Format(time.RFC822))
	fmt.Println(now.Format(time.RFC850))
	fmt.Println(now.Format(time.RFC1123Z))
	fmt.Println(now.Format(time.RFC3339Nano))
	fmt.Println(now.Format(time.RFC822Z))
	fmt.Println("")

	parsed, err := time.Parse(
		time.RFC3339,
		"1985-04-26T10:45:00+05:00")

	if err == nil {
		fmt.Println(parsed)
	}
}
