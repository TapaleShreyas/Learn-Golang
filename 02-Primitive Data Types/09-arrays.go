package main

import "fmt"

func main() {
	var intArr1 [3]int
	intArr1[0] = 234
	intArr1[1] = 789
	intArr1[2] = 345
	fmt.Println("intArr1:", intArr1)
	fmt.Println("intArr1:", intArr1[:2])
	fmt.Println("intArr1:", intArr1[0:3])

	var intArr2 = [4]int{123, 456, 890, 4534}
	fmt.Println("intArr2:", intArr2)
	fmt.Println("intArr2:", intArr2[1:])
	fmt.Println("intArr2:", intArr2[2:4])

	var rawString = `You can't assign intArr2 to intArr1 or vice versa because both are having different size.
	intArr1 = intArr2 // cannot use intArr2 (variable of type [4]int) as [3]int value in assignment compiler IncompatibleAssign
	intArr2 = intArr1 // cannot use intArr1 (variable of type [3]int) as [4]int value in assignment compiler IncompatibleAssign
	In both the above case compiler will give you an error saying IncompatibleAssign`

	fmt.Println(rawString)
}
