package main

import (
	"bytes"
	"fmt"
)

func main() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = new(bytes.Buffer)
	any = map[string]int{"one": 1}

	fmt.Println(any)

	fmt.Println("vim-go")
}
