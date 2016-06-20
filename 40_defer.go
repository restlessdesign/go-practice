package main

import (
	"fmt"
	"os"
)

func createFile(filepath string) *os.File {
	fmt.Println("Creating", filepath)

	file, err := os.Create(filepath)
	if err != nil {
		panic("Unable to create file")
	}

	return file
}

func closeFile(file *os.File) {
	fmt.Println("Closing", file.Name())
	file.Close()
}

func writeFile(data string, file *os.File) {
	fmt.Println("Writing", data, "to", file.Name())
	fmt.Fprintln(file, data)
}

func main() {
	file := createFile("/tmp/godefer.txt")
	// This will execute just before the end of this method, even if there’s an
	// error! Think of it like the “finally” clause in try, catch, finally…
	defer closeFile(file)
	writeFile("♫ I bless the rains down in Africaaaaaa… ♫", file)
}
