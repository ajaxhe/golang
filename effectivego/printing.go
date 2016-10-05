package main

import (
    "fmt"
)

type T struct {
    a int
    b float64
    c string
}

func PrintStruct() {
    //t := T{7, -2.35, "abc\tefg"}
    t := &T{7, -2.35, "abc\tefg"}
    fmt.Printf("%v\n", t)
    fmt.Println((*t).a)
    fmt.Println(t.a)
}

func main() {
    PrintStruct()
}
