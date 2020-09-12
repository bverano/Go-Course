package main

import "fmt"

func main() {

	name := "Antonio"

	if len(name) > 6 {
		fmt.Println("Your name is too long")
	} else {
		fmt.Println("Your name is normal")
	}

}
