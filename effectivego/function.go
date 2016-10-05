package main

import (
    "fmt"
)

func isDigit(i byte) bool {
    if '0' <= i && i <= '9' {
        return true
    } else {
        return false
    }
}

// 为何next_pos后面没有使用，也不会报错？
func NextInt(bs []byte, pos int) (value, next_pos int) {
    for ; pos < len(bs) && !isDigit(bs[pos]); pos++ {
    }
    for ; pos < len(bs) && isDigit(bs[pos]); pos++ {
        value = value * 10 + int(bs[pos]) - '0'
    }
    //next_pos = pos
    //return
    return value, pos
}

func ScanInt(bs []byte) {
    for pos := 0; pos < len(bs); {
        var value int
        value, pos = NextInt(bs, pos)
        //value, pos := NextInt(bs, pos)
        fmt.Println(value)
    }
}

func main() {
    bytes := []byte("1234a5bb122")
    ScanInt(bytes)
}
