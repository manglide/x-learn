// birthday
package main

import (
	"fmt"
)

type person struct {
	name, superpower string
	age              int
}
type perSon struct {
	name string
	age  int
}

func birthday(p *person) {
	p.age++
}

func (p *perSon) human() {
	p.age++
}
func main() {
	rebecca := person{
		name:       "Adeyem",
		superpower: "Jesus Christ of nazareth",
		age:        38,
	}
	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)
	joshua := perSon{
		name: "Josh UI/UX",
		age:  27,
	}
	joshua.human()
	fmt.Printf("%+v\n", joshua)
}
