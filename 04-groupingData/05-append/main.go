package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x)

	x = append(x, 9, 10)
	fmt.Println(x)

	y := []int{11, 12, 13, 14}

	x = append(x, y...)
	fmt.Println(x)

}
