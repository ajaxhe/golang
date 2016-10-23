package main

import "fmt"

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func (list *IntList) Insert(value int) {
	node := &IntList{value, nil}
	node.Tail = list.Tail
	list.Tail = node
}

func (list *IntList) Print() {
	for ; list != nil; list = list.Tail {
		fmt.Printf("%d ", list.Value)
	}
}

func main() {
	root := &IntList{0, nil}
	for i := 1; i < 10; i = i + 1 {
		root.Insert(i)
	}
	root.Print()
	fmt.Println()
	sum := root.Sum()
	fmt.Println(sum)
}
