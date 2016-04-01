package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[0] = 4
	a[1] = 2
	a[2] = 6
	fmt.Println(a)
	fmt.Println(a[0], a[1], a[2])

	prefilled_int_arr := [5]int{3, 1, 1, 4, 5}
	fmt.Println(prefilled_int_arr)

	prefilled_str_arr := [2]string{"foo", "bar"}
	fmt.Println(prefilled_str_arr)

	var multi_dim_arr [2][4]int
	multi_dim_arr[0][0] = 1
	fmt.Println(multi_dim_arr)

	// let compiler count for us
	psuedo_dynamic_arr := [...]string{"cat", "dog", "mouse"}
	fmt.Println(len(psuedo_dynamic_arr))
}
