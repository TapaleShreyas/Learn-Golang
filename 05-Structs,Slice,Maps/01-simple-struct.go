package main

import "fmt"

type Class struct {
	Students int
	Division string
}

func main() {

	// Initialize empty struct with new keyword
	// It will return memroy address of struct to class1 variable
	var class1 = new(Class)
	fmt.Println("class1:", class1)

	// Initialize empty struct (using empty constructor)
	// by default it will initialize all the variables with its default values.
	// default value for int is 0, default value for string is empty.
	class := Class{}
	fmt.Println("class:", class)

	// Initialize with parameterized constructor
	// If you are not mentioning the names of the struct member then order is important
	// classB := Class{"B", 40} //won't work.
	classA := Class{40, "A"}
	fmt.Println("classA:", classA)

	// Initialize with names
	classB := Class{Students: 50, Division: "B"}
	fmt.Println("classB:", classB)

	// Initialize with only one member
	// It will asign uninitialized struct member with default value.
	classC := Class{Division: "C"}
	fmt.Println("classC:", classC)

	// Modify value
	classC.Students = 80
	fmt.Println("classC:", classC)

}
