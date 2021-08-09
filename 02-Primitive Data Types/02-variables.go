package main

import (
	"fmt"
	"os"
)

func main() {

	// Command to run this program is
	// $ go run 02-variables.go hello

	// simple variable decalaration
	// syntax:
	//		var var_name type
	// later you can assign the value to variable
	var name string
	name = "Shreyas"
	fmt.Println("name:", name)

	// variable initialization
	// syntax:
	//		var var_name type = initial_value
	var number int = 10
	fmt.Println("number:", number)

	// variable declaration without using type
	// if we dont specify the type then compiler will automatically check the type
	// depends upon the value passed.
	var age = 12
	fmt.Println("age:", age)

	// variable declaration using shorthand syntax
	// syntax:
	//		var_name := initial_value
	salary := 20000
	fmt.Println("salary:", salary)

	if len(os.Args) < 2 {
		fmt.Println("please run the program by passing command line argument")
		os.Exit(1)
	}
	// read value of the variable from command line
	var input string = os.Args[1]
	fmt.Println("Reading value from command line:", input)

}
