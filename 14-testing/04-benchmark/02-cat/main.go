package main

import (
	"fmt"
	"strings"

	"github.com/bverano/Go-Course/14-testing/04-benchmark/02-cat/mystr"
)

const s = "Me gusta cuando callas, porque estás como ausente."

func main() {
	xs := strings.Split(s, " ")
	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}
