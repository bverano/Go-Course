package main

import "fmt"

func main() {

	superheroes := [][]string{
		[]string{"Batman", "Boss", "Black suit"},
		[]string{"Robin", "Assistant", "Colors suit"},
	}

	for i, v := range superheroes {
		fmt.Printf("Superhero #%d:\n", i+1)
		for _, v2 := range v {
			fmt.Printf("\t%s\n", v2)
		}
	}
}
