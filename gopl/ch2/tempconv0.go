package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	// var f float64 = 100
	fmt.Println(AbsoluteZeroC > 10)
	// fmt.Println(BoilingC > f) // false
	fmt.Printf("%s\n", AbsoluteZeroC)
	fmt.Println(FreezingC.String())
	fmt.Printf("%v\n", BoilingC)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
