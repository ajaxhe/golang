package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("start")
    go process()
    time.Sleep(time.Millisecond * 10)
    fmt.Println("dnoe")
}

func process() {
    fmt.Println("processing")
}
