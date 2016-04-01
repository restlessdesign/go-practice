package main

import "fmt"

func add(a int, b int) int {
	return a + b
}


// Generally speaking, /interface{} is an anti-pattern
func expect(tested interface{}, expected interface{}) bool {
	equal := tested == expected
	if equal {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL (Expected ", expected, " to equal ", tested)
	}

	return tested == expected
}

func main() {
	fmt.Println(add(1, 1))
	fmt.Println(add(2, 2))

	expect(add(1, 1), 2)
}
