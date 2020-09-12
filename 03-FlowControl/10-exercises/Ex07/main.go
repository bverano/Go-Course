package main

import "fmt"

func main() {

	age := 21

	if age < 18 {
		fmt.Println("You aren't of legal age")
	} else if age >= 18 && age <= 50 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are elderly")
	}

}
