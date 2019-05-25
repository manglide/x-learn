// kelvin
package main

import (
	"fmt"
)

type celsius float64
type kelvin float64

// kelvinToCelsius converts k to C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius(max kelvin) celsius {
	return celsius(k - max)
}

func main() {
	const v kelvin = 273.15
	var k kelvin = 294.0
	var j celsius = k.celsius(v)
	fmt.Print(k, "* K is ", j, "* C \n")
}
