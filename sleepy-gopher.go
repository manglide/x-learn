// sleepy-gopher
package main

import (
	"fmt"
	"time"
)

func sleepyGopher() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(" ... snoring ... ")
	}
	go dance()
}

func dance() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	go sleepyGopher()
}

func zenit() {
	for v := 0; v < 20; v++ {
		time.Sleep(1 * time.Second)
		fmt.Println(v)
	}
}

func main() {
	go dance()
	go zenit()
	go sleepyGopher()
	time.Sleep(22 * time.Second)
	fmt.Println("Pens Up!!!")
}
