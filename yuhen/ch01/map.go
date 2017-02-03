package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 1
	m["c"] = 3
	x, ok := m["b"]
	fmt.Println(x, ok)
	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "a")
}
