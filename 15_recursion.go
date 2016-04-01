package main

import "fmt"

func countdown(start int) int {
	fmt.Println(start)

	if start > 0 {
		return countdown(start - 1)
	}

	return 0
}

func main() {
	countdown(10)
}
