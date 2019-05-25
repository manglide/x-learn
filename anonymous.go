// anonymous
package main

import (
	"fmt"
)

var f = func() {
	fmt.Println("Dress up for the rehearsals")
}

func main() {
	func(message string) {
		fmt.Println(message)
	}("yawa be gas naw naw")
}
