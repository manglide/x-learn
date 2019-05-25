// sort
package main

import (
	"fmt"
	"sort"
)

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
	planets = append(planets, "Pluto", "Amala", "Ayala", "Adenuga")
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
}
