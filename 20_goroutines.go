package main

import "fmt"

func count(context string) {
	for i := 0; i < 10; i++ {
		fmt.Println(context, i)
	}
}

func main() {
	count("sync")

	go count("async a")
	go count("async b")

	// Our two function calls are running asynchronously in separate goroutines
	// now, so execution falls through to here. This `Scanln` code requires we
	// press a key before the program exits.
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
