package main

import (
	"bufio"
	"encoding/json"
	"strings"
	"fmt"
	"os"
	"strconv"
)

type t_dir struct {
	Dir			string		`json:"dir"`
	Files		[]string	`json:"files,omitempty"`
	Folders		[]t_dir		`json:"folders,omitempty"`
}

func isHackFile(filename string) bool {
	if len(filename) >= 5 && filename[len(filename) - 5:] == ".hack" {
		return true
	}
	return false
}

func countHackFiles(files []string) int {
	var n int

	for _, filename := range files {
		if isHackFile(filename) {
			n++
		}
	}
	return n
}

func allNestedFiles(json t_dir) int {
	var n int

	n = len(json.Files)
	for _, folder := range json.Folders {
		n += allNestedFiles(folder)
	}
	return n
}

// this is right answer!!! bro
func numberOfHackFiles(json t_dir) int {
	var n int = 0

	if countHackFiles(json.Files) != 0 {
		return allNestedFiles(json)
	}
	for _, folder := range json.Folders {
		n += numberOfHackFiles(folder)
	}
	return n
}

func read_input(scanner *bufio.Scanner, n int) string {
	var output []string

	for i := 0; i < n; i++ {
		scanner.Scan()
		output = append(output, scanner.Text())
	}
	return strings.Join(output, "")
}

func main() {
	var (
		in *bufio.Scanner
		n_cases int
		n_lines int
		json_data t_dir
	)

	in = bufio.NewScanner(os.Stdin)
	in.Scan()
	n_cases, _ = strconv.Atoi(in.Text())
	for i := 0; i < n_cases; i++ {
		json_data = t_dir{}
		in.Scan()
		n_lines, _ = strconv.Atoi(in.Text())
		data := read_input(in, n_lines)
		json.Unmarshal([]byte(data), &json_data)
		fmt.Println(numberOfHackFiles(json_data))
	}	
}
