package main

import "fmt"

func main() {
	var x [5]int
	x[3] = 15

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
}
