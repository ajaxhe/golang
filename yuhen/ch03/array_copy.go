package main

import "fmt"

func main() {
	data := [3]int{10, 20, 30}
	for i, x := range data {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}

	for i, x := range data[:] {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}
}
