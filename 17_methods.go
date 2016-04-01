package main

import "fmt"

type rect struct {
	width  int
	height int
}

// You can pass the receiver, (r rect), as a value…
func (r rect) area() int {
	return r.width * r.height
}

// …or as a pointer/reference
//
// When? Why?
// 1. If you need to actually alter the receiver (e.g., read/write vs. read)
// 2. If a deep-copy would be expensive
// 3. To maintain consistency of existing API
func (r *rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	test_rect := rect{5, 10}

	fmt.Println(test_rect.area())
	fmt.Println(test_rect.perimeter())
}
