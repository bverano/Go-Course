package main

import "fmt"

func main() {
	cr := make(<-chan int, 2)
	cs := make(chan<- int, 3)
	cb := make(chan int, 4)

	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
	fmt.Printf("%T\n", cb)

	//! fmt.Printf("%T\n", (chan int)(cr)) ERROR
	//! fmt.Printf("%T\n", (chan int)(cs)) ERROR
	fmt.Println("---------")
	fmt.Printf("%T\n", (chan<- int)(cb))
	fmt.Printf("%T\n", (<-chan int)(cb))

}
