package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	counter := 0

	gs := 100
	wg.Add(gs)

	for i := 1; i <= gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println("Counter:", counter)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
