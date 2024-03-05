package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	var a, b int

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &a, &b)
	fmt.Fprint(out, a - b)
}
