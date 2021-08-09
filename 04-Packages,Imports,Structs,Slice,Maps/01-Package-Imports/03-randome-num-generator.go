package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// If you don't use this line then rand.Int31n will generate the same numbe every time.
	rand.Seed(time.Now().UnixNano())

	// It will generate the random number for you within the range
	num1 := rand.Int31n(10)
	num2 := rand.Int31n(10)

	// math.Max will give you the maximum number between the 2
	max := int(math.Max(float64(num1), float64(num2)))
	fmt.Println("First number:", num1)
	fmt.Println("Second number:", num2)
	fmt.Println(max, "is bigger")
}
