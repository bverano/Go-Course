package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := `123456`
	bs, err := bcrypt.GenerateFromPassword([]byte(p), 4)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Hashed Password: %s\n", bs)

	err = bcrypt.CompareHashAndPassword(bs, []byte("123456"))
	if err != nil {
		fmt.Println("Wrong password")
		return
	}
	fmt.Println("Successfully logged in")

}
