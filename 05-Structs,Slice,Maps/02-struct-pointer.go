package main

import "fmt"

type Student struct {
	name       string
	rollNumber int
}

func main() {
	var shreyas = Student{name: "Shreyas"}
	var student Student

	student = shreyas
	fmt.Println("student:", student)

	// try to modify values of student now
	// shreyas struct values won't get change here
	// because both are having different memory addresses.
	student.rollNumber = 34
	fmt.Println("student:", student)
	fmt.Println("shreyas:", shreyas)

	fmt.Println("\n\nusing pointers\n")

	var john = Student{name: "John", rollNumber: 101}
	fmt.Println("john:", john)

	var randomStudent *Student

	randomStudent = &john
	fmt.Println("randomStudent:", randomStudent)

	// try to change value of randomStudent
	randomStudent.rollNumber = 89
	fmt.Println("randomStudent:", randomStudent)
	fmt.Println("john:", john)

}
