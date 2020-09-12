package main

import "fmt"

const a = 42
const b = 42.12
const c = /*string*/ "Bruno Verano"

type name string

var d name

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	d = c
}
