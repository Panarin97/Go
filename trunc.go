package main

import (
	"fmt"
	"strconv"
)

// a function in order to check whether the input is an actual number
func is_number(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func main() {
	var float_str string
	fmt.Printf("Enter a float: ")
	fmt.Scan(&float_str)

	// check input
	if is_number(float_str) {
		var float_num float64
		var err error

		if float_num, err = strconv.ParseFloat(float_str, 64); err == nil {
			fmt.Println(int(float_num))
		}

	} else {
		fmt.Println("Invalid input")
	}

	// the other solutions are to use the "math" package with Modf method or string conversion.

}
