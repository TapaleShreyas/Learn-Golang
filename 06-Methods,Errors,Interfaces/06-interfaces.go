package main

import (
	"fmt"
	"math"
)

/*
	- An Interface is a type that lists a set of methods but provides no implementation.

	- Define an Interface Geometry
	- Declare 2 methos for the interface
	- Create struct which implements this interface
	- Provide definition for both the interface method
	- create one more struct with interface method definition
*/

// define an interface with 2 methods
type Geometry interface {
	Area() float64
	Perimeter() float64
}

// define first implementation struct for an interface
// note we have not declared or mentioned any methods for Rectangle here
type Rectangle struct {
	height, width float64
}

// provide interface Area() method definition using rectangle struct
func (rectangle Rectangle) Area() float64 {
	return rectangle.height * rectangle.width
}

// provide interface Perimeter() method definition using rectangle struct
func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.height + rectangle.width)
}

// define a second implementation struct for an interface
// note we have not declared or mentioned any methods for Rectangle here
type Circle struct {
	radius float64
}

// provide interface Area() method definition using rectangle struct
func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// provide interface Perimeter() method definition using rectangle struct
func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

// normal function which accepts interface and calculate its area and perimeter
func calculateAreaAndPerimeter(geometry Geometry) {
	fmt.Println(fmt.Sprintf("Area is: %.2f", geometry.Area()))
	fmt.Println(fmt.Sprintf("Perimeter is: %.2f", geometry.Perimeter()))
	checkUnderlyingType(geometry)
}

func main() {

	// initialize rectangle with height and width
	rectangle := Rectangle{
		height: 12.0,
		width:  23.5,
	}
	fmt.Println("Calculate Area and Perimeter for Rectangle with height and width", rectangle)
	// note we are pssing rectangle struct object as interface reference
	calculateAreaAndPerimeter(rectangle)
	interfaceAssertion(rectangle)

	// initialize circle with radius
	circle := Circle{radius: 5.0}
	fmt.Println("\nCalculate Area and Perimeter for Rectangle with radius", circle)
	// note we are pssing circle struct object as interface reference
	calculateAreaAndPerimeter(circle)
	interfaceAssertion(circle)

}

// interface assertion
func interfaceAssertion(geometry Geometry) {
	fmt.Println("\nCheck geometry is of type using boolean flag")
	rectangle, flag := geometry.(Rectangle)
	if flag {
		fmt.Println("geometry is of type Rectangle.", rectangle)
	}

	circle, flag := geometry.(Circle)
	if flag {
		fmt.Println("geometry is of type Circle.", circle)
	}

}

// check underlying type
func checkUnderlyingType(geometry Geometry) {
	switch geometry.(type) {
	case Rectangle:
		fmt.Println("geometry is of type Rectangle", geometry)
	case Circle:
		fmt.Println("geometry is of type Circle", geometry)
	default:
		fmt.Println("geometry is of some other type")
	}
}
