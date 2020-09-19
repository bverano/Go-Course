package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(foo(s1...))
	fmt.Println(bar(s1))

}

func foo(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}

func bar(x []int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}
