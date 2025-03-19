package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		// read the keyboad input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// handle the execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// ErrNoPath is returned when 'cd' was called without a second arguments
var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	// remove the newline char
	input = strings.TrimSuffix(input, "\n")

	// split the input to serparate the command and the arguments
	args := strings.Split(input, " ")

	// check for built in commands
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)

	// set the correct output devicce
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// exacut the command
	return cmd.Run()
}
