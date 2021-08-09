package main

import (
	"fmt"
	"reflect"
)

func main() {

	// type inference is nothing but you want to specify value of some type and
	// compiler automatically assumes it is of some other type.
	// for example you want to initialize value of type byte
	// and you have initialize as below

	var i = 10
	fmt.Println("type of i:", reflect.TypeOf(i))

	var j byte = 20
	i = j

	// If I try to assign value of type byte to i then compiler is showing me an error.
	// cannot use j (variable of type byte) as int value in
	// That's why it is considered to be the best practice to mention type along with variables declaration.

}
