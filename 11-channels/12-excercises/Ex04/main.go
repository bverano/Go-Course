package main

import (
	"fmt"
)

func main() {
	quit := make(chan int)
	c := gen(quit)

	recieve(c, quit)

	fmt.Println("About to end.")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func recieve(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Value:", v)
		case <-q:
			fmt.Println("From quit channel")
			return
		}
	}
}
