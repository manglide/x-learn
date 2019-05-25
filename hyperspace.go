// hyperspace
package main

import (
	"fmt"
	"strings"
)

func hyperspace(planets []string) {
	for i := range planets {
		planets[i] = strings.TrimSpace(planets[i])
	}
}

func main() {
	planets := []string{" Gang   ", "See Me see   ", " Futures "}
	hyperspace(planets)
	fmt.Println(strings.Join(planets, ""))
}
