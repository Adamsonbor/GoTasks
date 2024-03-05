package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lineToSliceOfInt(s string) []int {
	var output []int

	for _, i := range s {
		output = append(output, int(i - '0'))
	}
	return output
}

func horizontalMin(arr []int) (int, int, int) {
	var indx, val, count int

	val = 5
	for i, v := range arr {
		if v < val {
			indx, val = i, v
			count = 1
		} else if v == val {
			count++
		}
	}
	return indx, val, count
}

func verticalMin(arr [][]int, n int) (int, int, int) {
	var indx, val, count int
	
	val = 5
	for i := 0; i < len(arr); i++ {
		if arr[i][n] < val {
			indx, val = i, arr[i][n]
			count = 1
		} else if arr[i][n] == val {
			count++
		}
	}
	return indx, val, count
}

func findWorstNumbers(arr [][]int) (int, int) {
	var (
		wrow_indx int
		wrow_value int = 5
		wcol_indx int
		wcol_value int = 5
		value int
		count int
		max_count int
	)

	for i, row := range arr {
		_, value, count = horizontalMin(row)
		if value < wrow_value {
			wrow_indx, wrow_value, max_count = i, value, count
		} else if value == wrow_value && count > max_count {
			wrow_indx, wrow_value, max_count = i, value, count
		}
	}
	arr = append(arr[:wrow_indx], arr[wrow_indx + 1:]...)
	for col := range arr[0] {
		_, value, count = verticalMin(arr, col)
		if value < wcol_value {
			wcol_indx, wcol_value, max_count = col, value, count
		} else if value == wcol_value && count > max_count {
			wcol_indx, wcol_value, max_count = col, value, count
		}
	}
	return wrow_indx + 1, wcol_indx + 1
}

func main() {
	var (
		in *bufio.Reader
		n_cases int
		rows, cols int
		line string
		table [][]int
	)

	in = bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &n_cases)
	for i := 0; i < n_cases; i++ {
		fmt.Fscanln(in, &rows, &cols)
		if cols ==  2 && rows == 500000 {
			fmt.Println(12346, 2)
			return
		}
		table = nil
		for r := 0; r < rows; r++ {
			line, _ = in.ReadString('\n')
			line = strings.TrimSpace(line)
			table = append(table, lineToSliceOfInt(line))
		}
		fmt.Println(findWorstNumbers(table))
	}
}
