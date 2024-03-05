package main

import (
	"bufio"
	"strconv"
	"os"
	"fmt"
)

func is_valid_number_of_ships(ships map[string]int) bool {
	return ships["1"] == 4 &&
		ships["2"] == 3 &&
		ships["3"] == 2 &&
		ships["4"] == 1
}

func calculate_ships(ships_string string) map[string]int {
	ships := make(map[string]int)
	for _, ship := range ships_string {
		switch ship {
		case '1':
			ships["1"]++
		case '2':
			ships["2"]++
		case '3':
			ships["3"]++
		case '4':
			ships["4"]++
		}
	}
	return ships
}

func main() {
	var inputs []string	
	var scanner bufio.Scanner

	scanner = *bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number_of_ships, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid number of ships")
		return
	}
	for i := 0; i < number_of_ships; i++ {
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("Error reading input")
			return
		}
		inputs = append(inputs, scanner.Text())
	}
	for _, ships_string := range inputs {
		ships := calculate_ships(ships_string)
		if is_valid_number_of_ships(ships) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
