package main

import (
	"fmt"
	"sort"
)

// Create a new sort that ranks by string length (shortest to longest)
type ByLength []string

// sort.Interface() requires Len(), Swap(), and Less()
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s ByLength) Less(a, b int) bool {
	return len(s[a]) < len(s[b])
}

// Create a new sort that separates even numbers from odd numbers, from lowest to highest
type EvenPriority []int

func (num EvenPriority) Len() int {
	return len(num)
}

func (num EvenPriority) Swap(a, b int) {
	num[a], num[b] = num[b], num[a]
}

func (num EvenPriority) Less(a, b int) bool {
	a_even := num[a]%2 == 0
	b_even := num[b]%2 == 0
	both_even := a_even && b_even

	if both_even {
		return num[a] < num[b]
	}

	return a_even
}

func main() {
	foo := []string{
		"xxxx",
		"x",
		"xxxxxx",
		"xxx",
		"xxxxxxxx",
		"xx",
		"xxxxxxx",
		"xxxxx",
	}

	fmt.Println(sort.IsSorted(ByLength(foo)))

	sort.Sort(ByLength(foo))
	for _, val := range foo {
		fmt.Println(val)
	}

	fmt.Println(sort.IsSorted(ByLength(foo)))

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums, sort.IsSorted(EvenPriority(nums)))
	sort.Sort(EvenPriority(nums))
	fmt.Println(nums, sort.IsSorted(EvenPriority(nums)))
}
