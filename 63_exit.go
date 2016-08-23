package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Deferred calls don’t run when `os.Exit` is triggered!")
	os.Exit(3)
}
