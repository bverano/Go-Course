package main

import "fmt"

func getSum(x int) func() int {
	var numSlice []int
	for ; x > 0; x-- {
		numSlice = append(numSlice, x)
	}
	return func() int {
		s := 0
		for _, v := range numSlice {
			s += v
		}
		return s
	}
}

func main() {
	f := getSum(90)
	fmt.Println(f())
}
