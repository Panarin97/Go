package main

import (
	"fmt"
	"strconv"
	"sort"
)

// a function in order to check whether the input is an actual number
func is_number(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

// a function to read the input
func ask_for_input() string{
	var input_str string
	fmt.Printf("Enter an int (to exit press X): ")
	fmt.Scan(&input_str)

	return input_str
}

func main() {
	slice := make([]int, 3)
	var flag int = 1
	var num_of_inputs = 0

	for flag == 1 {
		var input string
		input = ask_for_input()

		if input == "X" {
			flag = 0

		} else if is_number(input) {
			var int_num int
			int_num, _ = strconv.Atoi(input)

			num_of_inputs = num_of_inputs + 1

			switch {

			case num_of_inputs < 4:
				slice[0] = int_num
				sort.Ints(slice)
				fmt.Println(slice)
				
			case num_of_inputs >= 4:
				slice = append(slice, int_num)
				sort.Ints(slice)
				fmt.Println(slice)
			}
		} else {
			fmt.Println("Invalid input")
		}
	}
}