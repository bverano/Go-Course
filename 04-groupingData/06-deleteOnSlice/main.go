package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}

	x = append(x[:5], x[7:]...)
	fmt.Println(x)
}
