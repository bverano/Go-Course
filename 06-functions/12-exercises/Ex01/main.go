package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 30
}

func bar() (int, string) {
	return 21, "Bruno"
}
