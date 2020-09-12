package main

import "fmt"

var z = 41 // Declaring a variable using "var" and initializating it with a value

func main() {
	var y int // Declaring a variable using var and specifying it's type

	// Declaring variables
	x := 42 //Short declaration: Always use this kind of declaration inside of blocks, scopes, curly braces

	fmt.Println(x)
	fmt.Println(y)
	number()
}

func number() {
	fmt.Println(z)
}
