package main

import "fmt"

/*
	- slice are similar to array but without fixed length
	- operations on slice are similar to that of array
	- syntax
			var_name := make()(t Type, size ...IntegerType)
			example:   var1 := make([]int, 0)
	- The size specifies the length. The capacity of the slice is equal to its length.
	- A second integer argument may be provided to specify a different capacity;
	- it must be no smaller than the length.
*/

func main() {
	// declare slice of integers
	stringSlice := make([]string, 0)
	fmt.Println("length of string slice:", len(stringSlice))
	// stringSlice[0] = "Shreyas" // runtime error: index out of range [0] with length 0

	// add element to slice
	stringSlice = append(stringSlice, "Shreyas")
	fmt.Println("stringSlice:", stringSlice)

	fmt.Println("now you can access the first element:", stringSlice[0])

	intSlice := make([]int, 2)
	fmt.Println("length of intSlice:", len(intSlice))
	fmt.Println("slice is of length 2, you can access 1st element directly:", intSlice[0])

	fmt.Println("\nInitialize slice while declaring")
	cities := []string{"goa", "pune", "nashik", "satara"}
	fmt.Println("cities: ", cities)

	fmt.Println("print first 2 cities from slice:", cities[0:2])
	fmt.Println("print last 2 cities from slice:", cities[len(cities)-2:])

	fmt.Println("\nIterate over slice(cities) using for range loop")
	for offset, element := range cities {
		fmt.Println(element, "present at", offset, "index")
	}

}
