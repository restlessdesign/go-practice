package main

import (
	// "errors"
	"fmt"
)

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0 {
		// Why throw a generic error when you can create your own!
		// error := errors.New("Can’t divide by zero")
		// return dividend, error

		return dividend, &divideByZeroError{dividend}
	}
	return dividend / divisor, nil
}

// Types _______________________________________________________________________

type divideByZeroError struct {
	dividend float64
}

// Custom error structs need only to implement the Error method -- remember that
// structs implicitly implement an interface when their methods satisfy those
// defined by the interface!
func (e *divideByZeroError) Error() string {
	return fmt.Sprintf("Can’t divide %.f by zero", e.dividend)
}

// Main ________________________________________________________________________

func main() {
	fmt.Println(divide(5, 5))
	fmt.Println(divide(10, 5))
	fmt.Println(divide(10, 1))
	fmt.Println(divide(10, 0))
	fmt.Println()

	var last_val = -1.0
	numbers := []float64{0, 1, 2, 3, 4, 5, 10, 25}
	for _, cur_val := range numbers {
		result, err := divide(cur_val, last_val)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}

		last_val = cur_val
	}
}
