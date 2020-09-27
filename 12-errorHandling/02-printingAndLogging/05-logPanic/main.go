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
		log.Panicln(err)
	}
}

func foo() {
	fmt.Println("When we use panic, deferred functions will run")
}
