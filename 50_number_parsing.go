package main

import (
	"fmt"
	"strconv"
)

const BITS = 64

func main() {
	int1, _ := strconv.ParseInt("42", 10, BITS)
	fmt.Println(int1)

	float1, _ := strconv.ParseFloat("42.6", BITS)
	fmt.Println(float1)

	// If 0 is passed to ParseInt, Go will attempt to
	// infer the base from the string.
	int2, _ := strconv.ParseInt("0xff", 0, BITS)
	fmt.Println(int2)

	// C’s atoi method is provided as a convenience
	// method for base-10 conversions
	int3, _ := strconv.Atoi("9001")
	fmt.Println(int3)

	// The strconv package can also parse things
	// like boolean values!
	b1, _ := strconv.ParseBool("1")
	b2, _ := strconv.ParseBool("true")
	fmt.Println(b1)
	fmt.Println(b2)
}
