// celsius
package main

import (
	"fmt"
)

func main() {
	type celsius float64
	const degrees = 20
	var temperature celsius = degrees
	temperature += 10
	fmt.Println(temperature)

	var warmup float64 = 10

	temperature += celsius(warmup)

	fmt.Println(temperature)
}
