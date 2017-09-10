package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 print the text of each line that appears more than
// once in the standard input, precede by its count.

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "end" {
			break
		}
	}
	//
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
