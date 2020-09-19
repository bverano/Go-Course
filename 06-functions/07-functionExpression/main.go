package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Mi first function expression")
	}

	f()

	g := func(x int) {
		fmt.Println("America was discovered in", x)
	}
	g(1492)
}
