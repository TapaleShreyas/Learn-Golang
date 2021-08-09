package main

import "fmt"

/*
	- All function calls in GO are called by value.
		(copy of a value of a variable passed to a function not a reference)
	- you can declare functions inside a function. make sure inner function should be anonymous.
	- you can pass function to another function
		example: http.HandleFunc("/hello", printHello)
	- you can return a function from a function
*/

// function to add 1
func add(n int) int {
	return n + 1
}

// you can return function from function
// it adds 2 numbers
func addNumbers(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func main() {
	// you can create reference to a function and then can call the actual function using reference
	addOne := add
	fmt.Println("function call to add one:", addOne(1))

	// Anonymous function or function inside function
	addTwo := func(n int) int {
		return n + 2
	}
	fmt.Println("addTwo is a refernce to anonymous function:", addTwo(4))

	fmt.Println("\nfunction that returns a function")
	addFive := addNumbers(5)
	fmt.Println("function that adds 5 and 10 together:", addFive(10))

}
