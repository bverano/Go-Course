package main

import (
	"fmt"
	"sync"
)

func main() {
	pair := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	//send
	go send(pair, odd)

	//recieve
	go recieve(pair, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Finishing...")
}

func send(p, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			p <- i
		} else {
			o <- i
		}
	}
	close(p)
	close(o)
}

func recieve(p, o <-chan int, f chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range p {
			f <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			f <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(f)
}
