package main

import (
	"bufio"
	"fmt"
	"os"
)

type t_robot struct {
	name byte
	x int
	y int
}

func intAbs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func findRobot(field [][]byte, name byte) (t_robot, error) {
	for y, row := range field {
		for x, col := range row {
			if col == name {
				return t_robot{name:name, x: x, y: y}, nil
			}
		}
	}
	return t_robot{}, fmt.Errorf("no robot A")
}

func goToTop(field [][]byte, r t_robot) {
	var n byte = r.name + 32
	if r.y > 0 && field[r.y - 1][r.x] == '#' && r.x > 0 {
		r.x--
		field[r.y][r.x] = n
	}

	for i := r.y - 1; i >= 0; i-- {
		field[i][r.x] = n
	}
	for i := r.x - 1; i >= 0; i-- {
		field[0][i] = n
	}
}

func goToBot(field [][]byte, r t_robot) {
	var n byte = r.name + 32

	if r.y < len(field) - 1 && field[r.y + 1][r.x] == '#' && r.x < len(field[0]) - 1 {
		r.x++
		field[r.y][r.x] = n
	}

	for i := r.y + 1; i < len(field); i++ {
		field[i][r.x] = n
	}
	for i := r.x + 1; i < len(field[0]); i++ {
		field[len(field) - 1][i] = n
	}
}

func calcPath(field [][]byte) {
	var (
		A, B t_robot
		n int
	)

	A, _ = findRobot(field, 'A')
	B, _ = findRobot(field, 'B')
	n = (A.x - B.x) + (A.y - B.y)
	if n  >= 0 {
		goToTop(field, B)
		goToBot(field, A)
	} else {
		goToTop(field, A)
		goToBot(field, B)
	}
}

func main() {
	var (
		in *bufio.Reader
		n_cases int
		rows, cols int
		line string
		field [][]byte
	)

	in = bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n_cases)
	for i := 0; i < n_cases; i++ {
		fmt.Fscan(in, &rows, &cols)
		field = nil
		for r := 0; r < rows; r++ {
			fmt.Fscan(in, &line)
			field = append(field, []byte(line))
		}
		calcPath(field)
		for _, bytes := range field {
			fmt.Println(string(bytes))
		}
	}
}
