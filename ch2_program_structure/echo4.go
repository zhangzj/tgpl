package main

import (
	"flag"
	"fmt"
	"strings"
)

// Echo4 print its command-line arguments


var n = flag.Bool("n", false, "omitting trailing new line")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
