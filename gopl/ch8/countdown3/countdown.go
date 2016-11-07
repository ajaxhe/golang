package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// do nothing
		case <-abort:
			fmt.Println("Launch abort!")
			return
		default:
			// do something
			// will get problem
			fmt.Println("test")
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
