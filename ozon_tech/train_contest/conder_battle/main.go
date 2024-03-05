package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_input(number_of_lines int, scanner *bufio.Scanner) []string {
	var input_strings []string

	if (scanner == nil) {
		panic("Scanner is nil")
	}
	for i := 0; i < number_of_lines; i++ {
		scanner.Scan()
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		input_strings = append(input_strings, scanner.Text())
	}
	return input_strings
}


func main() {
	var (
		scanner *bufio.Scanner
		input_strings []string
		number_of_lines int
	)

	scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	number_of_lines, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	input_strings = read_input(number_of_lines, scanner)

}
