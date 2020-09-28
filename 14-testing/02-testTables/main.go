package main

import "fmt"

func main() {
	fmt.Println("The sum of 2 + 4 is:", mySum(2, 4))
	fmt.Println("The sum of 1 + 5 is:", mySum(1, 5))
	fmt.Println("The sum of 4 + 4 is:", mySum(4, 4))
	fmt.Println("The sum of 7 + 3 is:", mySum(7, 3))
}

func mySum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
