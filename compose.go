// compose
package main

import (
	"fmt"
)

type report struct {
	sol int
	temperature
	location
}
type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (r report) average() celsius {
	return r.temperature.average()
}

func main() {
	report := report{sol: 15, temperature: temperature{high: -1.0, low: -78.0}, location: location{-4.5895, 137.4417}}
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v C\n", report.temperature.average())
	fmt.Printf("average %v C\n", report.average())
	fmt.Printf("%v\n", report.high)
	report.high = 33
	fmt.Printf("%v C\n", report.temperature.high)
}
