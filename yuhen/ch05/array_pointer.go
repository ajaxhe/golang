package main

import "fmt"

func main() {
	a := [...]int{1, 2}
	p := &a
	aa := a

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", &aa)

	p[1] += 10
	aa[1] += 10

	println(a[1])
}
