package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When you divide %d by 4, the rest is %d\n", i, i%4)
	}
}
