package main

import (
	"errors"
	"fmt"
)

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func log(x int, err error) {
	fmt.Println(x, err)
}

func test() (int, error) {
	return div(5, 0)
}

func test2() (int, string, error) {
	return 0, "", nil
}

func test3() (int, s string, e error) {
	return 0, "", nil
}

func main() {
	log(test())
	fmt.Println(test2())
	fmt.Println(test3())
}
