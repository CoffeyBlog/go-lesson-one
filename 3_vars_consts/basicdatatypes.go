//Go code goes here:

package main

import "fmt"


// Start of Go Program
func main(){

	var a int32				// Here we are asking Go to create us a "32 bit (4 groups of 8 bits) / 8 bytes - block of storage in memory location"
	a = 33					// The we are putting / assigning an int of 30 to that blocked of location of storage.

	fmt.Println("The value of a is ", a)



	var b bool          	//We created a "block of memory" for a bool data type
	b = true				//The data type is true - thus it is = to - 1 - ONE

	fmt.Println("The value of b is", b)



	var c bool    			//We created a "block of memory" for a bool data type
	c = false				//The data type is true - thus it is = to - 0 -   ZERO

	fmt.Println("The value of c is", c)



	var d float32 			//We created a "block of memory" for a float data type - it is = to 4 bits of memory
	d = 3.1415927

	fmt.Println("The value of d is", d)



	var e string			//We created a "block of memory" for a string data type
	e = "Go is Awesome"

	fmt.Println(e)

}

