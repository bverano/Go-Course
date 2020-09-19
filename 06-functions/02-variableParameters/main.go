package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := foo(s...)
	fmt.Println(sum)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
