package main

import "fmt"

/*
	- structs are not objects
	- no inheritance in structs unlike objects.
	- struts in GO are value type
	- embeding a struct
	- embedding provides delegation

*/

type Employee struct {
	empNo           uint
	salary          float32
	department      Department
	PersonalDetails // Embeding PersonalDetails struct
}

type PersonalDetails struct {
	name    string
	city    string
	pincode uint
}

type Department struct {
	deptId        uint
	deptName      string
	noOfEmployees uint
}

func main() {
	shreyas := PersonalDetails{
		name:    "Shreyas",
		city:    "Pune",
		pincode: 411043,
	}
	fmt.Println("shreyas:", shreyas)

	developers := Department{
		deptId:        1001,
		deptName:      "software development",
		noOfEmployees: 20,
	}
	fmt.Println("developers:", developers)

	employee := Employee{
		empNo:           404,
		salary:          10000,
		department:      developers,
		PersonalDetails: shreyas,
	}

	fmt.Println("employee:", employee)

	// To access department details you need to follow the hierarchie since you have not embeded it
	fmt.Println("department name: ", employee.department.deptName)
	employee.department.noOfEmployees = 30
	fmt.Println("employee:", employee)

	// Since the Personal Details in embeded into Employeed
	// You can access its members directly from employee
	fmt.Println("PersonalDetails shreyas:", employee.name)
	employee.pincode = 411009
	fmt.Println("employee:", employee)
}
