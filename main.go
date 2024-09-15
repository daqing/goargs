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
//
// For file name with space, we set :0 to the file name
// so:
// echo Frame 123.svg | goargs mv :0 :2
// will rename "Frame 123.svg" to "123.svg"
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

// Input examples:
// foo.go
// Frame 123.svg
//
// args examples:
// []string{":0", ":1.bak"}
// []string{":1", ":2"}
func execCmd(cmd string, args []string, input string) {
	// Compile the regular expression
	re := regexp.MustCompile(`:\d+`)

	var hasPlaceholder bool
	for _, arg := range args {
		if re.MatchString(arg) {
			hasPlaceholder = true
			break
		}
	}

	// if input has placeholder, we need to replace it with the input
	if hasPlaceholder {
		execCmdWithPlaceholders(cmd, args, input)
	} else {
		execSimpleCmd(cmd, args, input)
	}
}

func execCmdWithPlaceholders(cmd string, args []string, input string) {
	str := strings.Join(args, " ")

	runCommand(cmd, replacePlaceholders(str, input))
}

// given str as "mv :0 :2.bak", and input as "Frame 123.svg"
// then the result returned is:
// []string{"mv", "Frame\\ 123.svg", "123.svg.bak"}
func replacePlaceholders(str string, input string) []string {
	// Compile the regular expression
	re := regexp.MustCompile(`:(\d)`)

	// Split the input string into arguments
	args := strings.Split(str, " ")

	// Split the input string into values
	values := strings.Split(input, " ")

	// Create a result slice to store the replaced arguments
	result := make([]string, 0, len(args))

	for _, arg := range args {
		// Find all matches in the current argument
		matches := re.FindAllStringSubmatchIndex(arg, -1)

		if len(matches) == 0 {
			// If no placeholders found, add the argument as is
			result = append(result, arg)
			continue
		}

		// Create a new string builder for the current argument
		var newArg strings.Builder
		lastIndex := 0

		for _, match := range matches {
			if len(match) == 4 {
				// Append the part of the argument before the match
				newArg.WriteString(arg[lastIndex:match[0]])

				// Convert the matched group to an integer
				index, err := strconv.Atoi(arg[match[2]:match[3]])
				if err != nil {
					continue
				}

				// Replace the placeholder with the corresponding value
				if index == 0 {
					// for :0, use the whole input
					newArg.WriteString(input)
				} else if index > 0 && index <= len(values) {
					newArg.WriteString(values[index-1])
				}

				// Update the last index to the end of the current match
				lastIndex = match[1]
			}
		}

		// Append the remaining part of the argument
		newArg.WriteString(arg[lastIndex:])

		// Add the processed argument to the result
		result = append(result, newArg.String())
	}

	return result
}

func execSimpleCmd(cmd string, args []string, input string) {
	args = append(args, input)

	runCommand(cmd, args)
}

func runCommand(cmd string, args []string) {
	// fmt.Printf("cmd=%s, args=%v \n", cmd, args)

	command := exec.Command(cmd, args...)
	output, err := command.Output()
	if err != nil {
		fmt.Printf("Error: %s, output: %v\n", err, output)
		os.Exit(1)
	}

	fmt.Print(string(output))
}
