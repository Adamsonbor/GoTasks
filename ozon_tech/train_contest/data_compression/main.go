package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func parse_input_string(input_string string) []int {
	var output []int

	for _, s := range strings.Split(input_string, " ") {
		number, _ := strconv.Atoi(s)
		output = append(output, number)
	}
	return output
}

func colc_range(input_numbers []int) int {
	var counter int = 0
	var sign int

	if len(input_numbers) == 1 { return counter }
	if (input_numbers[counter + 1] - 1) == input_numbers[counter] {
		sign = 1
	} else if (input_numbers[counter + 1] + 1) == input_numbers[counter] {
		sign = -1
	} else {
		return counter
	}
	for counter + 1 < len(input_numbers) && input_numbers[counter + 1] - sign == input_numbers[counter] {
		counter++
	}
	return counter * sign
}

func data_compresion(input_numbers []int) []int {
	var output []int
	var counter int = 0

	for i := 0; i < len(input_numbers); i++ {
		output = append(output, input_numbers[i])
		counter = colc_range(input_numbers[i:])
		if counter < 0 {
			i -= counter
		} else {
			i += counter
		}
		output = append(output, counter)
	}
	return output
}

func main() {
	var number_of_cases int
	var input_numbers []int
	var output_numbers []int
	var scanner *bufio.Scanner

	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number_of_cases, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < number_of_cases; i++ {
		scanner.Scan()
		scanner.Scan()
		input_numbers = parse_input_string(scanner.Text())
		output_numbers = data_compresion(input_numbers)
		fmt.Println(len(output_numbers))
		for _, number := range output_numbers {
			fmt.Print(number, " ")
		}
		fmt.Println()
	}
}

