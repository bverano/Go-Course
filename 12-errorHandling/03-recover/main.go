package main

import "fmt"

func main() {
	f()
	fmt.Println("Returning normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returning normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer on g", i)
	fmt.Println("Printing g", i)
	g(i + 1)
}
