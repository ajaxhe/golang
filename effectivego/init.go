package main

import (
    "fmt"
    "os"
)

type ByteSize float64
const (
    // 通过iota迭代器初始化枚举置
    _ = iota
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func (b ByteSize) String() string{
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b / YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b / ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b / EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b / PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b / TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b / GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b / MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b / KB)
    }
    return fmt.Sprintf("%.2fB", b)
}

func GetEnvs() {
    var (
        home = os.Getenv("HOME")
        user = os.Getenv("USER")
        gopath = os.Getenv("GOPATH")
    )
    fmt.Println(home)
    fmt.Println(user)
    fmt.Println(gopath)
}

func init() {
    GetEnvs()
}

func InitConst() {
    var bz ByteSize = (1 << (10*3)) * 1.36
    //bz := ByteSize(10*YB)
    fmt.Printf("%v\n", bz)
}

func main() {
    InitConst()
}
