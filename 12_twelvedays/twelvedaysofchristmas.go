//twelve days of christmas - using switch statement and code blocks

package main

import (
	"fmt"
)

func main() {

	for day := 1; day <= 12; day++ { // if DAY 1 ---> is equal to - or less than 12 ---> increment days by one

		fmt.Println("On the ", day, " of Christmas, my true love gave to me ... ")

		switch day {
		case 12:
			fmt.Println("12 drummers drumming")

		case 11:
			fmt.Println("Eleven Pipers Piping")

		case 10:
			fmt.Println("Ten Lords a Leaping")

		case 9:
			fmt.Println("Nine Ladies Dancing")

		case 8:
			fmt.Println("Eight Maids a Milking")

		case 7:
			fmt.Println("Seven Swans a Swimming")

		case 6:
			fmt.Println("Six Geese a Laying")

		case 5:
			fmt.Println("Five Golden Rings")

		case 4:
			fmt.Println("Four Calling Birds")

		case 3:
			fmt.Println("Three French Hens")

		case 2:
			fmt.Println("Two Turtle doves")

		case 1:
			fmt.Println("A Patridge in a pear tree")

		}
	}
}
