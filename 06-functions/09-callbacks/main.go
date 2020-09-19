package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := sum(ii...)
	s2 := sumPairs(sum, ii...)
	s3 := sumOdds(sum, ii...)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

}

func sum(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}

func sumPairs(f func(x ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}

func sumOdds(f func(x ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
