package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	switch {
	case now.Year() == 2020:
		fmt.Println("Is coronavirus year")
	case now.Year() == 1999:
		fmt.Println("Is your birth year")
	default:
		fmt.Println("Is a normal year")
	}

}
