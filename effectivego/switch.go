package main

import (
    "fmt"
)

func Unhex(c byte) byte {
    switch {
        case '0' <= c && c <= '9':
            return c - '0'
        case 'a' <= c && c <= 'f':
            return c - 'a' + 10
        case 'A' <= c && c <= 'F':
            return c - 'A' + 10
    }
    return 0
}

func ShouldEscape(c byte) bool {
    switch c {
        case ' ', '?', '=', '#', '+':
            return true
        case '&':
            return true
    }
    return false
}

// Compare按字典顺序比较两个字节切片并返回一个整数
// 若a==b:return 0, 若a<b:return -1; 若a>b:return 1
func Compare(a, b []byte) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        switch {
            case a[i] < b[i]:
                return -1
            case a[i] > b[i]:
                return 1
        }
    }
    switch {
        case len(a) < len(b):
            return -1
        case len(a) > len(b):
            return 1
    }
    return 0
}

func main() {
    fmt.Println(Unhex('F'))
    fmt.Println(Unhex('0'))
    fmt.Println(ShouldEscape('+'))
    fmt.Println(ShouldEscape('1'))
    a := []byte("ajaxhe")
    b := []byte("ajax123")
    fmt.Println(Compare(a, b))
    fmt.Println(Compare(a, a))
    PrintType()
}
