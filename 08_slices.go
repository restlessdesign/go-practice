package main

import (
	"fmt"
	"reflect"
)

func main() {
	number_slice := make([]int, 2)
	fmt.Println(number_slice)

	string_slice := make([]string, 4)
	fmt.Println(string_slice)

	string_slice[0] = "foo"
	string_slice[1] = "bar"
	string_slice[2] = "baz"
	string_slice[3] = "qux"
	for i := 0; i < len(string_slice); i++ {
		fmt.Println(i, string_slice[i])
	}

	string_slice = append(string_slice, "apples")
	string_slice = append(string_slice, "bananas")
	fmt.Println(string_slice)

	copy_slice := make([]string, len(string_slice))
	copy(copy_slice, string_slice)
	copy_slice = append(copy_slice, "carrots")
	fmt.Println(copy_slice)

	sub_slice := copy_slice[0:3]
	fmt.Println(sub_slice)

	sub_slice = copy_slice[4:]
	fmt.Println(sub_slice)

	// Allegedly a slice...
	t := []string{"g", "h", "i"}
	fmt.Println(reflect.TypeOf(t))

	// Allegedly an array...
	// ...but it's the (almost) same initialization syntax?
	b := [3]string{"g", "h", "i"}
	fmt.Println(reflect.TypeOf(b))

	// So the TypeOf is the same...
	// The difference must be the length that is passed?
	// Calling append() should verify
	t = append(t, "j")
	fmt.Println(reflect.TypeOf(t))

	// Yep! This throws an error warning that the first argument to
	// append() must be a slice.
	// b = append(b, "j")
	// fmt.Println(reflect.TypeOf(b))

	// Definitely not the way to prepopulate a multi-dimensional array =/
	// multi_slice := []{[2]string{"foo", "bar"}, [2]string{"dog", "cat"}}
	// fmt.Println(multi_slice)

	// multi_slice := make([][]int, 3)
	// (to be continued...)
}
