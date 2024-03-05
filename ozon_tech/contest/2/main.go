package main

import (
	"os"
	"fmt"
	"math"
	"bufio"
)

func main() {
	var stdin *bufio.Reader
	var stdout *bufio.Writer
	var n_cases int
	var n_items int
	var discount_percentage float64
	var discount float64
	var price float64
	var total_lost float64

	stdin = bufio.NewReader(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
	fmt.Fscan(stdin, &n_cases)
	for i := 0; i < n_cases; i++ {
		fmt.Fscan(stdin, &n_items, &discount_percentage)
		total_lost = 0
		for j := 0; j < n_items; j++ {
			fmt.Fscan(stdin, &price)
			discount = price * discount_percentage / 100
			total_lost += discount - float64(int(discount))
		}
		fmt.Fprintf(stdout, "%.2f\n", math.Floor(total_lost * 1000) / 1000)
		stdout.Flush()
	}
}
