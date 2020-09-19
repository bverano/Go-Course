package main

import "fmt"

func main() {

	bruno := birthDay(21)
	luis := birthDay(19)

	fmt.Println(bruno())
	fmt.Println(bruno())
	fmt.Println(bruno())
	fmt.Println(bruno())
	fmt.Println(bruno())
	fmt.Println(bruno())
	fmt.Println("-------------------")
	fmt.Println(luis())
	fmt.Println(luis())
	fmt.Println(luis())
	fmt.Println(luis())
	fmt.Println(luis())
	fmt.Println(luis())
	fmt.Println(luis())

}

func birthDay(age int) func() int {
	return func() int {
		age++
		return age
	}
}
