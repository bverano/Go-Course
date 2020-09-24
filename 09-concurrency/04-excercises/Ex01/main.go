package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Simple hello world")

	wg.Add(2)

	go func() {
		fmt.Println("Go routine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Go routine 2")
		wg.Done()
	}()

	wg.Wait()
}
