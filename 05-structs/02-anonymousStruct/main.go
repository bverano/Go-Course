package main

import "fmt"

func main() {
	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Bruno",
		lastName:  "Verano",
		age:       21,
	}
	fmt.Println(p1)
}
