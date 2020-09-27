package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("non_existing.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}
