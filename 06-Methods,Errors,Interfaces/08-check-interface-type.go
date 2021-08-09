package main

import "fmt"

func checkType(obj interface{}) {
	switch obj.(type) {
	case int:
		fmt.Println("obj is of type int")
	case float64:
		fmt.Println("obj is of type float64")
	case string:
		fmt.Println("obj is of type string")
	default:
		fmt.Println("obj is of some other type")
	}
}

func main() {
	checkType("hello")
	checkType(10)
	checkType(10.78)

	var float32 float32 = 32.1
	checkType(float32)
}
