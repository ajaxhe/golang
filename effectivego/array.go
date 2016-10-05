package main

import (
    "fmt"
)

func Sum(a *[3]float64) (sum float64) {
    // *a与a都是ok，这是为何？
    // for _, v := range *a {
    for _, v := range a {
        sum += v
    }
    return sum
}

func main() {
    //v := [3]float64{1.1,-2.2,3.3}
    v := [...]float64{1.1,-2.2,3.3}
    fmt.Println(Sum(&v))
}
