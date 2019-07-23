package main

import (
	"fmt"
	"time"
)

func main() {


	// While Loop
	countdown := 10
	for countdown > 0 {
		fmt.Println(countdown )
		time.Sleep(time.Second)   // if we stopped here - we'd have an infinite loop because countdown is not decreasing

		countdown = countdown - 1
	}
	fmt.Println("Blastoff")



	//For Loop
	for countdownTwo := 10; countdownTwo > 0; countdownTwo-- {
		fmt.Println(countdownTwo)
		time.Sleep(time.Second) // if we stopped here - we'd have an infinite loop because countdown is not decreasing
	}
	fmt.Println("Second Blastoff")



	//For Loop - decrementing by twos
	for countbyTwo := 10; countbyTwo > 0; countbyTwo = countbyTwo - 2 {
		fmt.Println(countbyTwo)
		time.Sleep(time.Second) // if we stopped here - we'd have an infinite loop because countdown is not decreasing
	}
	fmt.Println("Count By Twos")


}






