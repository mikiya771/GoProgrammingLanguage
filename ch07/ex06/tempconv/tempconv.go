package tempconv

import (
	"flag"
	"fmt"
)

type celsiusFlag struct{ Celsius }

const boilingF = 212.0
const zeroK = -273.15

type Celsius float64
type Kelvin float64
type Fahrenheit float64

func (c Celsius) String() string    { return fmt.Sprintf("%g C", c) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g F", f) }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Scanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func CToK(c Celsius) Kelvin {
	return Kelvin(c - zeroK)
}
func KToC(k Kelvin) Celsius {
	return Celsius(k + zeroK)
}
func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}
func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}
