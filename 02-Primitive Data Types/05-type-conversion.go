package main

import "fmt"

func main() {

	var firstSemisterResult int = 60
	var secondSemisterResult float32 = 58.23

	// compiler won't allow below line since go does not supports auto type conversion
	// fmt.Println("total result:", firstSemisterResult+secondSemisterResult)

	// if you want to add both the result then either of them needs to be converted to other's type
	// so that both variable is if same type
	var intResult int = firstSemisterResult + int(secondSemisterResult)
	fmt.Println("result in int:", intResult)

	// syntax for type conversion
	//  type(variable)

	var floatResult float32 = float32(firstSemisterResult) + secondSemisterResult
	fmt.Println("result in float:", floatResult)

	// GO don't support type conversion between same type byt different bits.
	// for example betwen int8 to int32

	var num1 int8 = 8
	var num2 int16 = 16
	// var addition = num1 + num2
	// invalid operation: mismatched types int8 and int16

	var addition = int16(num1) + num2
	fmt.Println("addition:", addition)

}
