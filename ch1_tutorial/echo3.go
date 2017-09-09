package main

import (
	"fmt"
	"strings"
	"os"
	"time"
)

func main() {

	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%s elapsed.", time.Since(start).Nanoseconds())
	fmt.Println(os.Args[1:])
}
