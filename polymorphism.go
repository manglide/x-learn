// polymorphism
package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type martian struct{}

type crater struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func (z crater) talk() string {
	return "Women are Wonderful"
}

type starship struct {
	laser
}

func main() {
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martian{})
	shout(&martian{})
	shout(laser(4))
	shout(crater{})

	s := starship{laser(6)}
	fmt.Println(s.talk())
	shout(s)
}
