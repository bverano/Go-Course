package main

import "fmt"

func main() {

	p1 := struct {
		name string
		age  int
	}{
		name: "Bruno",
		age:  21,
	}

	fmt.Println(p1)

}
