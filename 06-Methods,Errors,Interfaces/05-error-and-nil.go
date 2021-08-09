package main

import "fmt"

/*
	- an error interface is considered nil only if it isn't associated with any underlying value,
	  not even nil underlying value.
	- you can define your own custom error class and provide a reference of it by assigning values.
*/

type MyError struct {
	A            int
	B            int
	ErrorMessage string
}

func (_error *MyError) Error() string {
	return fmt.Sprintf("values %d and %d produced error: %s", _error.A, _error.B, _error.ErrorMessage)
}

func reallyNil() error {
	var _error error
	fmt.Println("_error from reallyNil:", _error == nil)
	return _error
}

func notReallyNil() error {
	var _error *MyError
	fmt.Println("_error from notReallyNil:", _error == nil)
	return _error
}

func main() {
	_err := reallyNil()
	_myerr := notReallyNil()
	fmt.Println("in main, _err is nil:", _err == nil)
	fmt.Println("in main, _myerr is nil:", _myerr == nil)
}
