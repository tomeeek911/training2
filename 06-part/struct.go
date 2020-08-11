package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	person1 := person{
		first: "James",
		last:  "Bond",
		age:   27,
	}
	person2 := person{
		first: "M",
		last:  "M",
		age:   21,
	}
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.first, person1.last, person1.age)
	fmt.Println(person2.first, person2.last, person2.age)

}
