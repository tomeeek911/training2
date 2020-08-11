package main

import "fmt"

type person struct {
	firstname string
	age       int
}

type footballer struct {
	person
	skilled bool
}

func (y person) human() {
	fmt.Println("Hey, I am here: ", y.firstname, y.age)
}

func main() {
	p1 := person{
		firstname: "James",
		age:       42,
	}
	p2 := footballer{
		person: person{
			firstname: "John",
			age:       23,
		},
		skilled: true,
	}
	fmt.Println(p2)
	p1.human()
}
