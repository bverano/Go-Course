package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	by := 1999
	for by <= now.Year() {
		fmt.Println(by)
		by++
	}
}
