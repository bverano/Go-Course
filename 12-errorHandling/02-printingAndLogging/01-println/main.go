package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("non_existing.txt")
	if err != nil {
		fmt.Println(err)
	}
}
