package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])


	start := 2
	for i, arg := range os.Args[start:] {
		fmt.Println(i+start, arg)
	}
}
