package main

import "fmt"

type user struct {
	name string
	age  byte
}

func main() {
	var users [2]user
	fmt.Printf("%#v\n", users)
}
