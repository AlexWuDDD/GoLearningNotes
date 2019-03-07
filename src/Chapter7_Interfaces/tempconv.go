package main

import (
	"fmt"
	"tempconv"
)
type Value interface {
	String() string
	Set(string) error
}

type celsiusFlag struct {
	tempconv.Celsius
}

func(f *celsiusFlag) Sets(s string) error{
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit{
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
