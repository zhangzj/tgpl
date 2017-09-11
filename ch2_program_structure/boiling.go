package main

import "fmt"

// Boiling prints the boiling point of water

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g˚F or %g˚C\n", f, c)
	// output
	// boiling point = 212 ˚F or 100 ˚C
}
