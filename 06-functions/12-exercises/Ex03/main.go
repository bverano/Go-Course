package main

import "fmt"

func main() {
	defer bar()
	foo()
}

func foo() {
	fmt.Println("Hello world!!")
}

func bar() {
	fmt.Println("Something really cool")
}
