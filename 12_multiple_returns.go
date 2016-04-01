package main

import "fmt"

func echo(a string, b string) (string, string) {
	return a, b
}

func main() {
	foo, bar := echo("foo", "bar")
	fmt.Println(foo, bar)
}
