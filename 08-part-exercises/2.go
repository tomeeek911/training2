package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p := person{
		first: "Magda",
		last:  "Sabiniewicz",
	}
	fmt.Println(p)
	ChangeMe(&p)
	fmt.Println(p)
}

func ChangeMe(y *person) {
	y.first = "Elvis"
	(*y).first = "Presley"
}
