package main

import "fmt"

func sum(vals ...int) (total_val int) {
	for _, val := range vals {
		total_val += val
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	vals := []int{1, 2, 3}
	fmt.Println(sum(vals...))
}
