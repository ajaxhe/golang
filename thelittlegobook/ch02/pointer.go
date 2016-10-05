package main

import (
    "fmt"
)

type Saiyan struct {
    Name string
    Power int
}

func Super(s *Saiyan) {
    s.Power += 10000
}
func (s *Saiyan) Super2(){
    s.Power += 1000
}
func main() {
    goku := &Saiyan{"ajaxhe", 9000} 
    Super(goku)
    fmt.Println(goku)
    goku.Super2()
    fmt.Println(goku)
}
