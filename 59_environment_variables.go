package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FIRST_NAME", "Kevin")
	fmt.Println("FIRST_NAME:", os.Getenv("FIRST_NAME"))

	for _, arg := range os.Environ() {
		keyval := strings.Split(arg, "=")
		fmt.Println(keyval[0], keyval[1])
	}
}
