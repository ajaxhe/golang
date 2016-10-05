package main

import (
    "fmt"
    "io"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello world")
}

func main() {
    http.HandleFunc("/hello", hello)
    err := http.ListenAndServe(":17000", nil)
    if err != nil {
        fmt.Println(err)
    }
}
