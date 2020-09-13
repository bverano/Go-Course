package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	a[0] = 6
	a[1] = 7
	a[2] = 8
	a[3] = 9
	a[4] = 10

	for i, v := range a {
		fmt.Printf("Index: %d\tValue: %d\n", i, v)
	}

	fmt.Printf("%T", a)
}
