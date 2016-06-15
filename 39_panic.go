package main

import "os"

func main() {
	_, err := os.Open("/tmp/spooky")
	if err != nil {
		panic("File doesn’t exist…panic and freak out!!!")
	}
}
