package main

import "fmt"

func main() {
	// run can be used to asign single literals
	str := "Hello, "
	// here we are assinging earth literal to rune
	var r rune = 127757

	// this is how you can convert rune to string and then concatenate
	str = str + string(r)
	fmt.Println(str)

	// another way of initializing run is you assign direct emoji to
	var r2 rune = '🌍'
	str = str + string(r2)
	fmt.Println(str)
}
