package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("%v: %v\n", i, pc[i])
	}
}

func main() {
	fmt.Println("vim-go")
}
