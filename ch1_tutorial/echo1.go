package main

import (
	"fmt"
	"os"
)

func main() {

	var argsString, separator string
	separator = " "
	for i := 1; i < len(os.Args); i++ {
		argsString += separator + os.Args[i]
	}
	fmt.Println(argsString)
}
