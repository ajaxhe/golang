package main

import (
    "fmt"
)

func main() {
    var power int
    power = 9000
    //power := 1000
    //name, power := "ajaxhe", 1000
    power, name := 1000, "ajaxhe"
    fmt.Println(name, "power is:", power)
}
