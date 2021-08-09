package main

import "fmt"

func main() {
	uniHello := "👋 🌍"
	fmt.Println(uniHello)

	// convert string to byte slice
	bytes := []byte(uniHello)
	fmt.Println("slice(bytes):", bytes)

	// convert string to rune slice
	runes := []rune(uniHello)
	fmt.Println("slice(rune):", runes)

	// Modify runes element
	runes[1] = 'A'
	fmt.Println(runes)

	// print runes element in string format
	for _, element := range runes {
		fmt.Print(string(element), " ")
	}

	fmt.Println("\nrunes in string now:", string(runes))
}
