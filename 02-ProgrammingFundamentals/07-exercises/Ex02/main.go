package main

import "fmt"

func main() {
	a := 10 == 20
	b := 5 <= 10
	c := 2 >= 1
	d := 10 != 10
	e := 3 < 10
	f := 10 > 1

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
