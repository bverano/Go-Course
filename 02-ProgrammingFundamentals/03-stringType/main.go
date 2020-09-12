package main

import "fmt"

func main() {

	s1 := "Hello world"
	s2 := `This is a break
	line`

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", s2)

	fmt.Println("")

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i, element := range bs {
		fmt.Printf("Index: %d\tASCII: %d\tTranslation: %#U\n", i, element, element)
	}

	fmt.Printf("With the verb %q we can print the string %s", "%s", s1)

}
