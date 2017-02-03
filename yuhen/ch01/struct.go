package main

import "fmt"

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func (u user) toString() string {
	return fmt.Sprintf("%+v", u)
}

func (m manager) Print() {
	fmt.Printf("%+v\n", m)
}

type Printer interface {
	Print()
}

func main() {
	var m manager
	m.name = "ajaxhe"
	m.age = 28
	m.title = "cto"

	fmt.Println(m)
	println(m.toString())

	var p Printer = m
	p.Print()
}
