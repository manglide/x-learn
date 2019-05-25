package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count, "")
		time.Sleep(time.Second)
		if rand.Intn(100) == 32 {
			fmt.Println("Mission Aborted. Emergency Calls Alerted!!!!!!!")
			break
		} else {
			count--
		}
		if count == 0 {
			fmt.Println("Liftoff!!!")
		}
 	}
}