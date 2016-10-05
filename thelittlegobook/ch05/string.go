package main

import (
    "fmt"
)

func main() {
    str := "汉语"
    fmt.Println(len(str))
    fmt.Println(len("何露凡"))
    fmt.Println(len("bytes"))
    for i := 0; i < len(str); i++ {
        fmt.Println(str[i])
    }
    for idx, c := range(str) {
        fmt.Printf("idx:%d, val:%#U\n", idx, c)
    }
    byts := []byte(str)
    for _, c := range(byts) {
        fmt.Println(c)
    }
}
