//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CommandBye   = "bye"
	CommandHello = "hello"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nonBlankLine, commands := 0, 0
	for scanner.Scan() {
		theCommand := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if theCommand == "q" {
			break
		} else {
			switch theCommand {
			case CommandBye:
				commands += 1
				fmt.Println("Bye Bye!")
			case CommandHello:
				commands += 1
				fmt.Println("Hello There!")
			}
			if theCommand != "" {
				nonBlankLine += 1
			}
		}
	}

	// Reader method
	// r := bufio.NewReader(os.Stdin)
	// for {
	// 	input, inputErr := r.ReadString('\n')
	// 	n := strings.TrimSpace(input)
	// 	//* When the user enters either "hello" or "bye", the program
	// 	//  should respond with a message after pressing the enter key.
	// 	if n == "hello" {
	// 		commands += 1
	// 		nonBlankLine += 1
	// 		fmt.Println("Hello there")
	// 	}
	// 	if n == "bye" {
	// 		commands += 1
	// 		nonBlankLine += 1
	// 		fmt.Println("Bye Bye!")
	// 	}
	// 	//* Whenever the user types a "Q" or "q", the program should exit.
	// 	//* Upon program exit, some usage statistics should be printed
	// 	if n == "Q" || n == "q" {
	// 		fmt.Println("Goodbye!")
	// 		break
	// 	}
	// 	if n == "" {
	// 		continue
	// 	}
	// 	if inputErr != nil {
	// 		fmt.Println("Inout error:", inputErr)
	// 	}
	// }
	//  - The number of non-blank lines entered
	//  - The number of commands entered
	fmt.Printf("---Stats \n Non-blank lines: %v \n commands entered : %v", nonBlankLine, commands)

}
