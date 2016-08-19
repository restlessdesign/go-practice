package main

import (
	"flag"
	"fmt"
)

func main() {
	// "word" = flag
	// "foo" = default value
	// "a string" = description
	//
	// Note: `flag.String()`, `flag.Int()`, etc. return a POINTER—not a value!
	// This will be important later…
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("number", 1, "an integer")
	boolPtr := flag.Bool("fork", false, "a boolean")

	var svar string
	// The `flag.StringVar()` method allows us to assign a flag back to a
	// pre-existing variable in our program. In this case, we assign it to the
	// pointer to `svar`, which we declare above.
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Execute command line parsing now that flags have been declared
	flag.Parse()

	// Remember when we said that methods like `flag.String()` return a pointer?
	// Now is the time to access the actual values using asterisk notation,
	// which returns a pointer’s value!
	fmt.Println("wordPtr", *wordPtr)
	fmt.Println("numPtr", *numPtr)
	fmt.Println("boolPtr", *boolPtr)

	fmt.Println("svar", svar)

	fmt.Println("tail:", flag.Args())
}

// ./58_command_line_flags --word="foobar" --number=5 --fork
// ./58_command_line_flags --word="foobar" --fork trailing args go here
// ./58_command_line_flags -h
// ./58_command_line_flags --help
