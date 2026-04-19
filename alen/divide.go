package main

import "fmt"

func divide(i int) (float64, error) {
	if i == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return float64(10) / float64(i), nil
}
