package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
	"runtime"
	"sync"
)

// a function to read the input
func ask_for_input() []string {
	fmt.Printf("Enter a series of ints with spaces > ")

	consoleReader := bufio.NewReader(os.Stdin)
	input, _ := consoleReader.ReadString('\n')
	s := strings.Fields(input)

	return s
}

// a function to check if input is a set of ints
func check_input(s []string) bool{
	for _, num := range s {
		if _, err := strconv.Atoi(num); err != nil {
			return false
		}
	}
	return true
}

// a function to convert an slice of strings to ints
func convert_input(s []string) []int {
	s_1 := make([]int, 0)
	for _, num := range s {
		x, _ := strconv.Atoi(num)
		s_1 = append(s_1, x)
	}
	return s_1
}

// a function to divide the input into 4 rough categories
func partition_slice(input [] int) [][] int {
	output := make([][]int, 0)

	if len(input) <= 4 {
		for i, _ := range input {
			output = append(output, input[i:i+1])
		}

	} else {
		remainder := len(input) % 4
		length := (len(input) - remainder)/4

		for i := 0; i <= 3; i++  {
			output = append(output, input[i*length:length*(i+1)])
		}

		for n := 0; n <= remainder - 1; n++  {
			slice := make([]int, len(output[n]), len(output[n])+1)
			copy(slice, output[n])
			slice = append(slice, input[len(input)-remainder + n])

			output[n] = slice
		}

	}

	return output
}

// a function to print and sort the arrays
func sort_arrays(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Initial array: ")
	fmt.Println(slice)

	sort.Ints(slice)

	fmt.Printf("Sorted array: ")
	fmt.Println(slice)
	fmt.Println()
}

func main() {

	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup

	input_ints := make([]int, 0)

	input := ask_for_input()

	if check_input(input) {
		input_ints = convert_input(input)
		part_ints := partition_slice(input_ints)
		final := make([]int, 0)

		for i := 0; i <= len(part_ints)-1; i++ {
			wg.Add(1)
			go sort_arrays(part_ints[i], &wg)
			final = append(final, part_ints[i]...)
		}

		wg.Wait()

		sort.Ints(final)

		fmt.Printf("Final sorted array: ")
		fmt.Println(final)

	} else {
		fmt.Println("Invalid input")
	}
}