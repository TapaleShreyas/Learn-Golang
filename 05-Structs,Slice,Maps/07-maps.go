package main

import "fmt"

/*
	- a map is a data type that associates a value of one data type to a value of another data type.
	- map works on key value pair like maps in other langauge (e.g. java)
	- Syntax:
		m := make (map[string]int)
	- type of the value anything you want it to be
	- but key can be a slice, map, function, or it can be a type that contains slice, map or function.
	- maps are reference types
*/

func main() {

	cities := map[int]string{
		412301: "saswad",
		411043: "dhankawadi",
		411009: "sahakarnagar",
		415001: "hadapsar",
	}
	fmt.Println("cities map:", cities)

	// get city with pincode 412301
	fmt.Println("city with pincode 412301 is:", cities[412301])

	// get city with pincode 412000
	fmt.Println("city with pincode 412000 is:", cities[412000])

	fmt.Println("\ncheck element exists or not?")
	if city, ok := cities[412000]; ok {
		fmt.Println(city, "with pincode 412000 is present in map")
	} else {
		fmt.Println("city with pincode 412000 is not present in map")
	}

	fmt.Println("\nIterate over map")
	for key, value := range cities {
		fmt.Println(key, "is having city", value)
	}
	fmt.Println("\nGet keys(city) of the map")
	keys := make([]int, 0, len(cities))
	// map can return only values also
	for key := range cities {
		keys = append(keys, key)
	}
	fmt.Println("only cities from maps:", keys)

	fmt.Println("\ndelete city with pincode 411043")
	delete(cities, 411043)
	fmt.Println("cities after delete:", cities)

	fmt.Println("\nMaps are referene types")
	cities2 := cities
	cities2[411043] = "Dhankawadi"
	fmt.Println("cities2:", cities2)
	fmt.Println("cities:", cities)

}
