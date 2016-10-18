package main

import (
	"fmt"
	"os"
)

func main() {
	var sep, s string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
