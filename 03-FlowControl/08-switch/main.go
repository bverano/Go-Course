package main

import "fmt"

func main() {

	switch "Apple" {
	case "Lemon", "Pear":
		fmt.Println("Green fruits")
	case "Apple", "Plum", "Strawberry":
		fmt.Println("Red fruits")
		fallthrough //run the function of the statement bellow, even if it's not true
	default:
		fmt.Println("Printing from default")
	}

}
