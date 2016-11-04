package main

import "fmt"

func SQLQuery(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case uint, int:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		} else {
			return "FALSE"
		}
	case string:
		return x
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func main() {
	var i int = 17
	var s string = "ajaxhe"
	var b bool = false
	var f interface{}
	var m map[string]int = nil
	fmt.Printf("%v\t%T\n", SQLQuery(i), i)
	fmt.Printf("%v\t%T\n", SQLQuery(s), s)
	fmt.Printf("%v\t%T\n", SQLQuery(b), b)
	fmt.Printf("%v\t%T\n", SQLQuery(f), f)
	fmt.Printf("%v\t%T\n", SQLQuery(m), m)
}
