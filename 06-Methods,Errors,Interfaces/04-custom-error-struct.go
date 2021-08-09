package main

import (
	"fmt"
	"os"
	"strconv"
)

type MyError struct {
	A            int
	B            int
	ErrorMessage string
}

func (_error *MyError) Error() string {
	return fmt.Sprintf("values %d and %d produced error: %s", _error.A, _error.B, _error.ErrorMessage)
}

func main() {
	// commands to run this program are
	// go run 04-custom-error-struct.go 12 0
	// go run 04-custom-error-struct.go 12 3
	// go run 04-custom-error-struct.go 12 sh

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
		return 0, 0, &MyError{A: a, B: b, ErrorMessage: "Can not divide by zero"}
	}
	return a / b, a % b, nil
}
