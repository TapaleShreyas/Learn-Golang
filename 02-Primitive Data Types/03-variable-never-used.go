package main

import "fmt"

func main() {
	var marks int = 82
	fmt.Println("marks:", marks)

	// Go compiler gives you an error if we never used declared variables in program.
	// In below example variable economics is declared but it is never used in program thats why compiler is showing red
	// on line 13 you can see, I have reassigned the value to variable economics but
	// still compiler is showing an error because this variable is not used.
	// if we try to print that variable then error will get vanished since we are using it t print
	economics := marks
	economics = 90
	//fmt.Println("economics:", economics)
}
