package main

import "fmt"

func main() {
	x := 0
	fmt.Println("X Before", x)
	fmt.Println("X Before", &x)
	foo(&x)
	fmt.Println("X After", x)
	fmt.Println("X After", &x)
}

func foo(y *int) {
	fmt.Println("Y Before", y)
	fmt.Println("Y Before", *y)
	*y = 21
	fmt.Println("Y After", y)
	fmt.Println("Y After", *y)
}
