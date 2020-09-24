// go run -race main.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("Go routines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			runtime.Gosched()
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go routines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Count:", counter)
}
