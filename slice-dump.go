// slice-dump
package main

import (
	"fmt"
)

func dump(label string, slice []string) {
	fmt.Printf("%v: lenght %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestial := planets[0:4]
	worlds := append(terrestial, "Ceres")
	fmt.Println(planets)
	fmt.Println(worlds)
}
