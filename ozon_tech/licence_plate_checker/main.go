package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
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

func is_valid_licence_plate_second_type(licence_plate string) bool {
	if (unicode.IsLetter(rune(licence_plate[0])) &&
		unicode.IsDigit(rune(licence_plate[1])) &&
		unicode.IsDigit(rune(licence_plate[2])) &&
		unicode.IsLetter(rune(licence_plate[3])) &&
		unicode.IsLetter(rune(licence_plate[4]))) {
		return true
	}
	return false
}

func is_valid_licence_plate_first_type(licence_plate string) bool {
	if (unicode.IsLetter(rune(licence_plate[0])) &&
		unicode.IsDigit(rune(licence_plate[1])) &&
		unicode.IsLetter(rune(licence_plate[2])) &&
		unicode.IsLetter(rune(licence_plate[3]))) {
		return true
	}
	return false
}

func is_valid_licence_plate(licence_plate string) bool {
	return is_valid_licence_plate_first_type(licence_plate) || is_valid_licence_plate_second_type(licence_plate)
}

func find_licence_plate(input_string string) []string {
	var length int
	var output []string

	length = len(input_string)
	for i := 0; i < length; {
		if (length - i < 4) { return nil }
		result := is_valid_licence_plate_first_type(input_string[i:i+4])
		if result {
			output = append(output, input_string[i:i+4])
			i += 4
			continue
		}

		if (length - i < 5) { return nil }
		result = is_valid_licence_plate_second_type(input_string[i:i+5])
		if result {
			output = append(output, input_string[i:i+5])
			i += 5
			continue
		}

		return nil
	}

	return output
}

func main() {
	var number_of_lines int
	var scanner *bufio.Scanner
	var input_strings []string

	scanner = bufio.NewScanner(os.Stdin)
	fmt.Scan(&number_of_lines)

	input_strings = read_input(number_of_lines, scanner)
	for _, value := range input_strings {
		result := find_licence_plate(value)
		if result != nil {
			for _, value := range result {
				fmt.Printf("%s ", value)
			}
			fmt.Println()
		} else {
			fmt.Println("-")
		}
	}
}
