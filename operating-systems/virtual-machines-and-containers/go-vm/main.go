package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run <container> <command> <args>
// go run main.go run cmd args
func main() {
	command := os.Args[1]
	switch command {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("wat should I do with " + command)
	}
}

// Represents main process
func run() {
	command := os.Args[2]
	args := os.Args[3:]
	fmt.Printf("Running %s with %v\n", command, args)

	cmd := exec.Command(command, args...)

	// Redirects the output of the command to the standard output
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Namespace for the child process to not be able to access parent
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	// Runs the command structure
	must(cmd.Run())
}

// Represents child process
func child() {
	fmt.Println("Child")
}

// Helper to check that it alwayws runs
func must(err error) {
	if err != nil {
		panic(err)
	}
}
