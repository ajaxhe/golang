package main

var x int = 100

func main() {
	println(&x, x)
	x = 10
	println(&x, x)
}
