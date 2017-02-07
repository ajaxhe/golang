package main

import "fmt"

func test(s string, a ...int) {
	for i := range a {
		a[i] += 100
	}
	fmt.Printf("%T, %v\n", a, a)
}

func main() {
	test("abc", 1, 2, 3, 4)
	var s []int = []int{1, 2, 3}
	test("ajaxhe", s[:]...)
	fmt.Println(s)
}
