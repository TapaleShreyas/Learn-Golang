package main

import (
	"fmt"
	"os"
)

func main() {

	// Command to run this program is
	// $ go run 05-switch-cases.go hello

	/*
		- What is fallthrough in Go?
		- there is no break statement for case in go
		- it is fine even if you don't add code for case
		- you can have multiple conditions into a single case seperated by comma's.
		- switch statement can have initialization variable to, life of variable is switch.
		- you can have a switch statement without passing a switch variable for comparison
		- you can have one case with numeric variable comparison and other with string comparison also.
		- you can define switch without switch variable but make sure to put semicolon
	*/

	/*
		input := os.Args[1] // reading value from command line screen
		fmt.Println("command line input", input)
		if input == "hello" {
			fmt.Println("Hi yourself")
		} else if input == "goodbye" {
			fmt.Println("So long!")
		} else if input == "greetings" {
			fmt.Println("Salutations")
		} else {
			fmt.Println("I don't know what you said")
		}
	*/

	// you can achieve the same using switch code as below
	// simple switch case
	input := os.Args[1]
	switch input {
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye":
		fmt.Println("So long!")
	case "greetings":
		fmt.Println("Salutations")
	default:
		fmt.Println("I don't know what you said")
	}

	fmt.Println("\n\n")
	// Command to run this program is
	// $ go run 05-switch-cases.go hi
	fmt.Println("use of fallthrough")
	// fallthrough makes flow to execute next case statement
	switch input {
	case "hi":
		fmt.Println("Very informal!")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye", "go":
		fmt.Println("So long!")
	case "greetings":
		fmt.Println("Salutations")
	default:
		fmt.Println("I don't know what you said")
	}

	fmt.Println("\n\n")
	fmt.Println("comma separated condition for switch case")
	fmt.Println("and switch case without case body")

	// Command to run this program is
	// $ go run 05-switch-cases.go bye
	switch input {
	case "hi":
		fmt.Println("Very informal!")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	case "farewell":
	case "goodbye", "bye":
		fmt.Println("So long!")
	case "greetings":
		fmt.Println("Salutations")
	default:
		fmt.Println("I don't know what you said")
	}

	fmt.Println("\n\n")
	fmt.Println("you can pass own string(variable) for case , different from switch variable")
	// Command to run this program is
	// $ go run 05-switch-cases.go greetings
	greet := "greetings"
	switch input {
	case "hi":
		fmt.Println("Very informal!")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye", "bye":
		fmt.Println("So long!")
	case greet:
		fmt.Println("Salutations")
	default:
		fmt.Println("I don't know what you said")
	}

	fmt.Println("\n\n")
	fmt.Println("local variable for switch case")
	// Command to run this program is
	// $ go run 05-switch-cases.go nomatch
	greet = "greetings"
	switch l := len(input); input {
	case "hi":
		fmt.Println("Very informal!")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	case "farewell":
	case "goodbye", "bye":
		fmt.Println("So long!")
	case greet:
		fmt.Println("Salutations")
	default:
		fmt.Println("I don't know what you said but it was", l, "characters long.")
	}

	fmt.Println("\n\n")
	fmt.Println("switch variable along with local variable condition case")
	fmt.Println("we have not passed variable to switch")
	fmt.Println("note semicolon after local variable initialization")
	// Command to run this program is
	// $ go run 05-switch-cases.go crackerjack
	word := os.Args[1]
	c := "crackerjack"
	switch l := len(word); {
	case word == "hi":
		fmt.Println("Very informal!")
		fallthrough
	case word == "hello":
		fmt.Println("Hi yourself")
	case l == 1:
		fmt.Println("I don't know any one letter words")
	case 1 < l && l < 10, word == c:
		fmt.Println("This word is either", c, "or 2-9 characters long")
	default:
		fmt.Println("I don't know what you said but it was", l, "characters long.")
	}
}
