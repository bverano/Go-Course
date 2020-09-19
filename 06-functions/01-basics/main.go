package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("MoneyPenny")
	fmt.Println(s1)
	x, y := sayHi("Bruno", "Verano")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func sayHi(n string, l string) (string, bool) {
	p := fmt.Sprint(n, " ", l, ` says "hi".`)
	q := true
	return p, q
}
