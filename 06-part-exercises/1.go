package main

import "fmt"

type person struct {
	first string
	last  string
	fif   []string
}

func main() {
	p1 := person{
		first: "Tomasz",
		last:  "Sabiniewicz",
		fif: []string{
			"blueberry",
			"strawberry",
		},
	}
	p2 := person{
		first: "Magdalena",
		last:  "Sabiniewicz",
		fif: []string{
			"chocolate",
			"carmel",
		},
	}
	fmt.Println(p1.first, p1.last)
	for i, v := range p1.fif {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first, p2.last)
	for i, v := range p2.fif {
		fmt.Println(i, v)
	}
}
