package main

import "fmt"

func main() {

	//Buffered channel
	ch := make(chan int, 2)
	ch <- 42
	ch <- 43
	fmt.Println(<-ch) //First In First Out
	fmt.Println(<-ch)
}
