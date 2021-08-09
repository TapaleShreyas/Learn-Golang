package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// commands to run this program are
	// go run 03-errors.go 12 0
	// go run 03-errors.go 12 3
	// go run 03-errors.go 12 sh

	if len(os.Args) < 3 {
		fmt.Println("Provide 2 input integers")
		os.Exit(1)
	}

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	div, mod, err := divisionAndMod(int(a), int(b))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%d / %d == %d and %d %% %d == %d\n", a, b, div, a, b, mod)

}

// methods that takes 2 integers, and returns division, mod, error if any.
func divisionAndMod(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Can not divide by 0")
	}
	return a / b, a % b, nil
}
