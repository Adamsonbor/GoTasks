package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func parse_date(date string) (int, int, int, error) {
	var date_parts []string

	// Check if the date is in the correct format
	date_parts = strings.Split(date, " ")
	if len(date_parts) != 3 {
		return 0, 0, 0, fmt.Errorf("Invalid date format")
	}
	day, err := strconv.Atoi(date_parts[0])
	if err != nil {
		return 0, 0, 0, err
	}
	month, err := strconv.Atoi(date_parts[1])
	if err != nil {
		return 0, 0, 0, err
	}
	year, err := strconv.Atoi(date_parts[2])
	if err != nil {
		return 0, 0, 0, err
	}
	return day, month, year, nil
}

func is_valid_date(date string) bool {
	var months_length [12]int = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	day, month, year, err := parse_date(date)
	if err != nil {
		panic(err)
	}

	// update the length of February if it's a leap year
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		months_length[1] = 29
	}

	// Check if the date is valid
	if (month < 1 || month > 12) {
		return false
	}
	if (day < 1 || day > months_length[month - 1]) {
		return false
	}
	if (year < 1) {
		return false
	}
	return true
}

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
	var number_of_lines int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scan(&number_of_lines)
	input := read_input(number_of_lines, scanner)
	for _, date := range input {
		result := is_valid_date(date)
		if result {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
