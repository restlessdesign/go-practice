package main

import "fmt"

const s string = "you_shall_not_reassign!"

func main() {
	fmt.Println(s)

	const a = 426
	const b = 5 * 5

	fmt.Println(a)
	fmt.Println(b)
}
