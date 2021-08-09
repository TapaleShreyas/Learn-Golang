package main

import "fmt"

/*
	- Closures : A local function that has access to the variable that exists in
				 local environment in which it is declared.
				-local variable values can aso be modified using closures.
*/
func main() {

	// anonymous function can access the local variables in which it is declared
	// but the varibles defined in anonymous function are not accesible outside that function
	a := 10
	addition := func(b int) int {
		return a + b
	}
	fmt.Println("addition 10 and 20 is:", addition(20))

	// below line will give you an error.
	// undeclared name: b compiler UndeclaredName
	//fmt.Println("trying to access variable b:", b)

	// we can modify the local variable values also from within the function
	num := 400
	fmt.Println("local variable value before closure is:", num)
	ModifyVariable := func(n int) {
		num = num + n
	}
	ModifyVariable(500)
	fmt.Println("local variable value after calling closure is:", num)
}
