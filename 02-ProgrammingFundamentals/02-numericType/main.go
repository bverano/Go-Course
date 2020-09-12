package main

import (
	"fmt"
	"runtime"
)

var x int

func main() {

	x = 42

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//How to see the OS architecture
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
