package main

import (
	"fmt"
	"strconv"
)

// a function to read the input
func ask_for_input() string{
	var input_str string
	fmt.Printf("Enter an int (to exit press X, max 10 ints): ")
	fmt.Scan(&input_str)

	return input_str
}

// a function in order to check whether the input is an actual number
func is_number(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func Swap(slice []int, pos int) {
	var first int = slice[pos]
	var second int = slice[pos+1]

	var temp int = first
	slice[pos] = second
	slice[pos+1] = temp
}

func BubbleSort(slice []int) {
	length := len(slice)
	var swapped bool
	for i := 0; i < length-1; i++  {
		swapped = false
		for j := 0; j < length - i - 1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
				swapped = true
			}
		}
		// IF no two elements were swapped, break 
		if swapped == false {
			break
		}
	}
}


func main() {
	slice := make([]int, 0, 10)
	var flag int = 1
	var num_of_inputs = 0

	for flag == 1 {
		var input string

		//check for input capacity
		if num_of_inputs < 10 {
			input = ask_for_input()

		} else {
			flag = 0
			fmt.Println("Reached max capacity!")
		}

		if input == "X" {
			flag = 0 

		} else if is_number(input) {
			var int_num int
			int_num, _ = strconv.Atoi(input)
			num_of_inputs = num_of_inputs + 1

			slice = append(slice, int_num)
		} else {
			fmt.Println("Invalid input")
		}
	}
	BubbleSort(slice)
	fmt.Println(slice)
}