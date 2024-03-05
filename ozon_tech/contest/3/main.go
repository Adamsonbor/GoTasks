package main

import (
	"os"
	"fmt"
	"bufio"
)

func isOk(line string) bool {
	var is_run bool = false
	var commands = map[byte]bool{
		'M': true,
		'D': true,
		'C': true,
		'R': true,
	}

	for i:=0;i<len(line);i++{
		_, ok := commands[line[i]]
		if !ok { return false }
		if is_run {
			if line[i] == 'M' {
				return false
			} else if line[i] == 'C' && len(line) > i + 1 && line[i + 1] == 'M' {
				is_run = false
			} else if line[i] == 'D' {
				is_run = false
			} else if line[i] == 'R' && len(line) > i + 1 && line[i + 1] == 'C' {
			} else {
				return false
			}
		} else {
			if line[i] == 'M' {
				is_run = true
			} else {
				return false
			}
		}
	}
	return !is_run
}

func main() {
	var (
		in *bufio.Reader
		n_cases int
		line string
	)

	in = bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n_cases)
	for i := 0; i < n_cases; i++ {
		fmt.Fscan(in, &line)
		if isOk(line) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
