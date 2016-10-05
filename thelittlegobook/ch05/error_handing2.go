package main

import (
	"fmt"
	"io"
)

func main() {
	var input int
	_, err := fmt.Scan(&input)
	if err == io.EOF {
		fmt.Println("no more input!")
	} else {
		fmt.Println(input)
	}
}
