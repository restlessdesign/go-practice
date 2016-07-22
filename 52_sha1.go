package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	s := "some string value"
	h := sha1.New()

	// Have to convert the string to bytes
	b := []byte(s)
	h.Write(b)

	// Summing gets our actual SHA-1 value
	// The argument passed in can be used to append to
	// an exiting byte slice, and usually isn’t needed.
	sum := h.Sum(nil)

	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(sum)

	// Sums are still bytes, so convert to hex for display
	fmt.Printf("%x\n", sum)

	// We can also use the hex class to do this
	fmt.Println(hex.EncodeToString(sum))
}
