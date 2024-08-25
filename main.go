package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
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

	scanner := bufio.NewScanner(os.Stdin)
	// Read all input from stdin
	for scanner.Scan() {
		input := scanner.Text()
		execCmd(cmd, args, input)
	}

	if err := scanner.Err(); err != nil {
		// fmt.Println("Error reading from stdin:", err)
		os.Exit(1)
	}
}

func execCmd(cmd string, args []string, input string) {
	// Compile the regular expression
	re := regexp.MustCompile(`:\d+`)

	var hasArgs bool
	for _, arg := range args {
		if re.MatchString(arg) {
			hasArgs = true
			break
		}
	}

	if hasArgs {
		execCmdWithArgs(cmd, args, input)
	} else {
		execSimpleCmd(cmd, args, input)
	}
}

func execCmdWithArgs(cmd string, args []string, input string) {
	str := strings.Join(args, " ")
	values := strings.Split(input, " ")

	x := replacePlaceholders(str, values)
	newArgs := strings.Split(x, " ")

	runCommand(cmd, newArgs)
}

func replacePlaceholders(str string, values []string) string {
	// Compile the regular expression
	re := regexp.MustCompile(`:(\d)`)

	// Find all matches and their submatches
	matches := re.FindAllStringSubmatchIndex(str, -1)

	// Create a result string builder
	var result string
	lastIndex := 0

	for _, match := range matches {
		if len(match) == 4 {
			// Append the part of the input string before the match
			result += str[lastIndex:match[0]]

			// Convert the matched group to an integer
			index, err := strconv.Atoi(str[match[2]:match[3]])
			if err != nil {
				continue
			}

			// Replace the placeholder with the corresponding value
			if index > 0 && index <= len(values) {
				result += values[index-1]
			}

			// Update the last index to the end of the current match
			lastIndex = match[1]
		}
	}

	// Append the remaining part of the input string
	result += str[lastIndex:]

	return result
}

func execSimpleCmd(cmd string, args []string, input string) {
	args = append(args, input)

	runCommand(cmd, args)
}

func runCommand(cmd string, args []string) {
	command := exec.Command(cmd, args...)
	output, err := command.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}

	fmt.Print(string(output))
}
