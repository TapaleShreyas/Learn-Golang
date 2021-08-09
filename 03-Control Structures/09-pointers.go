package main

import "fmt"

func main() {
	var a int = 10 // initializing value for a
	var b = &a     // assigning address of a to b
	var c = a      // assigning value of a to c
	fmt.Println("a =", a, "\nb =", *b, "\nc =", c)

	fmt.Println("\nmodify a now to 15")
	fmt.Println("b becomes 15 since it is pointing to address of 1")
	fmt.Println("c will still remains 10 since it is not referring to address")
	a = 15
	fmt.Println("a =", a, "\nb =", *b, "\nc =", c)

	fmt.Println("\nlet's try to change the value of b to 25")
	fmt.Println("a becomes 25 since b is pointing to address of a")
	fmt.Println("c will still remains 10 as it is")
	*b = 25
	fmt.Println("a =", a, "\nb =", *b, "\nc =", c)

	fmt.Println("\nnow change the value of c to 20")
	fmt.Println("a and b will remain same as it is")
	c = 20
	fmt.Println("a =", a, "\nb =", *b, "\nc =", c)

}
