package main

import "fmt"

func main() {

	//Unbuffered channel
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch)
}
