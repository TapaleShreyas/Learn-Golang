package main

import "fmt"

func main() {

	// In Go, the traditional while true loop, found in many programming languages,
	// can done through the for keyword. Below are two alternative versions,
	// a for can work as an infinite loop without any parameters, or with a true boolean.
	// for {}
	// or
	// for true {}
	// we can conditionally break the infinite loop
	// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	i := 10
	for {
		if i > 50 {
			break
		}
		fmt.Println(i)
		i++
	}

	i = 10
	for true {
		if i > 15 {
			break
		}
		fmt.Println(i)
		i++
	}
}
