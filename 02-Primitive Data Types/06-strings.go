package main

import "fmt"

func main() {

	// strings are sequence of bytes
	// strings are immutable in Go
	// String literals should not contains the ascii characters

	var welcomeMsg string
	welcomeMsg = "Welcome to the Go programming"

	// different ways to print string
	fmt.Println("welcomeMsg =", welcomeMsg)
	fmt.Println("lenght of string:", len(welcomeMsg))
	fmt.Println("substring [0:7] =", welcomeMsg[0:7])
	fmt.Println("substring [0:7] length=", len(welcomeMsg[0:7]))
	fmt.Println("substring [:7] =", welcomeMsg[:7])
	fmt.Println("substring [:7] length=", len(welcomeMsg[:7]))
	fmt.Println("substring [8:29] =", welcomeMsg[8:29])
	fmt.Println("substring [8:29] length=", len(welcomeMsg[8:29]))
	fmt.Println("substring [8:] =", welcomeMsg[8:])
	fmt.Println("substring [8:] length =", len(welcomeMsg[8:]))

	// string contains newline or tab characters
	var str = "\nHello, \n\t\"world!\" with a backslash \\"
	fmt.Println(str)

	// raw string
	// raw string always enclosed in stray keys(``)
	// it is very helpful when we are reading values from UI or database and that contains new lines, tabs etc.
	var rawString = `Hello, 
	"world!" with a backslash \`
	fmt.Println("\nrawString:\n", rawString)

}
