package main

import "fmt"

func main() {
	defer foo()
	defer woo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func woo() {
	fmt.Println("woo")
}
