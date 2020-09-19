package main

import "fmt"

func main() {
	f := factorial(10)
	fmt.Println(f)
	f2 := factorial2(10)
	fmt.Println(f2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorial2(n int) int {
	f := 1
	for ; n > 0; n-- {
		f *= n
	}
	return f
}
