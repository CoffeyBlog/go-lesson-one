package main

import "fmt"

func main() {

	age := 39


	if age < 12 {
		fmt.Println("Wow you're young")
	}

	if age > 13 && age < 19 {
		fmt.Println("You are a teenager")
	}

	if age > 20 && age < 30 {
		fmt.Println("You are in your twenties")
	}

	if age > 31 && age < 40 {
		fmt.Println("You are a thirties")
	}

	if age > 41 && age < 50 {
		fmt.Println("You are in your fourties")
	}

	if age > 51 && age < 60 {
		fmt.Println("You are a fifties")
	}

	if age > 61 && age < 70 {
		fmt.Println("You are in your sixties")
	}

}

