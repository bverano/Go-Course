package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range s {
		fmt.Printf("Index: %d\tValue: %d\n", i, v)
	}
	fmt.Printf("%T", s)
}
