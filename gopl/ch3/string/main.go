package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const GoUsage = `Go is a tool for managing Go source code
	Usage:
	go command [arguments]
	...`
	fmt.Println(GoUsage)
	s1 := "世界"
	s2 := "\xe4\xb8\x96\xe7\x95\x8c"
	s3 := "\u4e16\u754c"
	s4 := "\U00004e16\U0000754c"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
