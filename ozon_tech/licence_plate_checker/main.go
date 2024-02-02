package main

import (
	"fmt"
	"bufio"
	"os"
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

func is_valid_licence_plate(licence_plate string) bool {
	return true
}

func find_licence_plate(input_string string) []string {
	var length int
	var output []string

	length = len(input_string)
	for i := 0; i < length; {
		if (length - i < 3) { return nil }
		result := is_valid_licence_plate(input_string[i:i+3])
		if result {
			output = append(output, input_string[i:i+3])
			i += 3
		}
	}
	return []string{}
}

func main() {
	var number_of_lines int
	var scanner *bufio.Scanner
	var input_strings []string

	scanner = bufio.NewScanner(os.Stdin)
	fmt.Scan(&number_of_lines)
	
	input_strings = read_input(number_of_lines, scanner)
	result := find_licence_plate(input_strings[0])
	if result != nil {
		fmt.Println("-")
	} else {
		for _, value := range result {
			fmt.Printf("%s ", value)
		}
		fmt.Println()
	}
}
