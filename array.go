// array
package main

import (
	"fmt"
)

func gigi(planets [3]string) {
	for i := range planets {
		planets[i] = "Ambode"
	}
}

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	pp := planets[:]
	fmt.Println(terrestrial, gasGiants, iceGiants)
	fmt.Println(gasGiants[0])
	terrestrialPowers := terrestrial[1:2]
	fmt.Println(terrestrialPowers)
	fmt.Println((terrestrial))
	fmt.Println(pp)

	f := "Aparadija"
	fmt.Println(f[3:])

	question := "¿Cómo estás?"
	fmt.Println(question[:6])
}
