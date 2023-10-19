package main

import (
	"fmt"
	"errors"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

func Divide(a, b float64) (float64, error)  {

	if b == 0 {
		// division by zero is not allowed
		return 0, errors.New("Invalid operation")
	}

	return a/b, nil

}
