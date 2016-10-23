package main

import "fmt"

func main() {
	var anonyFuncs []func()
	slice := []string{"ajaxhe", "zhuzhu"}
	for _, n := range slice {
		name := n
		anonyFuncs = append(anonyFuncs, func() {
			fmt.Println(name)
		})
	}
	for _, fn := range anonyFuncs {
		fn()
	}
}
