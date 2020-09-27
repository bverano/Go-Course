package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {

	dir, err := os.Getwd()

	f, err := os.Create(path.Join(dir, "02-printingAndLogging", "03-logSetOutput", "log.txt"))
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("non_existing.txt")
	if err != nil {
		log.Println("An error has ocurred:", err)
	}
	defer f2.Close()
	fmt.Println("Check Log.txt file on directory")
}
