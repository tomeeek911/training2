package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (x person) speak() {
	fmt.Println("This is a method which calls: ", x.first, x.last, x.age)
}

func main() {
	p1 := person{
		first: "Tomasz",
		last:  "Sabiniewicz",
		age:   29,
	}
	p2 := person{
		first: "Magda",
		last:  "Sabiniewicz",
		age:   28,
	}
	p3 := person{
		first: "Hania",
		last:  "Sabiniewicz",
		age:   2,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	p3.speak()
}
