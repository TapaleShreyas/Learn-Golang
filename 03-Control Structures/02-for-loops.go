package main

import "fmt"

func main() {
	// simple for loop
	fmt.Println("simple for loop prints 1 to 10 numbers")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// for loop with break
	fmt.Println("\nfor loop with break statement prints 1 to 5 numbers")
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	// for loop with continue
	fmt.Println("\nfor loop with continue statement prints 5 to 10 numbers")
	for i := 1; i <= 10; i++ {
		if i > 5 {
			fmt.Println(i)
		}
		continue
	}

	// there is no while loop in go ut you can achive the same using for loop
	fmt.Println("\nachieve while loop behaviour using for loop")
	j := 5
	for j >= 1 {
		fmt.Println(j)
		j--
	}

}
