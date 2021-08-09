package main

import "fmt"

/*
	- you can define as many function you want.(after or before main func, doesn't matter)
	- no function overloading is allowed in Go.
	- function can return multiple values
	-
*/

// func that takes 2 arguments and prints the addition
func addition(a int, b int) {
	fmt.Println("Addition of", a, "+", b, "is", a+b)
}

// func that returns addition
func getAddition(a int, b int) int {
	return a + b
}

// function to return multiple values
func divisionAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

// function to check changing the values of passed arguments.
func modifyVariables(n int, arr [3]int, str string) {
	n = n * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	str = str + " append"
	fmt.Println("inside modify function values are \nn=", n, "\narr=", arr, "\nstr=", str)
}

func main() {
	// making a function call
	addition(5, 10)

	addition := getAddition(10, 10)
	fmt.Println("Addition of 10 + 10 by calling getAddition function:", addition)

	fmt.Println("\n")
	div, remainder := divisionAndRemainder(10, 4)
	fmt.Println("div, remainder = ", div, remainder)

	div, _ = divisionAndRemainder(10, 4)
	fmt.Println("division=", div)

	_, remainder = divisionAndRemainder(100, 4)
	fmt.Println("remainder=", remainder)

	n := 10
	arr := [3]int{2, 4, 6}
	str := "hello"
	modifyVariables(n, arr, str)
	fmt.Println("After calling modify function, inside main function values are \nn=", n, "\narr=", arr, "\nstr=", str)
}
