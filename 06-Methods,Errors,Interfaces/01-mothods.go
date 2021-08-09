package main

import "fmt"

/*
	- provides object oriented features
	- value receiver and reference receiver
	- syntax:
		func (s struct_name) method_name() return_type {
			method body...
		}
	-
*/

// define struct
type Person struct {
	name  string
	city  string
	phone int64
	email string
}

// Get contact details
func (person Person) getContactDetails() string {
	return fmt.Sprintf("phone:%d, email=%s", person.phone, person.email)
}

// updating email of current struct
func (person *Person) updateEmailAddress(newEmail string) *Person {
	person.email = newEmail
	return person
}

func main() {
	person := Person{
		name:  "Shreyas",
		city:  "Pune",
		phone: 1231231,
		email: "abc@gmail.com",
	}
	fmt.Println("person:", person)
	fmt.Println("get contact details:", person.getContactDetails())

	fmt.Println("\nUpdate Email Address to xyz@yahoo.com")
	fmt.Println(person.updateEmailAddress("xyz@yahoo.com"))
	fmt.Println("person after email address update:", person)

	fmt.Println("\nChange city of return person")
	p := person.updateEmailAddress("xyz@yahoo.com")
	p.city = "Saswad"
	fmt.Println("p:", p)
	fmt.Println("person:", person)
	// above both p and Person are referring to same location
	// and hence if we change value of either of then it will get reflect for both
}
