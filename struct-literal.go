// struct-literal
package main

import (
	"fmt"
)

func main() {
	type location struct {
		lat, long float64
	}
	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	insight := location{-14.5684, 175.472636}
	fmt.Printf("%v\n", insight)
	fmt.Printf("%+v\n", insight)
}
