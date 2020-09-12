package main

import "fmt"

const (
	_  = 2020 + iota
	y1 = 2020 + iota
	y2 = 2020 + iota
	y3 = 2020 + iota
	y4 = 2020 + iota
)

func main() {
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)

}
