package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cwd, err := os.Getwd()
	check(err)

	fp := cwd + "/" + "54_file_reading.go"

	// Read an entire file’s contents into memory
	dat, err := ioutil.ReadFile(fp)
	check(err)
	fmt.Println(string(dat))

	// Open a file object for more granular reading
	f, err := os.Open(fp)
	check(err)

	// Read the first line of this file (“package main”)
	// First, create a byte array with a predefined length
	b1 := make([]byte, len("package main"))

	// File.Read() will populate our byte array.
	// It also returns the number of bytes read (n1) and an err
	n1, err := f.Read(b1)
	check(err)

	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// Can also do some seeking within the file.
	// The second (“whence”) parameter indicates where to seek from:
	// - 0 = Beginning of file
	// - 1 = Relative to current offset/cursor location
	// - 2 = Relative to end of file
	f.Seek(int64(len("package main")), 0)
	b2 := make([]byte, 10)
	f.Read(b2)
	fmt.Println(string(b2))

	// Seek ahead 20 characters from current offset
	f.Seek(20, 1)

	// Use `io.ReadAtLeast()` to fill an expected number of bytes into buffer
	// Note! If the `min` argument passed to this method is larger than the size
	// of the buffer, an error will be thrown! An error will also be thrown if
	// EOF is reached before `min` bytes have been read into the buffer
	b3 := make([]byte, 50)
	_, err = io.ReadAtLeast(f, b3, 20)
	check(err)
	fmt.Println(string(b3))

	fmt.Println("")
	fmt.Println("=============================================================")
	r4 := bufio.NewReader(f)
	f.Seek(0, 0)
	b4, err := r4.Peek(128)
	check(err)
	fmt.Println(string(b4))
	fmt.Println("=============================================================")
	fmt.Println("")

	// Close the file (normally done using defer immediately after `os.Open()`)
	f.Close()
}
