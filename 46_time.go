package main

import (
	"fmt"
	"time"
)

const hours_per_year = 365 * 24

func main() {
	now := time.Now()
	fmt.Println(now)

	dob := time.Date(1985, 04, 26, 10, 45, 0, 0, time.UTC)
	fmt.Println(dob)

	fmt.Printf("I was born on %s, %s %d, %d\n",
		dob.Weekday(),
		dob.Month(),
		dob.Day(),
		dob.Year())

	age := time.Since(dob)
	hours := age.Hours()
	years := hours / hours_per_year

	fmt.Printf("That makes me %d years old!\n", int(years))
}
