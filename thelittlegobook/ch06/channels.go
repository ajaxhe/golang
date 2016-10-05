package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Worker struct {
    id int
}

func (w Worker) process(c chan int) {
    for {
        data := <- c
        fmt.Printf("worker %d got %d\n", w.id, data)
    }
}

func main() {
    c := make(chan int)
    for i := 0; i < 4; i++ {
        woker := Worker{id: i}
        go woker.process(c)
    }

    // write data
    for {
        c <- rand.Int()
        time.Sleep(time.Millisecond * 50)
    }
}
