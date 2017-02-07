package main

import "fmt"

func asParam(f func()) {
	f()
}

func returnAdd() func(x, y int) int {
	return func(x, y int) int {
		return x + y
	}
}

func main() {
	// call anonymous functon
	func(s string) {
		fmt.Println(s)
	}("Hello, anonymous func")

	// value to var
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	// as params
	asParam(func() {
		fmt.Println("Hello, as param")
	})

	// as return
	fmt.Println(returnAdd()(3, 4))
}
