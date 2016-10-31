package main

import "fmt"

type ByteCounter int

func (p *ByteCounter) Write(bytes []byte) (int, error) {
	*p += ByteCounter(len(bytes))
	return len(bytes), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("ajaxhe"))
	fmt.Println(c)
	var name = "dabao"
	c = 0
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
	fmt.Println("vim-go")
}
