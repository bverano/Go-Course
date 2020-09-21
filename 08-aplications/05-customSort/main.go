package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byName []person

func (n byName) Len() int           { return len(n) }
func (n byName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n byName) Less(i, j int) bool { return n[i].Name < n[j].Name }

func main() {
	p1 := person{"Eduar", 32}
	p2 := person{"María", 25}
	p3 := person{"Carolina", 56}
	p4 := person{"Adriana", 60}

	persons := []person{p1, p2, p3, p4}

	sort.Sort(byName(persons))

	fmt.Println(persons)
}
