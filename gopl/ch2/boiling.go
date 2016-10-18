package main

import "fmt"

func main() {
	fmt.Printf("freezing point = %g°F or %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("boiling point = %g°F or %g°C\n", boilingF, fToC(boilingF))
}

// boilingF的声明可以放在使用后，仅针对包级变量有效
const freezingF, boilingF = 32.0, 212.0

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
