package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("non_existing.txt")
	if err != nil {
		log.Println("An error has ocurred:", err)
	}
}
