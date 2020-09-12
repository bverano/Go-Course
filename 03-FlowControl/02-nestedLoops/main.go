package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Main loop:", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tNested loop: %d\n", j)
		}
	}
}
