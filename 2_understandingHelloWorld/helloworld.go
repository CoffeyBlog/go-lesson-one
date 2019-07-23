
//Go like most modern languages is case sensitive

//Go is very particular about the curly braces - And has a built in style guide - it "seems" like the designers of the go language go tired of everyone wasting time
// and bitching about how code should be written and instead just decided themselves. -- At first you will hate this but then you will LOVE IT.




//Packages are like folders they help us organize things
// When we create package main - we are creating an executable program that we can run
// If we declare it with a different package name --  we are then creating a shared library, which is code that has been grouped together to be used together
package main

//Sometimes you will need to import things into your program, like: import "math/rand" - this allows you to draw a random number
//As mentioned the designers of Go were tired intolerant of those who aregued about "code style" thus was born the "fmt" or format library
import "fmt"


//This is the function declaration - as you can see it looks very similiar to Python + Java
// The fmt.Println allows the program to print the data but in the correct format.
// Important note The program does not actually start running until line #24
func main() {
	fmt.Println("Hello Cat")
// White space is not sensitive to ... but you should always write clean professional code.
	fmt.Println("Hello Dog")
}



//Note you can't have any unused variables - nor unused imported libraries for "clean code" purposes