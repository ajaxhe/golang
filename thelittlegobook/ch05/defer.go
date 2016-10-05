package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file_not_exist")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}
