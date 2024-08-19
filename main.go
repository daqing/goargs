package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
)

// How to use goargs
// find . -name '*.go' | goargs wc -l
// find . -name '*.go' | goargs mv :1 :1.bak
// find . -name '*.go' | awk -F/ '{print $1, $2}' | goargs cat :1/:2
func main() {
	// read command from args
	if len(os.Args) == 1 {
		fmt.Println("Please provide a command")
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	reader := bufio.NewReader(os.Stdin)
	// Read all input from stdin
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		}

		if err == io.EOF {
			fmt.Println("All lines are processed.")
			os.Exit(0)
		}

		// Print the input read from stdin
		fmt.Println("Line:", input)

		execCmd(cmd, args, input[:len(input)-1])
	}
}

func execCmd(cmd string, args []string, input string) {
	fmt.Printf("cmd={%s}, input={%s}\n", cmd, input)

	// Compile the regular expression
	re := regexp.MustCompile(`:\d+`)

	if re.MatchString(cmd) {
		execCmdWithArgs(cmd, args, input)
	} else {
		execSimpleCmd(cmd, args, input)
	}
}

func execCmdWithArgs(cmd string, args []string, input string) {

}

func execSimpleCmd(cmd string, args []string, input string) {
	// simpleCmd := cmd + " " + input
	// fmt.Println("Simple command:", simpleCmd)

	args = append(args, input)

	command := exec.Command(cmd, args...)
	output, err := command.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
