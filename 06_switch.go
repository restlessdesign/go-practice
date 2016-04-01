package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Weekday(), now.Month(), now.Day(), now.Year())

	switch now.Weekday() {
	case 5:
	case 6:
		fmt.Println("Party on, Wayne!")
	default:
		fmt.Println("Back to the grind")
	}
}
