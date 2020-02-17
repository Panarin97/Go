package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
type Id struct {
    fname string
    lname string
}

// ok so, I REALLY wanted to make their solution work, but since it is so cluttered, here is an easy solution on how you actually should read files...

func read_file(name string) []Id {
	Ids := make([]Id, 0)
	file, _ := os.Open(name)

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		names := strings.Split(scanner.Text(), " ")
		id := new(Id)
		id.fname = names[0]
		id.lname = names[1]
		Ids = append(Ids, *id)
	}
	return Ids
}

// a function to read the input
func ask_for_input() (string){
	var name string

	fmt.Printf("Enter a file name: ")
	fmt.Scan(&name)

	return name
}

func main() {
	name := ask_for_input()
	names := read_file(name)

	for _, id := range names {
		fmt.Printf("First name: %s | Last name: %s\n", id.fname, id.lname)
	}

}