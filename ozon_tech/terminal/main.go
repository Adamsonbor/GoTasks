package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
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

func minOf(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

type terminal_buffer struct {
	buffer []string
	xcursor int
	ycursor int
}

func newTerminalBuffer() *terminal_buffer {
	return &terminal_buffer{buffer: []string{""}, xcursor: 0, ycursor: 0}
}

func (tb *terminal_buffer) addChar(char string) {
	str := tb.buffer[tb.ycursor]
	tb.buffer[tb.ycursor] = str[:tb.xcursor] + char + str[tb.xcursor:]
	tb.xcursor++
}

func (tb *terminal_buffer) moveUp() {
	if tb.ycursor > 0 {
		tb.ycursor--
	}
	tb.xcursor = minOf(tb.xcursor, len(tb.buffer[tb.ycursor]))
}

func (tb *terminal_buffer) moveDown() {
	if tb.ycursor < len(tb.buffer) - 1 {
		tb.ycursor++
	}
	tb.xcursor = minOf(tb.xcursor, len(tb.buffer[tb.ycursor]))
}

func (tb *terminal_buffer) moveLeft() {
	if tb.xcursor > 0 {
		tb.xcursor--
	}
}

func (tb *terminal_buffer) moveRight() {
	if tb.xcursor < len(tb.buffer[tb.ycursor]) {
		tb.xcursor++
	}
}

func (tb *terminal_buffer) moveStart() {
	tb.xcursor = 0
}

func (tb *terminal_buffer) moveEnd() {
	tb.xcursor = len(tb.buffer[tb.ycursor])
}

func (tb *terminal_buffer) newLine() {
	str := tb.buffer[tb.ycursor]
	// add new string
	tb.buffer = append(tb.buffer, "")
	// move the rest of the strings down
	copy(tb.buffer[tb.ycursor + 1:], tb.buffer[tb.ycursor:])
	// set string to the left of the cursor
	tb.buffer[tb.ycursor] = str[:tb.xcursor]
	// set string to the right of the cursor
	tb.buffer[tb.ycursor + 1] = str[tb.xcursor:]
	// move cursor down
	tb.ycursor++
	// set cursor to the start of the new line
	tb.xcursor = 0
}

func (tb *terminal_buffer) printBuffer() {
	for _, line := range tb.buffer {
		fmt.Println(line)
	}
}

func (tb *terminal_buffer) execute(ch string) {
	var commands = map[string]func(){
		"L": tb.moveLeft,
		"R": tb.moveRight,
		"U": tb.moveUp,
		"D": tb.moveDown,
		"N": tb.newLine,
		"B": tb.moveStart,
		"E": tb.moveEnd,
	}
	if command, ok := commands[ch]; ok {
		command()
	} else {
		tb.addChar(ch)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number_of_lines, _ := strconv.Atoi(scanner.Text())
	input_strings := read_input(number_of_lines, scanner)

	for _, str := range input_strings {
		tb := newTerminalBuffer()
		for _, ch := range str {
			tb.execute(string(ch))
		}
		tb.printBuffer()
		fmt.Println("-")
	}
}
