package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "d", "b"}
	fmt.Println(strs, sort.StringsAreSorted(strs))
	sort.Strings(strs)
	fmt.Println(strs, sort.StringsAreSorted(strs))

	ints := []int{5, 3, 1, 6, 4, 2}
	fmt.Println(ints, sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println(ints, sort.IntsAreSorted(ints))
}
