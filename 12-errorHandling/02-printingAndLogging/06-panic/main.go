package main

import (
	"os"
)

func main() {
	_, err := os.Open("non_existing.txt")
	if err != nil {
		panic(err)
	}
}
