// variadic
package main

import (
	"fmt"
)

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
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
	twoWorlds := terraform("Hello World!", planets...)
	fmt.Println(twoWorlds)
}
