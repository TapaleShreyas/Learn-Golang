package main

import "fmt"

func main() {
	var b *int     // initialize b with value type present at memory address int
	fmt.Println(b) // it will return value of b
	// fmt.Println(*b) // runtime error: invalid memory address or nil pointer dereference

	var a = new(int)
	fmt.Println(a)  // it will return memmory address of a
	fmt.Println(*a) // it will return the value of a here default value 0

	fmt.Println("\nuse above concepts in program now")
	n := 50
	fmt.Println("value of a before calling function:", n)
	add10(&n)
	fmt.Println("value of a after calling function:", n)

	fmt.Println("\ncall to add10Fail function")
	n = 50
	fmt.Println("value of a before calling function:", n)
	add10Fail(&n)
	fmt.Println("value of a after calling function:", n)
}

func add10(n *int) {
	*n = *n + 10
}

func add10Fail(n *int) {
	n = new(int) // assigning new address to a
	*n = *n + 10 // changes are being made to value present at new address
}
