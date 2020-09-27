package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go send(c)

	//recieve
	recieve(c)

	fmt.Println("Ending...")
}

func send(c chan<- int) {
	c <- 42
}

func recieve(c <-chan int) {
	fmt.Println(<-c)
}
