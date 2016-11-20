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
	fmt.Fprintf(&c, "hello, %s", name) // 这里会调用write函数
	fmt.Println(c)                     // 这里仅仅是输出c的值而已
}
