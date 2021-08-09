package main

import "fmt"

func main() {
	number := 10

	// Simple if else structure
	// no need of braces for condition 
	// body must be enclosed in curly braces
	// else part is optional
	// if works only with boolean value
	// unlike java where it does not works for 0 and 1
	if number < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is greater than or equal to 10")
	}

	// sometimes we create variables inside if block and not able to access it in else block
	// to overcome this behavior go has introduced new thin where
	// you can create variables local to your if else block.
	// syntax
	//		if variable initialization ; condition {
	//		} else { }

	var marksTotal float64 = 619 // out of 650
	var outOf float64 = 650.0

	if percentage := (marksTotal / outOf) * 100; percentage > 85 {
		msg := fmt.Sprintf("%.2f is greater than 85", percentage)
		fmt.Println(msg)
	} else {
		msg := fmt.Sprintf("%.2f is less than or equal to 85", percentage)
		fmt.Println(msg)
	}

	// scope of the percentage variable is for if else block only.
	// undeclared name: percentage compiler UndeclaredName
	// fmt.Println("percentage: ", percentage)

	// multiple conditions in if
	// all conditions needs to be satisfied
	marks := 65.23
	if marks > 60 && marks < 75 {
		fmt.Println("first class.")
	}

	// multiple either or condition
	// any one condition needs to be satisfied
	if marks > 60 || marks > 65 {
		fmt.Println("first class")
	}

	// if else ladder
	if marks > 75 {
		fmt.Println("distinction")
	} else if marks > 60 {
		fmt.Println("first class")
	} else if marks > 40 {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}

}
