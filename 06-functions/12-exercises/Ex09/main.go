package main

import "fmt"

func sum(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}

func getSum(n int, f func(x ...int) int) int {
	var numSlice []int
	for ; n > 0; n-- {
		numSlice = append(numSlice, n)
	}
	return f(numSlice...)
}

func main() {

	fmt.Println(getSum(10, sum))

}
