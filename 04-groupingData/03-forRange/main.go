package main

import "fmt"

func main() {

	x := []int{2, 3, 4, 5}

	fmt.Println(x)
	fmt.Println("Capacity:", cap(x))

	for i, v := range x {
		fmt.Printf("Index: %d\tValue: %d\n", i, v)
	}

}
