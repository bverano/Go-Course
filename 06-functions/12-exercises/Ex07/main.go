package main

import "fmt"

func main() {
	sum := func(x ...int) int {
		s := 0
		for _, v := range x {
			s += v
		}
		return s
	}

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
