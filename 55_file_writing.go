package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\nworld")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	// write byte slice to file
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("Wrote %d bytes\n", n2)

	// write string to file
	n3, err := f.WriteString("writes\n")
	fmt.Printf("Wrote %d bytes\n", n3)

	// flush writes to stable storage
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("Wrote %d bytes\n", n4)

	// flush writes to stable storage
	w.Flush()
}
