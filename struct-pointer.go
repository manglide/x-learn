// struct-pointer
package main

import (
	"fmt"
)

type person struct {
	name, superpower string
	age              int
}

func main() {
	timmy := &person{
		name: "Adeyemi",
		age:  38,
	}
	fmt.Printf("%+v\n", timmy)
	timmy.superpower = "Jesus Son Of God"
	fmt.Printf("%+v\n", timmy)

	superpowers := &[3]string{"one", "two", "three"}
	fmt.Println(superpowers[0])
}
