package main

// import (
// 	"errors"
// 	"fmt"
// )

// // In Go, the 'error' type is an interface with a single method: Error() string.
// // Any type that implements this method satisfies the error interface and can be used as an error.
// // The standard library's errors.New returns a value of type error with a simple error message.

// // divide divides two integers and returns an error if the denominator is zero.
// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		// Return a new error if division by zero is attempted.
// 		return 0, errors.New("cannot divide by zero")
// 	}
// 	// No error, return the result and nil error.
// 	return a / b, nil
// }

// func ErrorExample() {
// 	// Try dividing 10 by 0 to trigger an error.
// 	result, err := divide(10, 0)
// 	if err != nil {
// 		// Handle the error by printing it.
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	// If no error, print the result.
// 	fmt.Println("Result:", result)
// }

// CustomDivideError is a custom error type that implements the error interface.
// type CustomDivideError struct {
// 	Dividend int
// 	Divisor  int
// 	Msg      string
// }

// // Error implements the error interface for CustomDivideError.
// func (e *CustomDivideError) Error() string {
// 	return fmt.Sprintf("cannot divide %d by %d: %s", e.Dividend, e.Divisor, e.Msg)
// }

// // divide divides two integers and returns an error if the denominator is zero.
// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		// Return a custom error if division by zero is attempted.
// 		return 0, &CustomDivideError{
// 			Dividend: a,
// 			Divisor:  b,
// 			Msg:      "division by zero",
// 		}
// 	}
// 	// No error, return the result and nil error.
// 	return a / b, nil
// }

// func ErrorExample() {
// 	// Try dividing 10 by 0 to trigger an error.
// 	result, err := divide(10, 0)
// 	if err != nil {
// 		// Handle the error by printing it.
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	// If no error, print the result.
// 	fmt.Println("Result:", result)
// }

// Error is a interface that represents an error with a message and a code.
// A type can be an error and as well as fulfill the error interface.
// Errors package provides simple error handling primitives.
