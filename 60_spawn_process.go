package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dateCmd := exec.Command("date")

	// `.Output()` runs the command, waits for it to finish, and returns the
	// command’s STDOUT result
	dateOut, err := dateCmd.Output()
	check(err)

	fmt.Println("date >", string(dateOut))

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	// Start will begin the command but NOT wait for it to complete, `.Wait()`
	// will take care of that.
	grepCmd.Start()

	// Write some data that will be piped into our `grep` command, then close the pipe
	grepIn.Write([]byte("hello world\ngoodbye world"))
	grepIn.Close()

	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("grep >", string(grepBytes))

	// NOTE: Arguments must be passed as an array!
	// We will unpack this using `...` syntax below…
	lsArgs := []string{"-l", "-a"}
	lsCmd := exec.Command("ls", lsArgs...)
	lsOut, err := lsCmd.Output()
	check(err)

	fmt.Println("ls >", string(lsOut))
}
