package main

import "fmt"

func main() {
	bv := []string{"Bruno", "Verano", "Skate"}
	jv := []string{"Jacobo", "verano", "Among us"}

	fmt.Println(bv)
	fmt.Println(jv)

	mp := [][]string{bv, jv}
	fmt.Println(mp)
}
