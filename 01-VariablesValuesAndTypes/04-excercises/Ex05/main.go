package main

import "fmt"

type age int

var x age
var y int

func main() {

	fmt.Println("X")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	fmt.Println("-------------------------------")

	fmt.Println("Y")
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
