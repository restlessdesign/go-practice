package main

import "fmt"

func incrementBy(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	add10To := incrementBy(10)
	fmt.Println(add10To(5))

	add100To := incrementBy(100)
	fmt.Println(add100To(25))
}
