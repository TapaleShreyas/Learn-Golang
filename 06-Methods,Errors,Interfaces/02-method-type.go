package main

import "fmt"

type Year int

// value receiver
func (year Year) isLeapYear() bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// value receiver
func (year Year) isEven() bool {
	return year%2 == 0
}

// reference receiver
func (year *Year) addTwoYears() Year {
	return *year + 2
}

func main() {
	var year Year = 2000
	fmt.Println("is Leap year 2000?:", year.isLeapYear())
	fmt.Println("is year 2000 is even?:", year.isEven())
	fmt.Println("add 3 more years to 2000 :", year.addTwoYears())

}
