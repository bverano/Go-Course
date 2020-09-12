package main

import "fmt"

var a int
var b string = "program"

func main() {

	// Print Line
	fmt.Println(a)
	fmt.Println(b)

	// Print Format
	fmt.Printf("The value of \"a\" is: %d\n", a)
	fmt.Printf("The value of \"b\" is: %s\n", b)
	fmt.Printf("The type of \"a\" is: %T\n", a)
	fmt.Printf("The type of \"b\" is: %T\n", b)

	// Sprint
	s1 := fmt.Sprint("The ", b, " says hello world.")
	fmt.Println(s1)
	fmt.Printf("%T", s1)

}
