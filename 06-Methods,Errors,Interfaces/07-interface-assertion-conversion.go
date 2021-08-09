package main

import "fmt"

/*
	- type conversion vs type assertion
	- type assertion can only be used on interface.
	- type conversion can be done on any type.
	- j := i.(string) this assert the concrete type behind the interface is string and assign it to j.
	-
	- type conversion you are changing the type. i.e. int to float64, int to string
	- type assertion is delaing a underlying type and striping away an interface that wraps it.
	- if you are not sure about the type then you can use
		k, ok := i.(int)
		if you don't specify the ok, and type is not correct then code will panic.
*/

func main() {

	var i interface{}

	i = "Hello"
	j := i.(string)

	k, ok := i.(int)
	fmt.Println("j, k, ok :", j, k, ok)
}
