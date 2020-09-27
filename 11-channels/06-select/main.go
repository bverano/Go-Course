package main

import "fmt"

func main() {
	pair := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(pair, odd, quit)

	//recieve
	recieve(pair, odd, quit)

	fmt.Println("Finishing...")
}

func send(p, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			p <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func recieve(p, o, q <-chan int) {
	for {
		select {
		case v := <-p:
			fmt.Println("From pair channel:", v)
		case v := <-o:
			fmt.Println("From odd channel:", v)
		case v := <-q:
			fmt.Println("From quit channel:", v)
			return
		}
	}
}
