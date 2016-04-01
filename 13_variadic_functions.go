package main

import "fmt"

func expect(tested interface{}, expected interface{}) bool {
	equal := tested == expected
	if equal {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL (Expected ", expected, " to equal ", tested, ")")
	}

	return tested == expected
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 1))
	fmt.Println(sum(2, 2, 6))

	expect(sum(1, 1), 2)
	expect(sum(2, 2, 6), 10)
	expect(sum(0, 0, 0), 0)
	expect(sum(1, 0, 0), 1)
	expect(sum(1, 1, 0), 2)
	expect(sum(1, 0, 1), 2)
	expect(sum(1, 1, 1), 3)

	nums := []int{10, 5, 5}
	expect(sum(nums...), 20)
}
