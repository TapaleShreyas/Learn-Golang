package main

import "fmt"

func main() {

	// you can use range for loop to iterate over array, slice, maps.
	// syntax
	//		for offser_var, value_var := range var_to_iterate {
	//		}
	//  it return 2 elements, firts is offset of the element and second is the actual element
	// offsets are majored in byte not in characters

	arr := [5]int{11, 78, 34, 99, 123}
	for offset, value := range arr {
		fmt.Println(value, "present at offset", offset)
	}

	// range for loop on string
	// string is nothing but the sequence of bytes.

	fmt.Println("\nrange for loop on string")
	welcomeString := "Hello World"
	for idx, val := range welcomeString {
		fmt.Println(val, "is present at offset", idx)
	}

	// range for loop on emoji string
	// here we can see each emoji is taking 4 bytes and space is taking 1 byte

	fmt.Println("\nrange for loop on emoji string")
	emojiString := "👋 🌍"
	for idx, val := range emojiString {
		fmt.Println(string(val), "is present at offset", idx)
	}

}
