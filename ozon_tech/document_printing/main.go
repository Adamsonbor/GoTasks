package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func parse_constraints(n_pages int, constraints []string) []bool {
	var constraints_bool []bool = make([]bool, n_pages + 1)
	var range_constraints []string

	for _, constraint := range constraints {
		range_constraints = strings.Split(constraint, "-")
		if len(range_constraints) == 1 {
			page, _ := strconv.Atoi(range_constraints[0])
			constraints_bool[page] = true
		} else {
			min, _ := strconv.Atoi(range_constraints[0])
			max, _ := strconv.Atoi(range_constraints[1])

			for i:=min; i<=max; i++ {
				constraints_bool[i] = true
			}
		}
	}
	return constraints_bool
}

func ranges_of_pages(constraints []bool) []string {
	var min, max int
	var range_of_pages []string

	for i := 1; i < len(constraints); {
		if !constraints[i] {
			min = i
			for i < len(constraints) && !constraints[i] {
				i++
			}
			max = i - 1
		} else {
			i++
			continue
		}
		if min == max {
			range_of_pages = append(range_of_pages, strconv.Itoa(min))
		} else {
			range_of_pages = append(range_of_pages, strconv.Itoa(min) + "-" + strconv.Itoa(max))
		}
	}
	return range_of_pages
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n_cases, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n_cases; i++ {
		scanner.Scan()
		n_pages, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		constraints := strings.Split(scanner.Text(), ",")
		constraints_bool := parse_constraints(n_pages, constraints)

		ranges := ranges_of_pages(constraints_bool)
		fmt.Println(strings.Join(ranges, ","))
	}
}
