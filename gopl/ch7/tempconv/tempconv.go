package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string { return fmt.Sprintf("%gºC", c) }

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

func (f *Celsius) Set(s string) error {
	var unit string
	var val float64
	fmt.Sscanf(s, "%f%s", &val, &unit)
	switch unit {
	case "C", "ºC":
		*f = Celsius(val)
		return nil
	case "F", "ºF":
		*f = FToC(Fahrenheit(val))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := Celsius(value)
	flag.CommandLine.Var(&f, name, usage)
	return &f
}
