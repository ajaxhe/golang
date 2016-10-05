package main

import (
    "fmt"
)

func DeferReturn() {
    fmt.Println("I am in DeferReturn")
    defer fmt.Printf("\n")
    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }
}

func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    DeferReturn()
    b()
}
