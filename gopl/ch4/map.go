package main

import "fmt"

func mapInit() {
	args := make(map[string]int)
	var args2 map[string]int
	fmt.Println(args == nil)
	fmt.Println(len(args))
	args["ajaxhe"] = 28
	// args2["ajaxhe"] = 28
	fmt.Println(args2 == nil)
	fmt.Println(len(args2))
	arg, ok := args["zhu"]
	fmt.Println(arg, ok)
}

func main() {
	mapInit()
}
