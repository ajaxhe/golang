package main

type color byte

const (
	black color = iota
	red
	blue
)

func test(c color) {
	println(c)
}

func main() {
	test(red)
	test(100)

	x := 2
	//test(x)
	test(color(x))
}
