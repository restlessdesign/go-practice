package main

import (
	"os"
	"os/exec"
	"syscall"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Find our command’s location within our execution PATH
	bin, lookErr := exec.LookPath("ls")
	check(lookErr)

	// First argument needs to be program name
	lsArgs := []string{"ls", "-l", "-a", "-h"}

	// Use existing shell environment
	env := os.Environ()

	execErr := syscall.Exec(bin, lsArgs, env)
	check(execErr)
}
