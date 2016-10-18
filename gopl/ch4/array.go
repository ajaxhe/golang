package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func useSHA256() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func main() {
	// 在这种形式的数组字面值形式中，初始化索引的顺序是无关紧要的，而且没用到的索引可以省略，和前面提到的规则一样，未指定初始值的元素将用零值初始化。
	symbol := [...]string{USD: "$", RMB: "¥", EUR: "E", GBP: "G"}
	fmt.Println(symbol[RMB])

	r := [...]int{99: -1}
	fmt.Println(r[98], r[99])

	// fmt.Println(symbol == r) // mismatch type

	useSHA256()

	t := [...]byte{31: 255}
	fmt.Println(t)
	zero(&t)
	fmt.Println(t)
}
