package main

import "fmt"

func main() {

	//Composite literal
	x := []int{2, 3, 4, 10}

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))

}
