package main

func test() (z int) {
	defer func() {
		println("defer:", z)
		z += 100
	}()
	return 100
}

func main() {
	println("test:", test())
}
