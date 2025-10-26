package main

import (
	"errors"
	"fmt"
)

type customError struct {
	code    int
	message string
	er      error
}

// Error returns the error message. Implementing Error() error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.er)
}

// Function that returns a custom error
// func doSomething() error {
//	return &customError{
//		code:    500,
//		message: "Something went wrong!",
//	}
// }

func doSomething() error {
	if err := doSomethingElse(); err != nil {
		return &customError{
			code:    500,
			message: "something went wrong",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("internal error")
}

func main() {
	if err := doSomething(); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Operation completed successfully")
}
