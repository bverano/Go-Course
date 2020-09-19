package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("This is an anonymous function")
	}()

	func(x int) {
		fmt.Println("The meaning of life is", x)
	}(42)
}

func foo() {
	fmt.Println("foo was executed")
}
