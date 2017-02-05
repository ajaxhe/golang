package main

import "fmt"

type color byte

func main() {
	var b byte = 10
	var ui uint8 = 20
	// byte = uint8?
	var n = b + ui
	var c color = 30
	n = n + c // mismatch type

	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", ui, ui)
	fmt.Printf("%T, %v\n", n, n)
	fmt.Printf("%T, %v\n", c, c)
}
