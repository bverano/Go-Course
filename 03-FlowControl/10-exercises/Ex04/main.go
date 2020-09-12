package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	bd := 1999

	for {
		fmt.Println(bd)
		bd++
		if bd > now.Year() {
			break
		}
	}
}
